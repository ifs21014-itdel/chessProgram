package tests

import (
	"chessProgram"
	"testing"
)

func TestPawnMoveOneStep(t *testing.T) {
	b := chessProgram.NewBoard()

	if !chessProgram.IsLegalMove(b, 1, 0, 2, 0) {
		t.Errorf("white pawn should move 1 step")
	}
}

func TestPawnMoveTwoSteps(t *testing.T) {
	b := chessProgram.NewBoard()

	if !chessProgram.IsLegalMove(b, 1, 1, 3, 1) {
		t.Errorf("white pawn should move 2 steps on first move")
	}
}

func TestPawnBlockedTwoSteps(t *testing.T) {
	b := chessProgram.NewBoard()
	b[2][1] = "X" // block the pawn

	if chessProgram.IsLegalMove(b, 1, 1, 3, 1) {
		t.Errorf("pawn should NOT be able to move 2 steps when blocked")
	}
}

func TestPawnCapture(t *testing.T) {
	b := chessProgram.NewBoard()
	b[2][1] = "p" // black pawn in diagonal

	if !chessProgram.IsLegalMove(b, 1, 0, 2, 1) {
		t.Errorf("pawn should capture diagonally")
	}
}

func TestKnightJump(t *testing.T) {
	b := chessProgram.NewBoard()

	if !chessProgram.IsLegalMove(b, 0, 1, 2, 2) {
		t.Errorf("knight should jump in L shape")
	}
}

func TestRookBlocked(t *testing.T) {
	b := chessProgram.NewBoard()

	if chessProgram.IsLegalMove(b, 0, 0, 3, 0) {
		t.Errorf("rook cannot pass through pieces")
	}
}

func TestQueenDiagonal(t *testing.T) {
	b := chessProgram.NewBoard()

	// clear diagonal path from (0,3) -> (4,7)
	b[1][4] = "."
	b[2][5] = "."
	b[3][6] = "."

	if !chessProgram.IsLegalMove(b, 0, 3, 4, 7) {
		t.Errorf("queen should move diagonally")
	}
}

func TestBishopPathBlocked(t *testing.T) {
	b := chessProgram.NewBoard()

	if chessProgram.IsLegalMove(b, 0, 2, 3, 5) {
		t.Errorf("bishop path is blocked by pawn")
	}
}
