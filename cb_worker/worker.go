package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"

	dcg "github.com/KeatonBrink/distrubuted-chessboard-generation/distributedchessboardgeneration"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type NextMoves struct {
	cMove  string
	nMoves []string
}

func main() {
	masterAddrPtr := flag.String("masterAddr", getLocalAddress()+":3410", "a String")
	portPtr := flag.String("port", "3411", "a String")
	tempDirPtr := flag.String("tempdir", filepath.Join(os.TempDir(), fmt.Sprintf("cbWorker.%d", os.Getpid())), "a String")
	myAddrPtr := flag.String("myAddr", getLocalAddress(), "a String")
	flag.Parse()
	myAddr := *myAddrPtr
	port := *portPtr
	masterAddress := *masterAddrPtr
	tempdir := *tempDirPtr
	// Create temporary directory for storing db files
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
	defer os.RemoveAll(tempdir)
	myAddress := myAddr + ":" + port
	http.Handle("/data/", http.StripPrefix("/data", http.FileServer(http.Dir(tempdir))))
	go func() {
		if err := http.ListenAndServe(myAddress, nil); err != nil {
			log.Printf("Error in HTTP server for %s: %v", myAddress, err)
		}
	}()
	for {
		// Grab board from master
		curReply := getBoard(masterAddress, myAddress)
		log.Println("Tasked Board:")
		dcg.CBoardify(curReply.GetBoard()).PrintSelf()
		done := curReply.GetIsFinished()
		if done {
			break
		}
		// SQLite routine
		isDoneAndFileName := make(chan string)
		moveChan := make(chan NextMoves, 100)
		go SQLMoves(moveChan, isDoneAndFileName, tempdir)
		// Creates first set of moves from given board
		curBoard := dcg.CBoardify(curReply.GetBoard())
		curPieceCount := curBoard.CountAllPieces()
		var curNextMoves []dcg.Chessboard
		moveChan <- NextMoves{cMove: curBoard.Stringify(), nMoves: dcg.StringifySliceCB(curNextMoves)}
		var todoBoardQueue []dcg.Chessboard
		todoBoardQueue = append(todoBoardQueue, curBoard)
		hasQueued := make(map[string]bool)
		hasQueued[curBoard.Stringify()] = true
		for len(todoBoardQueue) > 0 {
			// If the todo board is longer than 1, pop off the start for work
			if len(todoBoardQueue) > 1 {
				curBoard, curNextMoves, todoBoardQueue = todoBoardQueue[0], todoBoardQueue[0].CreateNextMoves(), todoBoardQueue[1:]
				// Else grab the only item and empty the queue
			} else {
				curBoard, curNextMoves, todoBoardQueue = todoBoardQueue[0], todoBoardQueue[0].CreateNextMoves(), []dcg.Chessboard{}
			}
			if _, ok1 := hasQueued[curBoard.Stringify()]; !ok1 {
				log.Fatalf("curBoard not in hasQueued: %v", curBoard.Stringify())
			}
			// Send the board and next moves to sql
			moveChan <- NextMoves{cMove: curBoard.Stringify(), nMoves: dcg.StringifySliceCB(curNextMoves)}
			// For all of the next moves
			for _, curNextMove := range curNextMoves {
				// If the next moves have not been queued or computed, then append
				if _, ok := hasQueued[curNextMove.Stringify()]; !ok && curNextMove.CountAllPieces() == curPieceCount {
					todoBoardQueue = append(todoBoardQueue, curNextMove)
					hasQueued[curNextMove.Stringify()] = true
				}
			}
		}
		close(moveChan)
		retFileName := <-isDoneAndFileName
		log.Println("Returning board")
		returnBoards(masterAddress, myAddress, retFileName)
	}
}

func getBoard(masterAddr, myAddr string) *dcg.ChessboardString {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(masterAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := dcg.NewChessboardTaskAssignmentClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()
	BoardReply, err := client.GetCb(ctx, &dcg.Message{Ip: myAddr})
	if err != nil {
		log.Fatalf("Could not get Chessboard: %v", err)
	}
	return BoardReply
}

func SQLMoves(inputChan <-chan NextMoves, isDone chan<- string, tempdir string) {
	curNextMoves := <-inputChan
	cfilename := fmt.Sprint(curNextMoves.cMove, ".db")
	cfilepath := filepath.Join(tempdir, cfilename)
	curDB, err := dcg.OpenDatabase(cfilepath)
	if err != nil {
		log.Fatalf("error openDatabase: %v", err)
	}
	defer curDB.Close()
	// Tracking variables
	curRowID := 1
	allRowIDs := make(map[string]int)
	_, err = curDB.Exec("CREATE TABLE chessboards (id INTEGER PRIMARY KEY, chessboard TEXT NOT NULL)")
	if err != nil {
		log.Fatalf("error SQLMoves: Create Table chessboards: %v", err)
	}
	_, err = curDB.Exec("CREATE TABLE moves (id INTEGER PRIMARY KEY, from_id INTEGER NOT NULL REFERENCES chessboards(id) ON DELETE CASCADE, to_id INTEGER NOT NULL)")
	if err != nil {
		log.Fatalf("error SQLMoves: Create Table Moves: %v", err)
	}
	_, err = curDB.Exec("INSERT INTO chessboards (chessboard) VALUES (?)", curNextMoves.cMove)
	if err != nil {
		log.Fatalf("error SQLMoves: INSERT INTO chessboards: %v", err)
	}
	allRowIDs[curNextMoves.cMove] = curRowID
	curRowID++
	for _, nmove := range curNextMoves.nMoves {
		_, err = curDB.Exec("INSERT INTO chessboards (chessboard) VALUES (?)", nmove)
		if err != nil {
			log.Fatalf("error SQLMoves: INSERT INTO chessboards: %v", err)
		}
		allRowIDs[nmove] = curRowID
		curRowID++
		_, err = curDB.Exec("INSERT INTO moves (from_id, to_id) VALUES (?, ?)", allRowIDs[curNextMoves.cMove], allRowIDs[nmove])
		if err != nil {
			log.Fatalf("error SQLMoves: INSERT INTO move: %v", err)
		}
	}
	for curNextMoves = range inputChan {
		if _, ok := allRowIDs[curNextMoves.cMove]; !ok {
			_, err = curDB.Exec("INSERT INTO chessboards (chessboard) VALUES (?)", curNextMoves.cMove)
			if err != nil {
				log.Fatalf("error SQLMoves: INSERT INTO chessboards: %v", err)
			}
			allRowIDs[curNextMoves.cMove] = curRowID
			curRowID++
		}
		for _, nmove := range curNextMoves.nMoves {
			if _, ok := allRowIDs[nmove]; !ok {
				_, err = curDB.Exec("INSERT INTO chessboards (chessboard) VALUES (?)", nmove)
				if err != nil {
					log.Fatalf("error SQLMoves: INSERT INTO chessboards: %v", err)
				}
				allRowIDs[nmove] = curRowID
				curRowID++
			}
			_, err = curDB.Exec("INSERT INTO moves (from_id, to_id) VALUES (?, ?)", allRowIDs[curNextMoves.cMove], allRowIDs[nmove])
			if err != nil {
				log.Fatalf("error SQLMoves: INSERT INTO move: %v", err)
			}
		}
	}
	isDone <- cfilename
}

// Return board will have a string indicating the original board, and worker ip for grabbing db
func returnBoards(masterAddr, myAddress, fileName string) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(masterAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := dcg.NewChessboardReturnAssignmentURLClient(conn)
	start := time.Now()
	defer closingBoard(start)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30000)
	defer cancel()
	_, err = client.ReturnCb(ctx, &dcg.ReturnMessage{Ip: myAddress, FileName: fileName})
	if err != nil {
		log.Fatalf("Could not return Chessboard: %v", err)
	}

}

func closingBoard(nower time.Time) {
	end := time.Now()
	log.Printf("Time required to return board: %v", end.Sub(nower))
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
