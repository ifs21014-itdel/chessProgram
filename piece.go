package chessProgram

import "math"

func IsLegalMove(board Board, sr, sc, tr, tc int) bool {
	piece := board[sr][sc]
	target := board[tr][tc]

	if piece == "." {
		return false
	}

	dr := tr - sr
	dc := tc - sc

	switch piece {

	// ============================
	// WHITE PAWN
	// ============================
	case "P":
		// 1 step forward
		if dr == 1 && dc == 0 && target == "." {
			return true
		}

		// 2 steps from initial
		if sr == 1 && dr == 2 && dc == 0 && target == "." && board[sr+1][sc] == "." {
			return true
		}

		// capture diagonal
		if dr == 1 && math.Abs(float64(dc)) == 1 && target != "." && isBlackPiece(target) {
			return true
		}

		return false

	// ============================
	// BLACK PAWN
	// ============================
	case "p":
		if dr == -1 && dc == 0 && target == "." {
			return true
		}
		if sr == 6 && dr == -2 && dc == 0 && target == "." && board[sr-1][sc] == "." {
			return true
		}
		if dr == -1 && math.Abs(float64(dc)) == 1 && target != "." && isWhitePiece(target) {
			return true
		}
		return false

	// ============================
	// ROOK
	// ============================
	case "R", "r":
		if dr == 0 || dc == 0 {
			return isPathClear(board, sr, sc, tr, tc)
		}
		return false

	// ============================
	// BISHOP
	// ============================
	case "B", "b":
		if math.Abs(float64(dr)) == math.Abs(float64(dc)) {
			return isPathClear(board, sr, sc, tr, tc)
		}
		return false

	// ============================
	// QUEEN
	// ============================
	case "Q", "q":
		if dr == 0 || dc == 0 || math.Abs(float64(dr)) == math.Abs(float64(dc)) {
			return isPathClear(board, sr, sc, tr, tc)
		}
		return false

	// ============================
	// KNIGHT
	// ============================
	case "N", "n":
		return (math.Abs(float64(dr)) == 2 && math.Abs(float64(dc)) == 1) ||
			(math.Abs(float64(dr)) == 1 && math.Abs(float64(dc)) == 2)

	// ============================
	// KING
	// ============================
	case "K", "k":
		return math.Abs(float64(dr)) <= 1 && math.Abs(float64(dc)) <= 1
	}

	return false
}

func isPathClear(board Board, sr, sc, tr, tc int) bool {
	dr := sign(tr - sr)
	dc := sign(tc - sc)

	r := sr + dr
	c := sc + dc

	for r != tr || c != tc {
		if board[r][c] != "." {
			return false
		}
		r += dr
		c += dc
	}
	return true
}
