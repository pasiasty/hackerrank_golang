package utils

import (
	"bufio"
	"fmt"
	"io"
	"log"
)

// SkipBOM removes BOM from the reader.
func SkipBOM(r io.Reader) io.Reader {
	br := bufio.NewReader(r)
	rr, _, err := br.ReadRune()
	if err != nil {
		log.Fatal(err)
	}
	if rr != '\uFEFF' {
		br.UnreadRune() // Not a BOM -- put the rune back
	}

	return br
}

// MustReadInt reads int from reader
func MustReadInt(r io.Reader) int {
	var res int
	_, err := fmt.Fscanf(r, "%d", &res)
	if err != nil {
		log.Fatalf("Failed to read int: %v", err)
	}
	return res
}

// MustReadNewLine reads newline from reader
func MustReadNewLine(r io.Reader) {
	_, err := fmt.Fscanln(r)
	if err != nil {
		log.Fatalf("Failed to read newline: %v", err)
	}
}
