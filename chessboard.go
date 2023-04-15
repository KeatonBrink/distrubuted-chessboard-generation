package main

import (
	"errors"
	"math"
)

type Chessboard struct {
	Board      [][]ChessPiece // Contains strings that can be deciphered into a specific piece
	hasCastled bool           // One of many logic checkers that will be needed
}

type Move struct {
	Row    int // Row of movement
	Column int //Column of movement
}

// Defining methods of Chessboard
// Move: Takes a starting and ending Move struct, and returns an error if the method could not
//
//	successfully move the piece for any reason
func (cb Chessboard) Move(starting, ending Move) error {
	// Todo
	canMove := cb.IsValidMove(starting, ending)
	if !canMove {
		return errors.New("Error, not a valid move")
	}
	// The piece will move regardless
	cb.Board[ending.Row][ending.Column] = cb.Board[starting.Row][starting.Column]
	cb.Board[starting.Row][starting.Column] = ChessPiece{Color: Neither, Name: Empty}
	// Check if castling, so the rook can be moved
	if cb.Board[ending.Row][ending.Column].Name == King && math.Abs(float64(starting.Column-ending.Column)) > 1 {
		if starting.Row == 0 {
			if ending.Column == 6 {
				cb.Board[0][5] = ChessPiece{Color: White, Name: Rook}
			} else {
				cb.Board[0][3] = ChessPiece{Color: White, Name: Rook}
			}
		} else {
			if ending.Column == 6 {
				cb.Board[0][5] = ChessPiece{Color: Black, Name: Rook}
			} else {
				cb.Board[0][3] = ChessPiece{Color: Black, Name: Rook}
			}
		}
		cb.hasCastled = true
	}
	// Todo: Implment en passant
	return nil
}

// IsValidMove: Takes a starting and ending move, and checks if the given piece can move.
//
//	Failure can occure due to check, checkmate, or moving on same color
func (cb Chessboard) IsValidMove(starting, ending Move) bool {
	// Check starting location is valid location (on board)
	if starting.Row < 0 || starting.Row > 7 || starting.Column < 0 || starting.Column > 7 {
		return false
	}
	// Check ending location is valid location (on board)
	if ending.Row < 0 || ending.Row > 7 || ending.Column < 0 || ending.Column > 7 {
		return false
	}
	// Check starting location contains a piece
	if cb.Board[starting.Row][starting.Column].Color == Neither {
		return false
	}
	// Check the piece being moved, doesn't move on itself
	if cb.Board[starting.Row][starting.Column].Color == cb.Board[ending.Row][ending.Column].Color {
		return false
	}

	// Start checking rules of each piece
	startingPiece := cb.Board[starting.Row][starting.Column]
	endingLocation := cb.Board[ending.Row][ending.Column]
	switch startingPiece.Name {
	// Rules for Pawn
	case Pawn:
		if startingPiece.Color == Black {
			// Can only move 2 spaces if on starting location
			deltaRow, deltaCol = FindDeltas(starting, ending)
			if starting.Column == ending.Column {
				if starting.Row 
			}
		} else {

		}

	}
	return true
}


func FindDeltas(starting, ending Move) (int, int) {
	deltaRow := starting.Row - ending.Row
}
