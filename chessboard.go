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
	// Check a move actually happens
	if starting.Row == ending.Row && starting.Column == ending.Column {
		return false
	}

	// Start checking rules of each piece
	startingPiece := cb.Board[starting.Row][starting.Column]
	endingLocation := cb.Board[ending.Row][ending.Column]
	deltaRow, deltaCol := FindDeltas(starting, ending)
	switch startingPiece.Name {
	// Rules for Pawn
	//  These rules have taken longer than anticipated
	case Pawn:
		if startingPiece.Color == Black {
			// If the piece moves Upwards
			if starting.Row <= ending.Row {
				return false
			}
			if deltaRow == 2 {
				if starting.Row != 6 {
					return false
				}
				// If the closest spot is not empty
				if cb.Board[starting.Row-1][starting.Column].Color != Neither {
					return false
				}
				// If two spots away is not empty
				if cb.Board[starting.Row-2][starting.Column].Color != Neither {
					return false
				}
			} else if deltaRow == 1 {
				if deltaCol > 1 {
					return false
				} else if deltaCol == 1 {
					if endingLocation.Color != White {
						return false
					}
				} else if deltaCol == 0 {
					if endingLocation.Color != Neither {
						return false
					}
				}
			} else {
				// Delta Row > 2
				return false
			}
		} else {
			// Color is White
			// If the piece moves downwards
			if starting.Row >= ending.Row {
				return false
			}
			if deltaRow == 2 {
				if starting.Row != 1 {
					return false
				}
				// If the closest spot is not empty
				if cb.Board[starting.Row+1][starting.Column].Color != Neither {
					return false
				}
				// If two spots away is not empty
				if cb.Board[starting.Row+2][starting.Column].Color != Neither {
					return false
				}
			} else if deltaRow == 1 {
				// Can only ever move one spot sideways
				if deltaCol > 1 {
					return false
				} else if deltaCol == 1 {
					//Can only move sideways if the spot is the opposite
					if endingLocation.Color != Black {
						return false
					}
				} else if deltaCol == 0 {
					if endingLocation.Color != Neither {
						return false
					}
				}
				// Delta Row > 2
			} else {
				return false
			}
		}
	// Rules for Rook
	case Rook:
		// Either Col or Row must not change
		if deltaCol != 0 && deltaRow != 0 {
			return false
		}
		// Go the 4 cardinal directions and check if a piece is in the way
		for up := starting.Row + 1; up < ending.Row; up++ {
			if cb.Board[up][starting.Column].Color != Neither {
				return false
			}
		}
		for down := starting.Row - 1; down > ending.Row; down-- {
			if cb.Board[down][starting.Column].Color != Neither {
				return false
			}
		}
		for left := starting.Column - 1; left > ending.Column; left-- {
			if cb.Board[left][starting.Column].Color != Neither {
				return false
			}
		}
		for right := starting.Column + 1; right < ending.Column; right++ {
			if cb.Board[right][starting.Column].Color != Neither {
				return false
			}
		}
	case Knight:
		// Only move 1 in a direction and 2 in the other
		if (deltaCol == 2 && deltaRow != 1) || (deltaRow == 2 && deltaCol != 1) {
			return false
		}
	case Bishop:
		if deltaCol != deltaRow {
			return false
		}
		if starting.Row < ending.Row {
			if starting.Column < ending.Column {
				for curRow, curCol := starting.Row+1, starting.Column-1; curRow < ending.Row; curRow, curCol = curRow+1, curCol-1 {
					if cb.Board[curRow][curCol].Color != Neither {
						return false
					}
				}
			} else {
				for curRow, curCol := starting.Row+1, starting.Column+1; curRow < ending.Row; curRow, curCol = curRow+1, curCol+1 {
					if cb.Board[curRow][curCol].Color != Neither {
						return false
					}
				}
			}
		} else {
			if starting.Column < ending.Column {
				for curRow, curCol := starting.Row-1, starting.Column-1; curRow < ending.Row; curRow, curCol = curRow-1, curCol-1 {
					if cb.Board[curRow][curCol].Color != Neither {
						return false
					}
				}
			} else {
				for curRow, curCol := starting.Row-1, starting.Column+1; curRow < ending.Row; curRow, curCol = curRow-1, curCol+1 {
					if cb.Board[curRow][curCol].Color != Neither {
						return false
					}
				}
			}
		}
	// Queen is a bishop and rook in one piece
	case Queen:
		if deltaCol == deltaRow {
			if starting.Row < ending.Row {
				if starting.Column < ending.Column {
					for curRow, curCol := starting.Row+1, starting.Column-1; curRow < ending.Row; curRow, curCol = curRow+1, curCol-1 {
						if cb.Board[curRow][curCol].Color != Neither {
							return false
						}
					}
				} else {
					for curRow, curCol := starting.Row+1, starting.Column+1; curRow < ending.Row; curRow, curCol = curRow+1, curCol+1 {
						if cb.Board[curRow][curCol].Color != Neither {
							return false
						}
					}
				}
			} else {
				if starting.Column < ending.Column {
					for curRow, curCol := starting.Row-1, starting.Column-1; curRow < ending.Row; curRow, curCol = curRow-1, curCol-1 {
						if cb.Board[curRow][curCol].Color != Neither {
							return false
						}
					}
				} else {
					for curRow, curCol := starting.Row-1, starting.Column+1; curRow < ending.Row; curRow, curCol = curRow-1, curCol+1 {
						if cb.Board[curRow][curCol].Color != Neither {
							return false
						}
					}
				}
			}
		} else if (deltaCol != 0 && deltaRow == 0) || (deltaCol == 0 && deltaRow != 0) {
			for up := starting.Row + 1; up < ending.Row; up++ {
				if cb.Board[up][starting.Column].Color != Neither {
					return false
				}
			}
			for down := starting.Row - 1; down > ending.Row; down-- {
				if cb.Board[down][starting.Column].Color != Neither {
					return false
				}
			}
			for left := starting.Column - 1; left > ending.Column; left-- {
				if cb.Board[left][starting.Column].Color != Neither {
					return false
				}
			}
			for right := starting.Column + 1; right < ending.Column; right++ {
				if cb.Board[right][starting.Column].Color != Neither {
					return false
				}
			}
		} else {
			return false
		}
	case King:
		if deltaCol > 1 || deltaRow > 1 {
			return false
		}
	default:
		return false
	}
	// Todo
	if cb.IsResultCheck(starting, ending) {
		return false
	}
	return true
}

func FindDeltas(starting, ending Move) (int, int) {
	deltaRow := int(math.Abs(float64(starting.Row - ending.Row)))
	deltaColumn := int(math.Abs(float64(starting.Column - ending.Column)))
	return deltaRow, deltaColumn
}

// Needs work
func (cb Chessboard) IsResultCheck(starting, ending Move) bool {
	return false
}
