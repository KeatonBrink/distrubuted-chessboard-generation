package main

import (
	"log"

	dbc "github.com/KeatonBrink/distrubuted-chessboard-generation"
)

func main() {
	dbc.Start()
	// var temp dbc.Chessboard
	// temp.Board = make([][]dbc.ChessPiece, 8)
	// for i := range temp.Board {
	// 	temp.Board[i] = make([]dbc.ChessPiece, 8)
	// }
	// temp.Board[3][0] = dbc.ChessPiece{Color: dbc.Black, Name: dbc.Rook}
	// // temp.Board[0][4] = dbc.ChessPiece{Color: dbc.Black, Name: dbc.King}
	// // temp.Board[7][4] = dbc.ChessPiece{Color: dbc.White, Name: dbc.King}
	// fmt.Println(temp.Stringify())
	// temp.PrintSelf()
}

func GenerateStarterBoards() {
	outputCBChan := make(chan dbc.Chessboard, 10)
	var isFinishedChan chan dbc.Nothing
	go dbc.NextIterativeBoard(outputCBChan, isFinishedChan)

	for val := range outputCBChan {
		val.PrintSelf()
		curCBs := val.CreateNextMoves()
		log.Println(len(curCBs))
		for _, curCB := range curCBs {
			curCB.PrintSelf()
			// time.Sleep(time.Second / 8)
		}
	}
}
