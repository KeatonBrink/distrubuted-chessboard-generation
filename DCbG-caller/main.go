package main

import (
	"log"

	dbc "github.com/KeatonBrink/distrubuted-chessboard-generation"
)

func main() {
	dbc.Start()
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
