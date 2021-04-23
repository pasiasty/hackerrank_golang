package main

import (
	"bufio"
	"os"

	"github.com/pasiasty/hackerrank_golang/solution"
	"github.com/pasiasty/hackerrank_golang/utils"
)

func main() {
	solution.Solution(utils.SkipBOM(bufio.NewReader(os.Stdin)), os.Stdout)
}
