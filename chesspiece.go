package distributedchessboardgeneration

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
