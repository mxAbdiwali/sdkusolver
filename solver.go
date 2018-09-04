package solver

import (
	"errors"
)

func solveSdkuBrute(p [][]int) (string, error) {
	if len(p) == 0 || p == nil {
		return "", errors.New("Error: wrong parameter")
	}
	return "", nil
}

func isValidRow(p [][]int, curr int) bool {
	r := p[curr]
	if len(r) == 0 {
		return false
	}

	for i := 0; i < len(r)-1; i++ {
		for j := i + 1; j < len(r); j++ {
			if r[i] == r[j] {
				return false
			}
		}
	}

	return true

}

func isValidCol(p [][]int, curr int) bool {
	return true
}
