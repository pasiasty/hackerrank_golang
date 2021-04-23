package solution

import (
	"bufio"
	"fmt"
	"io"

	"github.com/pasiasty/hackerrank_golang/utils"
)

type heapElement struct {
	height int
	x, y   int
}

func (e *heapElement) Key() int {
	return e.height
}

func (e *heapElement) desiredHeight(h int) int {
	if e.height < h {
		res := h - e.height
		e.height = h
		return res
	}
	return 0
}

func inGrid(x, y, rows, cols int) bool {
	if x < 0 || y < 0 || x >= cols || y >= rows {
		return false
	}
	return true
}

func testCase(rows, cols int, grid [][]int) int {
	heap := utils.NewMaxHeap()
	elements := [][]*heapElement{}

	for y, row := range grid {
		elements = append(elements, []*heapElement{})

		for x, el := range row {
			elements[y] = append(elements[y], &heapElement{height: el, x: x, y: y})
			heap.Push(elements[y][x])
		}
	}

	res := 0

	for heap.Len() > 0 {
		// for _, row := range elements {
		// 	for _, el := range row {
		// 		fmt.Printf("%+v ", el)
		// 	}
		// 	fmt.Printf("\n")
		// }
		el := heap.Peek()
		heap.Pop()

		hel := el.(*heapElement)

		dh, x, y := hel.height-1, hel.x, hel.y

		if inGrid(x-1, y, rows, cols) {
			res += elements[y][x-1].desiredHeight(dh)
			heap.UpdatePosition(elements[y][x-1])
		}
		if inGrid(x+1, y, rows, cols) {
			res += elements[y][x+1].desiredHeight(dh)
			heap.UpdatePosition(elements[y][x+1])
		}
		if inGrid(x, y-1, rows, cols) {
			res += elements[y-1][x].desiredHeight(dh)
			heap.UpdatePosition(elements[y-1][x])
		}
		if inGrid(x, y+1, rows, cols) {
			res += elements[y+1][x].desiredHeight(dh)
			heap.UpdatePosition(elements[y+1][x])
		}
	}

	return res
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
