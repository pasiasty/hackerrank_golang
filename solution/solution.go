package solution

import (
	"bufio"
	"fmt"
	"io"

	"github.com/pasiasty/hackerrank_golang/utils"
)

type gridElement struct {
	x, y   int
	height int
}

func inGrid(x, y, rows, cols int) bool {
	if x < 0 || y < 0 || x >= cols || y >= rows {
		return false
	}
	return true
}

type bucketer struct {
	buckets   map[int]map[*gridElement]interface{}
	maxHeight int
}

func newBucketer() *bucketer {
	return &bucketer{buckets: make(map[int]map[*gridElement]interface{})}
}

func peekFromMap(m map[*gridElement]interface{}) *gridElement {
	for e := range m {
		return e
	}
	panic("I shouldn't be here!")
}

func (b *bucketer) putElement(el *gridElement) {
	if el.height > b.maxHeight {
		b.maxHeight = el.height
	}
	if bucket, ok := b.buckets[el.height]; ok {
		bucket[el] = new(interface{})
	} else {
		newBucket := make(map[*gridElement]interface{})
		newBucket[el] = new(interface{})
		b.buckets[el.height] = newBucket
	}
}

func (b *bucketer) pop() *gridElement {
	el := peekFromMap(b.buckets[b.maxHeight])

	if len(b.buckets[b.maxHeight]) == 1 {
		b.maxHeight--
	} else {
		delete(b.buckets[b.maxHeight], el)
	}

	return el
}

func (b *bucketer) updateHeight(el *gridElement, newHeight int) {
	delete(b.buckets[el.height], el)
	el.height = newHeight
	b.putElement(el)
}

func fixBoxIfNeeded(x, y, rows, cols, desiredHeight int, ge [][]*gridElement, b *bucketer) int {
	res := 0

	if inGrid(x, y, rows, cols) {
		el := ge[y][x]
		if el.height < desiredHeight {
			res += desiredHeight - el.height
			b.updateHeight(el, desiredHeight)
		}
	}

	return res
}

func testCase(rows, cols int, grid [][]int) int {
	res := 0

	b := newBucketer()

	ge := [][]*gridElement{}

	for y, row := range grid {
		ge = append(ge, []*gridElement{})
		for x, height := range row {
			el := &gridElement{x: x, y: y, height: height}
			b.putElement(el)
			ge[y] = append(ge[y], el)
		}
	}

	for i := 0; i < (rows * cols); i++ {
		el := b.pop()
		dh := el.height - 1
		res += fixBoxIfNeeded(el.x-1, el.y, rows, cols, dh, ge, b)
		res += fixBoxIfNeeded(el.x+1, el.y, rows, cols, dh, ge, b)
		res += fixBoxIfNeeded(el.x, el.y-1, rows, cols, dh, ge, b)
		res += fixBoxIfNeeded(el.x, el.y+1, rows, cols, dh, ge, b)
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
