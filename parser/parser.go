package parser

import (
	"errors"
	"io/ioutil"
	"strconv"
	"strings"
)

func readline(line string) ([]int, error) {
	sliceint := []int{}
	if line == "" {
		return nil, errors.New("line parsing error")
	}

	str := strings.Replace(line, "*", "0", 9)
	arr := strings.Fields(str)

	for _, s := range arr {
		i, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		sliceint = append(sliceint, i)
	}
	return sliceint, nil
}

//
func parse(file string) ([][]int, error) {
	if file == "" {
		return nil, errors.New("Wrong parameter: nill or empty")
	}
	square := [][]int{}
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	dstr := string(b)
	darr := strings.Split(dstr, "\n")

	for _, s := range darr {
		var d, err = readline(s)
		if err != nil {
			return nil, err
		}
		square = append(square, d)
	}

	return square, nil
}

func parseGetText(file string) (string, error) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return "null", err
	}
	d := string(b)
	d = strings.Trim(d, "\t")
	return d, nil
}
