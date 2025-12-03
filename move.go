package chessProgram

func MovePiece(board *Board, sr, sc, tr, tc int) bool {
	piece := board.Get(sr, sc)
	target := board.Get(tr, tc)

	// cannot capture same color
	if isWhitePiece(piece) && isWhitePiece(target) {
		return false
	}
	if isBlackPiece(piece) && isBlackPiece(target) {
		return false
	}

	if !IsLegalMove(*board, sr, sc, tr, tc) {
		return false
	}

	// perform move
	board.Set(tr, tc, piece)
	board.Set(sr, sc, ".")

	return true
}
