package distributedchessboardgeneration

import (
	"flag"
	"log"
	"net"
	"net/rpc"
)

func Start() error {
	// Two options here worker and master debugging

	// Worker debug
	// rolePtr := flag.String("role", "worker", "a String: master, worker")
	// masterAddrPtr := flag.String("masterAddr", getLocalAddress()+":3410", "a String")
	// portPtr := flag.String("port", "3440", "a String")

	// Master Debug
	rolePtr := flag.String("role", "master", "a String: master, worker")
	masterAddrPtr := flag.String("masterAddr", getLocalAddress()+":3410", "a String")
	portPtr := flag.String("port", "3410", "a String")
	flag.Parse()
	if *rolePtr == "master" {
		err := masterNode(*portPtr)
		if err != nil {
			return err
		}
	} else {
		err := workerNode(*portPtr, *masterAddrPtr)
		if err != nil {
			return err
		}
	}
	return nil
}

func workerNode(port, masterAddress string) error {
	myAddress := getLocalAddress() + ":" + port
	for {
		curReply := getBoard(masterAddress)
		log.Println("Tasked Board:")
		curReply.Board.PrintSelf()
		done := curReply.IsDone
		if done {
			break
		}
		curBoard := curReply.Board
		curPieceCount := curBoard.CountAllPieces()
		boardMap := make(map[string][]string)
		curNextMoves := curBoard.CreateNextMoves()
		todoBoardQueue := curNextMoves
		hasQueued := make(map[string]bool)
		hasQueued[curBoard.Stringify()] = true
		boardMap[curBoard.Stringify()] = StringifySliceCB(curNextMoves)
		for len(todoBoardQueue) > 0 {
			if len(todoBoardQueue) > 1 {
				curBoard, curNextMoves, todoBoardQueue = todoBoardQueue[0], todoBoardQueue[0].CreateNextMoves(), todoBoardQueue[1:]
			} else {
				curBoard, curNextMoves, todoBoardQueue = todoBoardQueue[0], todoBoardQueue[0].CreateNextMoves(), []Chessboard{}
			}

			// Too many repeats are ending up in here
			boardMap[curBoard.Stringify()] = StringifySliceCB(curNextMoves)
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
	return nil
}

func getBoard(masterAddr string) GetBoardReply {
	client, err := rpc.DialHTTP("tcp", string(masterAddr))
	if err != nil {
		log.Fatalf("rpc.DialHTTP(Get_Board): %v", err)
	}
	var nothing Nothing
	var reply GetBoardReply
	if err = client.Call("MNode.Get_Board", nothing, &reply); err != nil {
		log.Fatalf("client.Call(Get_Board): %v", err)
	}
	if err = client.Close(); err != nil {
		log.Fatalf("client.Close(Get_Board) %v", err)
	}
	return reply
}
func returnBoards(masterAddr, myAddress string, curBoards map[string][]string) {
	client, err := rpc.DialHTTP("tcp", string(masterAddr))
	if err != nil {
		log.Fatalf("rpc.DialHTTP(Return_Board): %v", err)
	}
	var reply Nothing
	if err = client.Call("MNode.Return_Board", curBoards, &reply); err != nil {
		log.Fatalf("client.Call(Return_Board): %v", err)
	}
	if err = client.Close(); err != nil {
		log.Fatalf("client.Close(Return_Board) %v", err)
	}
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
