package utils

import (
	"bufio"
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
