package solver

import (
	"testing"
)

var sudoku = "1 * 3 * * 6 * 8 *\n" +
	"* 5 * * 8 * 1 2 *\n" +
	"7 * 9 1 * 3 * 5 6\n" +
	"* 3 * * 6 7 * 9 *\n" +
	"5 * 7 8 * * * 3 *\n" +
	"8 * 1 * 3 * 5 * 7\n" +
	"* 4 * * 7 8 * 1 *\n" +
	"6 * 8 * * 2 * 4 *\n" +
	"* 1 2 * 4 5 * 7 8"

var p = [][]int{{1, 0, 3, 0, 0, 6, 0, 8, 0}, {0, 5, 0, 0, 8, 0, 1, 2, 0}, {7, 0, 9, 1, 0, 3, 0, 5, 6},
	{0, 3, 0, 0, 6, 7, 0, 9, 0}, {5, 0, 7, 8, 0, 0, 0, 3, 0}, {8, 0, 1, 0, 3, 0, 5, 0, 7},
	{0, 4, 0, 0, 7, 8, 0, 1, 0}, {6, 0, 8, 0, 0, 2, 0, 4, 0}, {0, 1, 2, 0, 4, 5, 0, 7, 8}}

var solution = "1 2 3 4 5 6 7 8 9\n" +
	"4 5 6 7 8 9 1 2 3 \n" +
	"7 8 9 1 2 3 4 5 6\n" +
	"2 3 4 5 6 7 8 9 1\n" +
	"5 6 7 8 9 1 2 3 4\n" +
	"8 9 1 2 3 4 5 6 7\n" +
	"3 4 5 6 7 8 9 1 2\n" +
	"6 7 8 9 1 2 3 4 5\n" +
	"9 1 2 3 4 5 6 7 8"

type Test struct {
	in  [][]int
	out string
}

var tests = []Test{
	{[][]int{}, "Error: wrong parameter"},
	{nil, "Error: wrong paramete"},
	{nil, "Error: wrong paramete"},
}

/* You write a test by creating a file with a name ending in _test.go
that contains functions named TestXXX with signature
func (t *testing.T). The test framework runs each such function;
if the function calls a failure function such as t.Error or t.Fail,
the test is considered to have failed.  */

func TestSolve(t *testing.T) {
	for i, test := range tests {
		s, err := solveSdkuBrute(test.in)

		if err != nil {
			if err.Error() != test.out {
				t.Errorf("#%d: solveSdkuBrute(%d) = %s ; want %s", i, test.in, err.Error(), test.out)
			}
			break
		}

		if s != test.out {
			t.Errorf("#%d: solveSdkuBrute(%d)= %s; want %s", i, test.in, s, test.out)
		}
	}
}

func TestRowConstraintEmpty(t *testing.T) {
	p = [][]int{{}, {}, {}, {}, {}, {}, {}, {}, {}}
	var b = isValidRow(p, 1)
	if b != false {
		t.Errorf("solveSdkuBrute(%d)= %t; want %t", p, b, true)
	}
}

func TestRowConstraintInvalid(t *testing.T) {
	p = [][]int{{}, {1, 2, 3, 4, 8, 9, 7, 4, 5}, {}, {}, {}, {}, {}, {}, {}}
	var b = isValidRow(p, 1)
	if b != false {
		t.Errorf("solveSdkuBrute(%d)= %t; want %t", p, b, false)
	}
}

func TestRowConstraintValid(t *testing.T) {
	p = [][]int{{}, {1, 2, 3, 4, 8, 9, 7, 6, 5}, {}, {}, {}, {}, {}, {}, {}}
	var b = isValidRow(p, 1)

	if b != true {
		t.Errorf("solveSdkuBrute(%d)= %t; want %t", p, b, true)
	}
}

func TestColConstraindEmpty(t *testing.T) {
	p = [][]int{{}, {1, 2, 3, 4, 8, 9, 7, 6, 5}, {}, {}, {}, {}, {}, {}, {}}
	var b = isValidCol(p, 1)

	if b != false {
		t.Errorf("solveSdkuBrute(%d)= %t; want %t", p, b, true)
	}
}
