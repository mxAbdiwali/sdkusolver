package parser

import (
	"strings"
	"testing"
)

/* An import path is a pattern if it includes one or more "..." wildcards,
each of which can match any string, including the empty string and strings containing slashes.
Such a pattern expands to all package directories found in the GOPATH trees with names matching the patterns.
As a special case, x/... matches x as well as x's subdirectories.
For example, net/... expands to net and packages in its subdirectories. */

var sampletxt = "1 * 3 * * 6 * 8 *\n" +
	"* 5 * * 8 * 1 2 *\n" +
	"7 * 9 1 * 3 * 5 6\n" +
	"* 3 * * 6 7 * 9 *\n" +
	"5 * 7 8 * * * 3 *\n" +
	"8 * 1 * 3 * 5 * 7\n" +
	"* 4 * * 7 8 * 1 *\n" +
	"6 * 8 * * 2 * 4 *\n" +
	"* 1 2 * 4 5 * 7 8"
var samplearr = [][]int{{1, 0, 3, 0, 0, 6, 0, 8, 0}, {0, 5, 0, 0, 8, 0, 1, 2, 0}, {7, 0, 9, 1, 0, 3, 0, 5, 6},
	{0, 3, 0, 0, 6, 7, 0, 9, 0}, {5, 0, 7, 8, 0, 0, 0, 3, 0}, {8, 0, 1, 0, 3, 0, 5, 0, 7},
	{0, 4, 0, 0, 7, 8, 0, 1, 0}, {6, 0, 8, 0, 0, 2, 0, 4, 0}, {0, 1, 2, 0, 4, 5, 0, 7, 8}}

func Equal(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, x := range a {
		for j, y := range x {
			if y != b[i][j] {
				return false
			}
		}
	}
	return true
}
func TestParseEmptyFile(t *testing.T) {
	var expected = "Wrong parameter: nill or empty"
	_, err := parse("")
	if err.Error() != expected {
		t.Fatalf("expected to throw %s, but instead %s", expected, err.Error())
	}
}
func TestParseGetAsText(t *testing.T) {
	t.SkipNow()
	var plaintext, err = parseGetText("../sample.txt")
	if err != nil {
		t.Fatalf("%s", err)
	}

	if strings.EqualFold(plaintext, sampletxt) == false {
		t.Fatalf("%s != %s", plaintext, sampletxt)
	}

}

func TestParseGetAsArray(t *testing.T) {
	var soduku, err = parse("../sample.txt")
	if err != nil {
		t.Fatalf("%s", err)
	}

	if len(soduku) != len(samplearr) {
		t.Fatalf("wrong error not same length")
	}

	if !Equal(samplearr, soduku) {
		//t.Fatalf("expected %d , but found %d", samplearr, soduku)
		t.Fatalf("expected %d , but found %d", samplearr, soduku)
	}

}
