package tests

import (
	"chessProgram"
	"testing"
)

func TestBoardInitialization(t *testing.T) {
	b := chessProgram.NewBoard()

	// White pieces
	if b[0][0] != "R" || b[0][4] != "K" {
		t.Errorf("white pieces not initialized correctly")
	}

	// White pawns
	if b[1][0] != "P" || b[1][7] != "P" {
		t.Errorf("white pawns not initialized correctly")
	}

	// Black pieces
	if b[7][0] != "r" || b[7][4] != "k" {
		t.Errorf("black pieces not initialized correctly")
	}

	// Black pawns
	if b[6][0] != "p" || b[6][7] != "p" {
		t.Errorf("black pawns not initialized correctly")
	}

	// Middle should be empty
	for r := 2; r <= 5; r++ {
		for c := 0; c < 8; c++ {
			if b[r][c] != "." {
				t.Errorf("middle squares should be empty")
			}
		}
	}
}
