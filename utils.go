package chessProgram

func sign(x int) int {
	if x > 0 {
		return 1
	}
	if x < 0 {
		return -1
	}
	return 0
}

func isWhitePiece(p string) bool {
	return p >= "A" && p <= "Z"
}

func isBlackPiece(p string) bool {
	return p >= "a" && p <= "z"
}
