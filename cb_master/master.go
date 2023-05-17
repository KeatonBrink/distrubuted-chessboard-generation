package main

import (
	"context"
	"flag"
	"fmt"
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
	BoardsToSQL    chan map[string][]string // Might consider making this a stream or get rid of it altogether
	FinishedBoards map[string]bool
}

var Mutex sync.Mutex

func (s *MNodeServer) GetCb(ctx context.Context, in *dcg.Message) (*dcg.ChessboardString, error) {
	boardTask, ok := <-s.BoardsToDo
	finished := false
	if !ok {
		finished = true
	}
	log.Println("Tasking a Board: ", boardTask)
	return &dcg.ChessboardString{Board: boardTask, IsFinished: finished}, nil
}

func (s *MNodeServer) ReturnCb(ctx context.Context, in *dcg.ReturnMessage) (*dcg.Emptyy, error) {
	log.Println("Recieving a board from: ", in.Ip)
	dcg.Download(makeURL(in.Ip, in.FileName), filepath.Join(s.MyTempPath, in.FileName))
	return &dcg.Emptyy{}, nil
}

// func (n *MNode) Return_Board(curBoards map[string][]string, reply *Nothing) error {
// 	log.Println("Recieving Finished Board")
// 	Mutex.Lock()
// 	defer Mutex.Unlock()
// 	for mstring := range curBoards {
// 		if _, ok := n.FinishedBoards[mstring]; !ok {
// 			n.FinishedBoards[mstring] = true
// 		} else {
// 			// Note, a rook/Knight on one side will have all the same moves as the other side
// 			//  And consequently throw this error.  Slight bug that leads to extra processing,
// 			//  but not detrimental to overall system design
// 			// log.Printf("Error in Returning computed board:\n%v", mstring)

// 			// Returning nil will avoid double entries in the SQL database
// 			return nil
// 		}
// 	}
// 	n.BoardsToSQL <- curBoards
// 	return nil
// }

func newServer(myAddress, tempdir string) *MNodeServer {
	s := &MNodeServer{BoardsToDo: make(chan string, 10), BoardsToSQL: make(chan map[string][]string, 2), MyIP: myAddress, MyTempPath: tempdir}
	go dcg.NextIterativeBoard(s.BoardsToDo)
	return s
}

func main() {
	portPtr := flag.String("port", "3410", "a String")
	tempDirPtr := flag.String("tempdir", filepath.Join(os.TempDir(), fmt.Sprintf("cbMaster.%d", os.Getpid())), "a String")
	tempdir := *tempDirPtr
	flag.Parse()
	port := *portPtr
	myAddress := getLocalAddress() + ":" + port
	lis, err := net.Listen("tcp", myAddress)
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
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	dcg.RegisterChessboardTaskAssignmentServer(grpcServer, newServer(myAddress, tempdir))
	dcg.RegisterChessboardReturnAssignmentURLServer(grpcServer, newServer(myAddress, tempdir))
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
