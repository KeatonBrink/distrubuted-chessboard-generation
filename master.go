package distributedchessboardgeneration

import (
	"log"
	"net/http"
	"net/rpc"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

type MNode struct {
	BoardsToDo     chan Chessboard
	BoardsToSQL    chan map[string][]string
	FinishedBoards map[string]bool
}

type GetBoardReply struct {
	Board  Chessboard
	IsDone bool
}

type Nothing struct{}

var Mutex sync.Mutex

func (n *MNode) Get_Board(nothing Nothing, reply *GetBoardReply) error {
	Mutex.Lock()
	defer Mutex.Unlock()
	(*reply).Board = <-n.BoardsToDo
	(*reply).IsDone = false
	log.Println("Tasking a Board:")
	(*reply).Board.PrintSelf()
	return nil
}

func (n *MNode) Return_Board(curBoards map[string][]string, reply *Nothing) error {
	log.Println("Recieving Finished Board")
	Mutex.Lock()
	defer Mutex.Unlock()
	for mstring := range curBoards {
		if _, ok := n.FinishedBoards[mstring]; !ok {
			n.FinishedBoards[mstring] = true
		} else {
			// Note, a rook/Knight on one side will have all the same moves as the other side
			//  And consequently throw this error.  Slight bug that leads to extra processing,
			//  but not detrimental to overall system design
			// log.Printf("Error in Returning computed board:\n%v", mstring)

			// Returning nil will avoid double entries in the SQL database
			return nil
		}
	}
	n.BoardsToSQL <- curBoards
	return nil
}

func masterNode(port string) error {
	myAddress := getLocalAddress() + ":" + port
	masterNode := &MNode{BoardsToDo: make(chan Chessboard, 10), BoardsToSQL: make(chan map[string][]string), FinishedBoards: make(map[string]bool)}
	var isFinishedChan chan Nothing
	go NextIterativeBoard(masterNode.BoardsToDo, isFinishedChan)
	// This will eventually be replaced with GRPC
	go InsertBoardAsSQL(masterNode.BoardsToSQL)
	rpc.Register(masterNode)
	rpc.HandleHTTP()
	go func() {
		if err := http.ListenAndServe(myAddress, nil); err != nil {
			log.Printf("Error in HTTP server for %s: %v", myAddress, err)
		}
	}()
	log.Println("Created rpc with address ", myAddress)
	<-isFinishedChan
	return nil
}

func InsertBoardAsSQL(inputMapChan <-chan map[string][]string) {
	// "database.db" is found here and in the makefile
	sourceDB, err := openDatabase("database.db")
	if err != nil {
		log.Fatalf("error InsertBoardAsSQL: Could not open sql file database.db: %v", err)
	}
	defer sourceDB.Close()
	indexCount := 1
	for curMap := range inputMapChan {
		log.Println("Writing map to disk")
		chessboardIndex := make(map[string]int)
		for curBoard := range curMap {
			_, err := sourceDB.Exec("INSERT INTO chessboards (chessboard) VALUES (?)", curBoard)
			if err != nil {
				log.Fatalf("error InsertBoardAsSQL: INSERT INTO chessboards: %v", err)
			}
			chessboardIndex[curBoard] = indexCount
			indexCount++
		}
		// Learned a lesson here, go purposefully randomizes the map
		// Note, errors are occuring because chessboards are not created for boards
		// with a number of pieces less than the starter.
		// As such, the foreign key has been removed for the time being
		for curBoard, nextBoards := range curMap {
			for _, nextBoard := range nextBoards {
				_, err = sourceDB.Exec("INSERT INTO move (from_id, to_id) VALUES (?, ?)", chessboardIndex[curBoard], chessboardIndex[nextBoard])
				if err != nil {
					log.Println(chessboardIndex[curBoard], curBoard)
					log.Println(chessboardIndex[nextBoard], nextBoard)
					log.Printf("error InsertBoardAsSQL: INSERT INTO move: %v", err)
				}
			}
		}
		log.Println("Finished writing map to disk")
	}
}
