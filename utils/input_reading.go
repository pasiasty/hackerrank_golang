package utils

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// SkipBOM removes BOM from the reader.
func SkipBOM(r *bufio.Reader) io.Reader {
	rr, _, err := r.ReadRune()
	if err != nil {
		panic(err)
	}
	if rr != '\uFEFF' {
		r.UnreadRune() // Not a BOM -- put the rune back
	}

	return r
}

// MustReadInt reads int from reader
func MustReadInt(r io.Reader) int {
	var res int
	_, err := fmt.Fscanf(r, "%d", &res)
	if err != nil {
		panic(fmt.Sprintf("Failed to read int: %v", err))
	}
	return res
}

// MustReadFloat reads int from reader
func MustReadFloat(r io.Reader) float64 {
	var res float64
	_, err := fmt.Fscanf(r, "%f", &res)
	if err != nil {
		panic(fmt.Sprintf("Failed to read float: %v", err))
	}
	return res
}

// MustReadNewLine reads newline from reader
func MustReadNewLine(r io.Reader) {
	_, err := fmt.Fscanln(r)
	if err != nil {
		panic(fmt.Sprintf("Failed to read newline: %v", err))
	}
}

// MustReadLineOfInts reads line of ints and returns them as a slice.
func MustReadLineOfInts(r *bufio.Reader, wantNumOfResults int) []int {
	b, _, err := r.ReadLine()
	l := string(b)

	if err != nil && err != io.EOF {
		panic(fmt.Sprintf("Could not read string from reader: %v", err))
	}

	split := strings.Split(l, " ")
	res := []int{}

	for _, el := range split {
		i, err := strconv.Atoi(strings.Trim(el, " \n\t\r"))
		if err != nil {
			panic(fmt.Sprintf("Failed to convert '%s' to int: %v", el, err))
		}
		res = append(res, i)
	}

	if wantNumOfResults > 0 && len(res) != wantNumOfResults {
		panic(fmt.Sprintf("Wrong amount of results, want: %v got: %v", wantNumOfResults, len(res)))
	}
	return res
}

// MustReadLineOfFloats reads line of ints and returns them as a slice.
func MustReadLineOfFloats(r *bufio.Reader, wantNumOfResults int) []float64 {
	l, err := r.ReadString('\n')
	if err != nil && err != io.EOF {
		panic(fmt.Sprintf("Could not read string from reader: %v", err))
	}

	split := strings.Split(l, " ")
	res := []float64{}

	for _, el := range split {
		i, err := strconv.ParseFloat(strings.Trim(el, " \n\t\r"), 64)
		if err != nil {
			panic(fmt.Sprintf("Failed to convert '%s' to int: %v", el, err))
		}
		res = append(res, i)
	}

	if wantNumOfResults > 0 && len(res) != wantNumOfResults {
		panic(fmt.Sprintf("Wrong amount of results, want: %v got: %v", wantNumOfResults, len(res)))
	}
	return res
}
