package distributedchessboardgeneration

import (
	"testing"
)

// Test movement of white pawn sideways from reset
func TestMoveWPawnSidways(t *testing.T) {
	var cb Chessboard
	cb.Reset()
	got := cb.IsValidMove(Move{Row: 1, Column: 1}, Move{Row: 1, Column: 2})
	want := false
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

// Test movement of white pawn forwards 1 from reset
func TestMoveWPawnForwards(t *testing.T) {
	var cb Chessboard
	cb.Reset()
	got := cb.IsValidMove(Move{Row: 1, Column: 1}, Move{Row: 2, Column: 1})
	want := true
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

// Test movement of white pawn forwards 2 from reset
func TestMoveWPawnForwards2(t *testing.T) {
	var cb Chessboard
	cb.Reset()
	got := cb.IsValidMove(Move{Row: 1, Column: 1}, Move{Row: 3, Column: 1})
	want := true
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

// Test movement of white rook forwards 1 from reset, but blocked
func TestMoveWRookBlocked1(t *testing.T) {
	var cb Chessboard
	cb.Reset()
	got := cb.IsValidMove(Move{Row: 0, Column: 0}, Move{Row: 1, Column: 0})
	want := false
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

// Test movement of white rook forwards 2 from reset, but blocked
func TestMoveWRookBlocked2(t *testing.T) {
	var cb Chessboard
	cb.Reset()
	got := cb.IsValidMove(Move{Row: 0, Column: 0}, Move{Row: 2, Column: 0})
	want := false
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

// Test movement of white rook forwards 6 from reset, but blocked
func TestMoveWRookBlocked3(t *testing.T) {
	var cb Chessboard
	cb.Reset()
	got := cb.IsValidMove(Move{Row: 0, Column: 0}, Move{Row: 6, Column: 0})
	want := false
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

// Test movement of white rook forwards 1 from reset, not blocked
func TestMoveWRook1(t *testing.T) {
	var cb Chessboard
	cb.Reset()
	cb.Board[1][0] = ChessPiece{Name: Empty, Color: Neither}
	got := cb.IsValidMove(Move{Row: 0, Column: 0}, Move{Row: 1, Column: 0})
	want := true
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

// Test movement of white rook forwards 6 from reset, not blocked
func TestMoveWRook2(t *testing.T) {
	var cb Chessboard
	cb.Reset()
	cb.Board[1][0] = ChessPiece{Name: Empty, Color: Neither}
	got := cb.IsValidMove(Move{Row: 0, Column: 0}, Move{Row: 6, Column: 0})
	want := true
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

// Test movement of white rook forwards 5 from reset, not blocked
func TestMoveWRook3(t *testing.T) {
	var cb Chessboard
	cb.Reset()
	cb.Board[1][0] = ChessPiece{Name: Empty, Color: Neither}
	got := cb.IsValidMove(Move{Row: 0, Column: 0}, Move{Row: 5, Column: 0})
	want := true
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

// Test check of white king by black queen
func TestIsCheckW(t *testing.T) {
	var cb Chessboard
	cb.Reset()
	cb.Board[1][3] = ChessPiece{Name: Empty, Color: Neither}
	cb.Board[4][0] = ChessPiece{Name: Queen, Color: Black}
	got := cb.IsCheck(White)
	want := true
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

// Test check of white king by black queen blocked
func TestIsCheckW2(t *testing.T) {
	var cb Chessboard
	cb.Reset()
	// cb.Board[1][3] = ChessPiece{Name: Empty, Color: Neither}
	cb.Board[4][0] = ChessPiece{Name: Queen, Color: Black}
	got := cb.IsCheck(White)
	want := false
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

// Test
