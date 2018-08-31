package main

import (
	"errors"
)

func main() {
	puzzle := [][]int{{}, {}, {}, {}, {}, {}, {}, {}, {}}
	solveSdkuBrute(puzzle)
}

func solveSdkuBrute(p [][]int) (string, error) {
	return "", errors.New("unsovlable")
}
