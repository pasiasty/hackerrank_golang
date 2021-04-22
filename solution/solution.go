package solution

import (
	"fmt"
	"io"

	"github.com/pasiasty/hackerrank_golang/utils"
)

// Solution contains solution to the problem.
func Solution(r io.Reader, w io.Writer) {
	s := utils.MustReadLineOfInts(r, -1)

	sum := 0

	for _, el := range s {
		sum += el
	}

	w.Write([]byte(fmt.Sprintf("%d", sum)))
}
