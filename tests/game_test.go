package tests

import (
	"chessProgram"
	"testing"
)

func TestKingExists(t *testing.T) {
	b := chessProgram.NewBoard()

	if !chessProgram.KingExists(b) {
		t.Errorf("king should exist at start")
	}
}

func TestKingCaptured(t *testing.T) {
	b := chessProgram.NewBoard()

	// Remove black king (simulate capture)
	b[7][4] = "."

	if chessProgram.KingExists(b) {
		t.Errorf("king should NOT exist after capture")
	}
}
