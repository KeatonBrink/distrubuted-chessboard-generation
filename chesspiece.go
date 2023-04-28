package distributedchessboardgeneration

import "log"

// Color enumeration
const (
	Neither int = iota
	Black
	White
)

// Piece enumeration
const (
	Empty int = iota
	Pawn
	Rook
	Knight
	Bishop
	Queen
	King
)

type ChessPiece struct {
	Color int
	Name  int
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

func CPieceify(cpString string) ChessPiece {
	retPiece := ChessPiece{}
	if cpString[0] == 'B' {
		retPiece.Color = Black
	} else {
		retPiece.Color = White
	}
	switch cpString[1:] {
	case "Ro":
		retPiece.Name = Rook
	case "Kn":
		retPiece.Name = Knight
	case "Bi":
		retPiece.Name = Bishop
	case "Qu":
		retPiece.Name = Queen
	case "Ki":
		retPiece.Name = King
	case "Pa":
		retPiece.Name = Pawn
	default:
		log.Fatal("Error in CPiecify, invalid name ", cpString[1:])
	}
	return retPiece
}
