package chessProgram

type Board [8][8]string

func NewBoard() Board {
	var b Board

	// White pieces
	b[0] = [8]string{"R", "N", "B", "Q", "K", "B", "N", "R"}
	for i := 0; i < 8; i++ {
		b[1][i] = "P"
	}

	// Black pieces
	b[7] = [8]string{"r", "n", "b", "q", "k", "b", "n", "r"}
	for i := 0; i < 8; i++ {
		b[6][i] = "p"
	}

	// Fill empty squares
	for r := 2; r <= 5; r++ {
		for c := 0; c < 8; c++ {
			b[r][c] = "."
		}
	}

	return b
}

func (b *Board) Get(r, c int) string {
	return b[r][c]
}

func (b *Board) Set(r, c int, v string) {
	b[r][c] = v
}
