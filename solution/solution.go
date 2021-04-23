package solution

import (
	"bufio"
	"io"
	"log"

	"github.com/pasiasty/hackerrank_golang/utils"
)

// Solution contains solution to the problem.
func Solution(r *bufio.Reader, w io.Writer) {
	n := utils.MustReadLineOfInts(r, 1)[0]

	for i := 0; i < n; i++ {
		dims := utils.MustReadLineOfInts(r, 2)
		rows, cols := dims[0], dims[1]

		grid := [][]int{}

		for y := 0; y < rows; y++ {
			grid = append(grid, utils.MustReadLineOfInts(r, cols))
		}

		log.Printf("grid: %v", grid)
	}
}
