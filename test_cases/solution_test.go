package test_cases

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"

	"github.com/pasiasty/hackerrank_golang/solution"
)

var (
	currDir = getCurrDir()
)

type testCase struct {
	name      string
	input     string
	expOutput []string
}

func getCurrDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}

func prepareTestCase(name string) *testCase {
	inputPath := path.Join(currDir, fmt.Sprintf("%s_input.txt", name))
	outputPath := path.Join(currDir, fmt.Sprintf("%s_output.txt", name))

	ib, err := ioutil.ReadFile(inputPath)
	if err != nil {
		panic(err)
	}

	f, err := os.Open(outputPath)
	if err != nil {
		panic(err)
	}

	expOutput := []string{}

	rb := bufio.NewReader(f)
	for {
		lb, _, err := rb.ReadLine()
		expOutput = append(expOutput, string(lb))
		if err == io.EOF {
			break
		}
	}

	return &testCase{
		name:      name,
		input:     string(ib),
		expOutput: expOutput,
	}
}

func TestSolution(t *testing.T) {
	testNames := []string{}

	filepath.Walk(currDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			panic(err.Error())
		}
		wantSuffix := "_input.txt"
		if strings.Contains(info.Name(), wantSuffix) {
			testNames = append(testNames, info.Name()[:len(info.Name())-len(wantSuffix)])
		}
		return nil
	})

	tcs := []*testCase{}

	for _, n := range testNames {
		tcs = append(tcs, prepareTestCase(n))
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			r := bytes.NewBufferString(tc.input)
			w := bytes.NewBuffer(nil)

			solution.Solution(bufio.NewReader(r), w)

			resReader := bufio.NewReader(bytes.NewBuffer(w.Bytes()))
			res := []string{}

			for {
				lb, _, err := resReader.ReadLine()
				res = append(res, string(lb))
				if err == io.EOF {
					break
				}
			}

			if !reflect.DeepEqual(res, tc.expOutput) {
				t.Errorf(
					"Wrong result, want:\n%v got:\n%v",
					strings.Join(tc.expOutput, "\n"),
					strings.Join(res, "\n"))
			}
		})
	}
}
