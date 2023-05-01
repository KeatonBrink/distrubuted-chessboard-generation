package main

import (
	"context"
	"flag"
	"log"
	"net"
	"time"

	dcg "github.com/KeatonBrink/distrubuted-chessboard-generation/distributedchessboardgeneration"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	masterAddrPtr := flag.String("masterAddr", getLocalAddress()+":3410", "a String")
	portPtr := flag.String("port", "3411", "a String")
	flag.Parse()
	port := *portPtr
	masterAddress := *masterAddrPtr
	myAddress := getLocalAddress() + ":" + port
	for {
		curReply := getBoard(masterAddress, myAddress)
		log.Println("Tasked Board:")
		dcg.CBoardify(curReply.GetBoard()).PrintSelf()
		done := curReply.GetIsFinished()
		if done {
			break
		}
		curBoard := dcg.CBoardify(curReply.GetBoard())
		curPieceCount := curBoard.CountAllPieces()
		boardMap := make(map[string][]string)
		curNextMoves := curBoard.CreateNextMoves()
		todoBoardQueue := curNextMoves
		hasQueued := make(map[string]bool)
		hasQueued[curBoard.Stringify()] = true
		boardMap[curBoard.Stringify()] = dcg.StringifySliceCB(curNextMoves)
		for len(todoBoardQueue) > 0 {
			if len(todoBoardQueue) > 1 {
				curBoard, curNextMoves, todoBoardQueue = todoBoardQueue[0], todoBoardQueue[0].CreateNextMoves(), todoBoardQueue[1:]
			} else {
				curBoard, curNextMoves, todoBoardQueue = todoBoardQueue[0], todoBoardQueue[0].CreateNextMoves(), []dcg.Chessboard{}
			}

			// Too many repeats are ending up in here
			boardMap[curBoard.Stringify()] = dcg.StringifySliceCB(curNextMoves)
			for _, curNextMove := range curNextMoves {
				_, curHasQueued := hasQueued[curNextMove.Stringify()]
				if _, ok := boardMap[curNextMove.Stringify()]; !ok && curNextMove.CountAllPieces() == curPieceCount && !curHasQueued {
					todoBoardQueue = append(todoBoardQueue, curNextMove)
					hasQueued[curNextMove.Stringify()] = true
				}
			}
		}
		log.Println("Returning board")
		returnBoards(masterAddress, myAddress, boardMap)
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	BoardReply, err := client.GetCb(ctx, &dcg.Message{Ip: myAddr})
	if err != nil {
		log.Fatalf("Could not get Chessboard: %v", err)
	}
	return BoardReply
}
func returnBoards(masterAddr, myAddress string, curBoards map[string][]string) {
	log.Println("returnBoards is not currently implemented")
	// TODO
	// client, err := rpc.DialHTTP("tcp", string(masterAddr))
	// if err != nil {
	// 	log.Fatalf("rpc.DialHTTP(Return_Board): %v", err)
	// }
	// var reply Nothing
	// if err = client.Call("MNode.Return_Board", curBoards, &reply); err != nil {
	// 	log.Fatalf("client.Call(Return_Board): %v", err)
	// }
	// if err = client.Close(); err != nil {
	// 	log.Fatalf("client.Close(Return_Board) %v", err)
	// }
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
