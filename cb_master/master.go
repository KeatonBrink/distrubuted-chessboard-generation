package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path/filepath"
	"sync"

	dcg "github.com/KeatonBrink/distrubuted-chessboard-generation/distributedchessboardgeneration"
	"google.golang.org/grpc"
)

type MNodeServer struct {
	dcg.UnimplementedChessboardTaskAssignmentServer
	dcg.UnimplementedChessboardReturnAssignmentURLServer
	MyIP           string
	MyTempPath     string
	BoardsToDo     chan string
	FinishedBoards map[string]bool
}

var Mutex sync.Mutex

func (s *MNodeServer) GetCb(ctx context.Context, in *dcg.Message) (*dcg.ChessboardString, error) {
	finished := false
	boardTask := ""
	var ok bool
	for {
		boardTask, ok = <-s.BoardsToDo
		if !ok {
			finished = true
			break
		}
		_, boardIsDone := s.FinishedBoards[fmt.Sprintf("%v.db", boardTask)]
		if !boardIsDone {
			break
		}
	}
	log.Println("Tasking a Board: ", boardTask)
	return &dcg.ChessboardString{Board: boardTask, IsFinished: finished}, nil
}

func (s *MNodeServer) ReturnCb(ctx context.Context, in *dcg.ReturnMessage) (*dcg.Emptyy, error) {
	log.Println("Recieving a board from: ", in.Ip)
	dcg.Download(makeURL(in.Ip, in.FileName), filepath.Join(s.MyTempPath, in.FileName))
	return &dcg.Emptyy{}, nil
}

func newServer(myAddress, tempdir string, fileNames map[string]bool) *MNodeServer {
	s := &MNodeServer{BoardsToDo: make(chan string, 10), MyIP: myAddress, MyTempPath: tempdir, FinishedBoards: fileNames}
	go dcg.NextIterativeBoard(s.BoardsToDo)
	return s
}

func main() {
	portPtr := flag.String("port", "3410", "a String")
	tempDirPtr := flag.String("tempdir", "Data", "a String")
	myAddrPtr := flag.String("myAddr", getLocalAddress(), "a String")
	flag.Parse()
	tempdir := *tempDirPtr
	port := *portPtr
	myAddr := *myAddrPtr
	myAddress := myAddr + ":" + port
	lis, err := net.Listen("tcp", myAddress)
	fileNames := make(map[string]bool)
	if err != nil {
		log.Fatalln(err)
	}
	dirExists, err := exists(tempdir)
	if err != nil {
		log.Fatalf("Error creating tempdir: %v", err)
	}
	if !dirExists {
		err := os.Mkdir(tempdir, 0750)
		if err != nil {
			log.Fatalf("Error creating tempdir: %v", err)
		}
	} else {
		files, err := ioutil.ReadDir(tempdir)
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			fileNames[file.Name()] = true
		}
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	newServerVar := newServer(myAddress, tempdir, fileNames)
	dcg.RegisterChessboardTaskAssignmentServer(grpcServer, newServerVar)
	dcg.RegisterChessboardReturnAssignmentURLServer(grpcServer, newServerVar)
	log.Printf("Master started with address: %v", myAddress)
	grpcServer.Serve(lis)
}

func getLocalAddress() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}

// exists returns whether the given file or directory exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func makeURL(host, file string) string { return fmt.Sprintf("http://%s/data/%s", host, file) }
