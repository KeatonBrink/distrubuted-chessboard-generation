package distributedchessboardgeneration

import (
	"errors"
	"fmt"
	"log"
	"math"
)

/*
	chessboard.go contains the structs of Chessboard and Move, along with methods to implement valid moves.
	Note, the code is verbose potentially to a point of concern.  My current intention it create a working system and worry about optimization later.
		An additional bonus to the syntax should be easier debugging.
*/

type Chessboard struct {
	Board      [][]ChessPiece // Contains strings that can be deciphered into a specific piece
	HasCastled bool           // One of many logic checkers that will be needed
}

type Move struct {
	Row    int // Row of movement
	Column int //Column of movement
}

// Defining methods of Chessboard
// Move: Takes a starting and ending Move struct, and returns an error if the method could not
//
//	I am not positive I want this to be a method of itself
//
//	successfully move the piece for any reason
func (cb Chessboard) MoveSelf(starting, ending Move) error {
	canMove := cb.IsValidMove(starting, ending)
	if !canMove {
		return errors.New("isValidMove error, not a valid move")
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
		cb.HasCastled = true
	}
	// Todo: Implment en passant
	return nil
}

func MovePiece(cb Chessboard, starting, ending Move) (Chessboard, error) {
	canMove := cb.IsValidMove(starting, ending)
	if !canMove {
		return cb, errors.New("isValidMove error, not a valid move")
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
		cb.HasCastled = true
	}
	// Todo: Implment en passant
	return cb, nil
}

// Sets board to starting pieces
func (cb *Chessboard) Reset() error {
	cb.HasCastled = false
	cb.Board = make([][]ChessPiece, 8)
	for i := range cb.Board {
		cb.Board[i] = make([]ChessPiece, 8)
	}
	for curRow := 0; curRow <= 7; curRow++ {
		for curCol := 0; curCol <= 7; curCol++ {
			cb.Board[curRow][curCol].Name = Empty
			cb.Board[curRow][curCol].Color = Neither
		}
	}
	// Starting from bottom left to top right filling appropriate spots with pieces
	cb.Board[0][0] = ChessPiece{Color: White, Name: Rook}
	cb.Board[0][1] = ChessPiece{Color: White, Name: Knight}
	cb.Board[0][2] = ChessPiece{Color: White, Name: Bishop}
	cb.Board[0][3] = ChessPiece{Color: White, Name: Queen}
	cb.Board[0][4] = ChessPiece{Color: White, Name: King}
	cb.Board[0][5] = ChessPiece{Color: White, Name: Bishop}
	cb.Board[0][6] = ChessPiece{Color: White, Name: Knight}
	cb.Board[0][7] = ChessPiece{Color: White, Name: Rook}
	for curCol := 0; curCol <= 7; curCol++ {
		cb.Board[1][curCol] = ChessPiece{Color: White, Name: Pawn}
		cb.Board[6][curCol] = ChessPiece{Color: Black, Name: Pawn}
	}
	cb.Board[7][0] = ChessPiece{Color: Black, Name: Rook}
	cb.Board[7][1] = ChessPiece{Color: Black, Name: Knight}
	cb.Board[7][2] = ChessPiece{Color: Black, Name: Bishop}
	cb.Board[7][3] = ChessPiece{Color: Black, Name: Queen}
	cb.Board[7][4] = ChessPiece{Color: Black, Name: King}
	cb.Board[7][5] = ChessPiece{Color: Black, Name: Bishop}
	cb.Board[7][6] = ChessPiece{Color: Black, Name: Knight}
	cb.Board[7][7] = ChessPiece{Color: Black, Name: Rook}
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
			if cb.Board[starting.Row][left].Color != Neither {
				return false
			}
		}
		for right := starting.Column + 1; right < ending.Column; right++ {
			if cb.Board[starting.Row][right].Color != Neither {
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
				if cb.Board[starting.Row][left].Color != Neither {
					return false
				}
			}
			for right := starting.Column + 1; right < ending.Column; right++ {
				if cb.Board[starting.Row][right].Color != Neither {
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

// Needs more tests
func (cb Chessboard) IsResultCheck(starting, ending Move) bool {
	tempCB := cb
	tempCB.Board[ending.Row][ending.Column] = tempCB.Board[starting.Row][starting.Column]
	tempCB.Board[starting.Row][starting.Column] = ChessPiece{Color: Neither, Name: Empty}
	kingColor := tempCB.Board[starting.Row][starting.Column].Color
	return tempCB.IsCheck(kingColor)
}

func (cb Chessboard) IsCheck(kingColor int) bool {
	tempCB := cb
	var kingLocation Move
	for curRow := 0; curRow < 8; curRow++ {
		for curCol := 0; curCol < 8; curCol++ {
			if tempCB.Board[curRow][curCol].Color == kingColor && tempCB.Board[curRow][curCol].Name == King {
				kingLocation.Row = curRow
				kingLocation.Column = curCol
			}
		}
	}
	if kingColor == Black {
		// Check pawns
		if kingLocation.Row > 1 && ((kingLocation.Column > 0 && tempCB.Board[kingLocation.Row-1][kingLocation.Column-1].Name == Pawn && tempCB.Board[kingLocation.Row-1][kingLocation.Column-1].Color == White) ||
			(kingLocation.Column < 7 && tempCB.Board[kingLocation.Row-1][kingLocation.Column+1].Name == Pawn && tempCB.Board[kingLocation.Row-1][kingLocation.Column+1].Color == White)) {

			return true
		}
		// Check rook and 1/2 queen
		for up := kingLocation.Row + 1; up <= 7; up++ {
			if tempCB.Board[up][kingLocation.Column].Color == White {
				if tempCB.Board[up][kingLocation.Column].Name == Rook || tempCB.Board[up][kingLocation.Column].Name == Queen {
					return true
				} else {
					break
				}
			} else if tempCB.Board[up][kingLocation.Column].Color == Black {
				break
			}
		}
		for down := kingLocation.Row - 1; down >= 0; down-- {
			if tempCB.Board[down][kingLocation.Column].Color == White {
				if tempCB.Board[down][kingLocation.Column].Name == Rook || tempCB.Board[down][kingLocation.Column].Name == Queen {
					return true
				} else {
					break
				}
			} else if tempCB.Board[down][kingLocation.Column].Color == Black {
				break
			}
		}
		for left := kingLocation.Column - 1; left >= 0; left-- {
			if tempCB.Board[kingLocation.Row][left].Color == White {
				if tempCB.Board[kingLocation.Row][left].Name == Rook || tempCB.Board[kingLocation.Row][kingLocation.Column].Name == Queen {
					return true
				} else {
					break
				}
			} else if tempCB.Board[kingLocation.Row][left].Color == Black {
				break
			}
		}
		for right := kingLocation.Column + 1; right <= 7; right++ {
			if tempCB.Board[kingLocation.Row][right].Color == White {
				if tempCB.Board[kingLocation.Row][right].Name == Rook || tempCB.Board[kingLocation.Row][right].Name == Queen {
					return true
				} else {
					break
				}
			} else if tempCB.Board[kingLocation.Row][right].Color == Black {
				break
			}
		}
		// Bishop and last 1/2 of Queen
		for curRow, curCol := kingLocation.Row+1, kingLocation.Column-1; curRow <= 7 && curCol >= 0; curRow, curCol = curRow+1, curCol-1 {
			if tempCB.Board[curRow][curCol].Color == White {
				if tempCB.Board[curRow][curCol].Name == Bishop || tempCB.Board[curRow][curCol].Name == Queen {
					return true
				} else {
					break
				}
			} else if tempCB.Board[curRow][curCol].Color == Black {
				break
			}
		}
		for curRow, curCol := kingLocation.Row+1, kingLocation.Column+1; curRow <= 7 && curCol <= 7; curRow, curCol = curRow+1, curCol+1 {
			if tempCB.Board[curRow][curCol].Color == White {
				if tempCB.Board[curRow][curCol].Name == Bishop || tempCB.Board[curRow][curCol].Name == Queen {
					return true
				} else {
					break
				}
			} else if tempCB.Board[curRow][curCol].Color == Black {
				break
			}
		}
		for curRow, curCol := kingLocation.Row-1, kingLocation.Column+1; curRow >= 0 && curCol <= 7; curRow, curCol = curRow-1, curCol+1 {
			if tempCB.Board[curRow][curCol].Color == White {
				if tempCB.Board[curRow][curCol].Name == Bishop || tempCB.Board[curRow][curCol].Name == Queen {
					return true
				} else {
					break
				}
			} else if tempCB.Board[curRow][curCol].Color == Black {
				break
			}
		}
		for curRow, curCol := kingLocation.Row-1, kingLocation.Column-1; curRow >= 0 && curCol >= 0; curRow, curCol = curRow-1, curCol-1 {
			if tempCB.Board[curRow][curCol].Color == White {
				if tempCB.Board[curRow][curCol].Name == Bishop || tempCB.Board[curRow][curCol].Name == Queen {
					return true
				} else {
					break
				}
			} else if tempCB.Board[curRow][curCol].Color == Black {
				break
			}
		}
		// Knight
		if kingLocation.Row < 7 && kingLocation.Column > 1 && tempCB.Board[kingLocation.Row+1][kingLocation.Column-2].Color == White && tempCB.Board[kingLocation.Row+1][kingLocation.Column-2].Name == Knight {
			return true
		}
		if kingLocation.Row < 6 && kingLocation.Column > 0 && tempCB.Board[kingLocation.Row+2][kingLocation.Column-1].Color == White && tempCB.Board[kingLocation.Row+2][kingLocation.Column-1].Name == Knight {
			return true
		}
		if kingLocation.Row < 6 && kingLocation.Column < 7 && tempCB.Board[kingLocation.Row+2][kingLocation.Column+1].Color == White && tempCB.Board[kingLocation.Row+2][kingLocation.Column+1].Name == Knight {
			return true
		}
		if kingLocation.Row < 7 && kingLocation.Column < 6 && tempCB.Board[kingLocation.Row+1][kingLocation.Column+2].Color == White && tempCB.Board[kingLocation.Row+1][kingLocation.Column+2].Name == Knight {
			return true
		}
		if kingLocation.Row > 0 && kingLocation.Column < 6 && tempCB.Board[kingLocation.Row-1][kingLocation.Column+2].Color == White && tempCB.Board[kingLocation.Row-1][kingLocation.Column+2].Name == Knight {
			return true
		}
		if kingLocation.Row > 1 && kingLocation.Column < 7 && tempCB.Board[kingLocation.Row-2][kingLocation.Column+1].Color == White && tempCB.Board[kingLocation.Row-2][kingLocation.Column+1].Name == Knight {
			return true
		}
		if kingLocation.Row > 1 && kingLocation.Column > 0 && tempCB.Board[kingLocation.Row-2][kingLocation.Column-1].Color == White && tempCB.Board[kingLocation.Row-2][kingLocation.Column-1].Name == Knight {
			return true
		}
		if kingLocation.Row > 0 && kingLocation.Column > 1 && tempCB.Board[kingLocation.Row-1][kingLocation.Column-2].Color == White && tempCB.Board[kingLocation.Row-1][kingLocation.Column-2].Name == Knight {
			return true
		}
		// King
		if kingLocation.Row < 7 {
			if kingLocation.Column > 0 && tempCB.Board[kingLocation.Row+1][kingLocation.Column-1].Color == White && tempCB.Board[kingLocation.Row+1][kingLocation.Column-1].Name == King {
				return true
			}
			if tempCB.Board[kingLocation.Row+1][kingLocation.Column].Color == White && tempCB.Board[kingLocation.Row+1][kingLocation.Column].Name == King {
				return true
			}
			if kingLocation.Column < 7 && tempCB.Board[kingLocation.Row+1][kingLocation.Column+1].Color == White && tempCB.Board[kingLocation.Row+1][kingLocation.Column+1].Name == King {
				return true
			}
		}
		if kingLocation.Column < 7 && tempCB.Board[kingLocation.Row][kingLocation.Column+1].Color == White && tempCB.Board[kingLocation.Row][kingLocation.Column+1].Name == King {
			return true
		}
		if kingLocation.Row > 0 {
			if kingLocation.Column < 7 && tempCB.Board[kingLocation.Row-1][kingLocation.Column+1].Color == White && tempCB.Board[kingLocation.Row-1][kingLocation.Column+1].Name == King {
				return true
			}
			if tempCB.Board[kingLocation.Row-1][kingLocation.Column].Color == White && tempCB.Board[kingLocation.Row-1][kingLocation.Column].Name == King {
				return true
			}
			if kingLocation.Column > 0 && tempCB.Board[kingLocation.Row-1][kingLocation.Column-1].Color == White && tempCB.Board[kingLocation.Row-1][kingLocation.Column-1].Name == King {
				return true
			}
		}
		if kingLocation.Column > 0 && tempCB.Board[kingLocation.Row][kingLocation.Column-1].Color == White && tempCB.Board[kingLocation.Row][kingLocation.Column-1].Name == King {
			return true
		}
		// In case the King is white
	} else {
		if kingLocation.Row < 6 && ((kingLocation.Column > 0 && tempCB.Board[kingLocation.Row+1][kingLocation.Column-1].Name == Pawn && tempCB.Board[kingLocation.Row+1][kingLocation.Column-1].Color == Black) ||
			(kingLocation.Column < 7 && tempCB.Board[kingLocation.Row+1][kingLocation.Column+1].Name == Pawn && tempCB.Board[kingLocation.Row+1][kingLocation.Column+1].Color == Black)) {
			return true
		}
		// Check rook and 1/2 queen
		for up := kingLocation.Row + 1; up <= 7; up++ {
			if tempCB.Board[up][kingLocation.Column].Color == Black {
				if tempCB.Board[up][kingLocation.Column].Name == Rook || tempCB.Board[up][kingLocation.Column].Name == Queen {
					return true
				} else {
					break
				}
			} else if tempCB.Board[up][kingLocation.Column].Color == White {
				break
			}
		}
		for down := kingLocation.Row - 1; down >= 0; down-- {
			if tempCB.Board[down][kingLocation.Column].Color == Black {
				if tempCB.Board[down][kingLocation.Column].Name == Rook || tempCB.Board[down][kingLocation.Column].Name == Queen {
					return true
				} else {
					break
				}
			} else if tempCB.Board[down][kingLocation.Column].Color == White {
				break
			}
		}
		for left := kingLocation.Column - 1; left >= 0; left-- {
			if tempCB.Board[kingLocation.Row][left].Color == Black {
				if tempCB.Board[kingLocation.Row][left].Name == Rook || tempCB.Board[kingLocation.Row][kingLocation.Column].Name == Queen {
					return true
				} else {
					break
				}
			} else if tempCB.Board[kingLocation.Row][left].Color == White {
				break
			}
		}
		for right := kingLocation.Column + 1; right <= 7; right++ {
			if tempCB.Board[kingLocation.Row][right].Color == Black {
				if tempCB.Board[kingLocation.Row][right].Name == Rook || tempCB.Board[kingLocation.Row][right].Name == Queen {
					return true
				} else {
					break
				}
			} else if tempCB.Board[kingLocation.Row][right].Color == White {
				break
			}
		}
		// Bishop and last 1/2 of Queen
		for curRow, curCol := kingLocation.Row+1, kingLocation.Column-1; curRow <= 7 && curCol >= 0; curRow, curCol = curRow+1, curCol-1 {
			if tempCB.Board[curRow][curCol].Color == Black {
				if tempCB.Board[curRow][curCol].Name == Bishop || tempCB.Board[curRow][curCol].Name == Queen {
					return true
				} else {
					break
				}
			} else if tempCB.Board[curRow][curCol].Color == White {
				break
			}
		}
		for curRow, curCol := kingLocation.Row+1, kingLocation.Column+1; curRow <= 7 && curCol <= 7; curRow, curCol = curRow+1, curCol+1 {
			if tempCB.Board[curRow][curCol].Color == Black {
				if tempCB.Board[curRow][curCol].Name == Bishop || tempCB.Board[curRow][curCol].Name == Queen {
					return true
				} else {
					break
				}
			} else if tempCB.Board[curRow][curCol].Color == White {
				break
			}
		}
		for curRow, curCol := kingLocation.Row-1, kingLocation.Column+1; curRow >= 0 && curCol <= 7; curRow, curCol = curRow-1, curCol+1 {
			if tempCB.Board[curRow][curCol].Color == Black {
				if tempCB.Board[curRow][curCol].Name == Bishop || tempCB.Board[curRow][curCol].Name == Queen {
					return true
				} else {
					break
				}
			} else if tempCB.Board[curRow][curCol].Color == White {
				break
			}
		}
		for curRow, curCol := kingLocation.Row-1, kingLocation.Column-1; curRow >= 0 && curCol >= 0; curRow, curCol = curRow-1, curCol-1 {
			if tempCB.Board[curRow][curCol].Color == Black {
				if tempCB.Board[curRow][curCol].Name == Bishop || tempCB.Board[curRow][curCol].Name == Queen {
					return true
				} else {
					break
				}
			} else if tempCB.Board[curRow][curCol].Color == White {
				break
			}
		}
		// Knight
		if kingLocation.Row < 7 && kingLocation.Column > 1 && tempCB.Board[kingLocation.Row+1][kingLocation.Column-2].Color == Black && tempCB.Board[kingLocation.Row+1][kingLocation.Column-2].Name == Knight {
			return true
		}
		if kingLocation.Row < 6 && kingLocation.Column > 0 && tempCB.Board[kingLocation.Row+2][kingLocation.Column-1].Color == Black && tempCB.Board[kingLocation.Row+2][kingLocation.Column-1].Name == Knight {
			return true
		}
		if kingLocation.Row < 6 && kingLocation.Column < 7 && tempCB.Board[kingLocation.Row+2][kingLocation.Column+1].Color == Black && tempCB.Board[kingLocation.Row+2][kingLocation.Column+1].Name == Knight {
			return true
		}
		if kingLocation.Row < 7 && kingLocation.Column < 6 && tempCB.Board[kingLocation.Row+1][kingLocation.Column+2].Color == Black && tempCB.Board[kingLocation.Row+1][kingLocation.Column+2].Name == Knight {
			return true
		}
		if kingLocation.Row > 0 && kingLocation.Column < 6 && tempCB.Board[kingLocation.Row-1][kingLocation.Column+2].Color == Black && tempCB.Board[kingLocation.Row-1][kingLocation.Column+2].Name == Knight {
			return true
		}
		if kingLocation.Row > 1 && kingLocation.Column < 7 && tempCB.Board[kingLocation.Row-2][kingLocation.Column+1].Color == Black && tempCB.Board[kingLocation.Row-2][kingLocation.Column+1].Name == Knight {
			return true
		}
		if kingLocation.Row > 1 && kingLocation.Column > 0 && tempCB.Board[kingLocation.Row-2][kingLocation.Column-1].Color == Black && tempCB.Board[kingLocation.Row-2][kingLocation.Column-1].Name == Knight {
			return true
		}
		if kingLocation.Row > 0 && kingLocation.Column > 1 && tempCB.Board[kingLocation.Row-1][kingLocation.Column-2].Color == Black && tempCB.Board[kingLocation.Row-1][kingLocation.Column-2].Name == Knight {
			return true
		}
		// King
		if kingLocation.Row < 7 {
			if kingLocation.Column > 0 && tempCB.Board[kingLocation.Row+1][kingLocation.Column-1].Color == Black && tempCB.Board[kingLocation.Row+1][kingLocation.Column-1].Name == King {
				return true
			}
			if tempCB.Board[kingLocation.Row+1][kingLocation.Column].Color == Black && tempCB.Board[kingLocation.Row+1][kingLocation.Column].Name == King {
				return true
			}
			if kingLocation.Column < 7 && tempCB.Board[kingLocation.Row+1][kingLocation.Column+1].Color == Black && tempCB.Board[kingLocation.Row+1][kingLocation.Column+1].Name == King {
				return true
			}
		}
		if kingLocation.Column < 7 && tempCB.Board[kingLocation.Row][kingLocation.Column+1].Color == Black && tempCB.Board[kingLocation.Row][kingLocation.Column+1].Name == King {
			return true
		}
		if kingLocation.Row > 0 {
			if kingLocation.Column < 7 && tempCB.Board[kingLocation.Row-1][kingLocation.Column+1].Color == Black && tempCB.Board[kingLocation.Row-1][kingLocation.Column+1].Name == King {
				return true
			}
			if tempCB.Board[kingLocation.Row-1][kingLocation.Column].Color == Black && tempCB.Board[kingLocation.Row-1][kingLocation.Column].Name == King {
				return true
			}
			if kingLocation.Column > 0 && tempCB.Board[kingLocation.Row-1][kingLocation.Column-1].Color == Black && tempCB.Board[kingLocation.Row-1][kingLocation.Column-1].Name == King {
				return true
			}
		}
		if kingLocation.Column > 0 && tempCB.Board[kingLocation.Row][kingLocation.Column-1].Color == Black && tempCB.Board[kingLocation.Row][kingLocation.Column-1].Name == King {
			return true
		}
	}
	return false
}

// Prints board in a semi-readable fashion
func (cb Chessboard) PrintSelf() {
	for curRow := 7; curRow >= 0; curRow-- {
		fmt.Printf("_%v_", curRow)
		for curCol := 0; curCol <= 7; curCol++ {
			curPiece := cb.Board[curRow][curCol]
			if curPiece.Color == Neither {
				fmt.Printf(" ___ ")
				continue
			}
			fmt.Printf(" ")
			if curPiece.Color == White {
				fmt.Printf("W")
			} else {
				fmt.Printf("B")
			}
			switch curPiece.Name {
			case Pawn:
				fmt.Print("Pa")
			case Rook:
				fmt.Print("Ro")
			case Knight:
				fmt.Print("Kn")
			case Bishop:
				fmt.Print("Bi")
			case Queen:
				fmt.Print("Qu")
			case King:
				fmt.Print("Ki")
			default:
				log.Fatal("PrintSelf error, Piece with unknown name")
			}
			fmt.Printf(" ")

		}
		fmt.Println()
		fmt.Println()
	}
	fmt.Println("     0    1    2    3    4    5    6    7   ")
}

// Generate all next moves, utilized by worker
func (cb Chessboard) CreateNextMoves() []Chessboard {
	allPieceLocations := cb.FindAllPieceLocations()
	var nextBoards []Chessboard
	// For each location on the board, try to create new boards
	for curRow := 0; curRow <= 7; curRow++ {
		for curCol := 0; curCol <= 7; curCol++ {
			// At each location, test movement of all pieces to that location
			for _, piece := range allPieceLocations {
				testCB, err := MovePiece(cb, piece, Move{Row: curRow, Column: curCol})
				// Note, an error occurs when the move is invalid, and should be thrown away
				if err == nil {
					nextBoards = append(nextBoards, testCB)
				}
			}
		}
	}
	return nextBoards
}

// Takes a ChessBoard, and makes a string.  Will be used for mapping completed boards and prevent double execution
func (cb Chessboard) Stringify() string {
	stringCB := ""
	for curRow := 0; curRow <= 7; curRow++ {
		for curCol := 0; curCol <= 7; curCol++ {
			stringCB += cb.Board[curRow][curCol].Stringify()
		}
	}
	return stringCB
}

// Takes a ChessPiece and makes it a string
func (p ChessPiece) Stringify() string {
	curPieceString := ""
	if p.Color == Neither {
		return "_"
	}
	if p.Color == Black {
		curPieceString += "B"
	} else if p.Color == White {
		curPieceString += "W"
	}

	switch p.Name {
	case Pawn:
		curPieceString += "Pa"
	case Rook:
		curPieceString += "Ro"
	case Knight:
		curPieceString += "Kn"
	case Bishop:
		curPieceString += "Bi"
	case Queen:
		curPieceString += "Qu"
	case King:
		curPieceString += "Ki"
	default:
		log.Fatal("PrintSelf error, Piece with unknown name")
	}
	return curPieceString
}

func (cb Chessboard) FindAllPieceLocations() []Move {
	var allPieceLocations []Move
	for curRow := 0; curRow <= 7; curRow++ {
		for curCol := 0; curCol <= 7; curCol++ {
			if cb.Board[curRow][curRow].Color != Neither {
				allPieceLocations = append(allPieceLocations, Move{Row: curRow, Column: curCol})
			}
		}
	}
	return allPieceLocations
}

// Need an IsCheckmate clause
// And built on top is an IsStalemate

// Utilized by Master
func GenerateBoards() {
	outputCBChan := make(chan Chessboard, 10)
	go NextIterativeBoard(outputCBChan)
	for val := range outputCBChan {
		log.Println(val.Stringify())
		val.PrintSelf()
		// time.Sleep(time.Second * 1)
	}
}

// NextIterativeBoard takes in a board and essentially finds the next iteration of the board with new pieces
func NextIterativeBoard(outputCBChan chan<- Chessboard) {
	bitBoardChan := make(chan int, 20)
	go GetAll30PowerSet(bitBoardChan)

	for bitBoard := range bitBoardChan {
		outputCBChan <- DecodeBitBoard(bitBoard)
	}
}
