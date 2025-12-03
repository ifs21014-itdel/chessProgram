package chessProgram

import (
	"bufio"
	"fmt"
	"os"
)

func PrintBoard(b Board) {
	fmt.Println()
	for r := 7; r >= 0; r-- {
		for c := 0; c < 8; c++ {
			fmt.Print(b[r][c], " ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func KingExists(b Board) bool {
	whiteExists := false
	blackExists := false

	for _, row := range b {
		for _, cell := range row {
			if cell == "K" {
				whiteExists = true
			}
			if cell == "k" {
				blackExists = true
			}
		}
	}

	return whiteExists && blackExists
}

func Game() {
	board := NewBoard()
	sc := bufio.NewScanner(os.Stdin)

	for {
		PrintBoard(board)
		fmt.Print("Move (ex: e2 e4): ")

		if !sc.Scan() {
			break
		}

		inp := sc.Text()

		sr, sc2, tr, tc, err := ParseInput(inp)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		if !MovePiece(&board, sr, sc2, tr, tc) {
			fmt.Println("Illegal move.")
			continue
		}

		if !KingExists(board) {
			fmt.Println("King captured! Game over.")
			break
		}
	}
}
