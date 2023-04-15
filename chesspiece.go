package main

// Color enumeration
const (
	Black int = iota
	White
	Neither
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
