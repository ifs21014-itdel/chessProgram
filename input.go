package chessProgram

import (
	"errors"
	"strconv"
)

func ParsePosition(pos string) (int, int, error) {
	if len(pos) != 2 {
		return 0, 0, errors.New("invalid position format (ex: e2)")
	}

	col := int(pos[0] - 'a')
	row, err := strconv.Atoi(string(pos[1]))
	if err != nil {
		return 0, 0, errors.New("invalid row value")
	}

	return row - 1, col, nil
}

func ParseInput(input string) (int, int, int, int, error) {
	if len(input) != 5 {
		return 0, 0, 0, 0, errors.New("use format: e2 e4")
	}

	sr, sc, err := ParsePosition(input[:2])
	if err != nil {
		return 0, 0, 0, 0, err
	}

	tr, tc, err := ParsePosition(input[3:])
	if err != nil {
		return 0, 0, 0, 0, err
	}

	return sr, sc, tr, tc, nil
}
