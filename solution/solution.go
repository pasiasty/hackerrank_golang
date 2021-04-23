package solution

import (
	"bufio"
	"fmt"
	"io"

	"github.com/pasiasty/hackerrank_golang/utils"
)

func testCase(rows, cols int, grid [][]int) int {
	return 0
}

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

		w.Write([]byte(fmt.Sprintf("Case #%d: %d\n", i+1, testCase(rows, cols, grid))))
	}
}
