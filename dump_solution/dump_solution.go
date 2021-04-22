package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

var (
	fileRegexp = regexp.MustCompile(`import \((?P<imports>[\s\S]+?)\)(?P<rest>[\s\S]+)`)
)

func dumpFileWithoutPackage(filename string, imports map[string]interface{}, cw io.Writer) {
	ub, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Failed to read %s: %v", filename, err))
	}

	matches := fileRegexp.FindStringSubmatch(string(ub))

	importsIdx := fileRegexp.SubexpIndex("imports")
	restIdx := fileRegexp.SubexpIndex("rest")

	for _, l := range strings.Split(matches[importsIdx], "\n") {
		if trimmed := strings.Trim(l, "\t "); trimmed != "" {
			if !strings.Contains(trimmed, "github.com/pasiasty/hackerrank_golang") {
				imports[trimmed] = new(interface{})
			}
		}
	}
	cw.Write([]byte(fmt.Sprintf("\n// -- %v -- ", filename)))
	cw.Write([]byte(strings.ReplaceAll(matches[restIdx], "utils.", "")))
}

func main() {
	cw := bytes.NewBuffer(nil)
	imports := make(map[string]interface{})

	imports[`"os"`] = new(interface{})

	filepath.Walk("utils", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			panic(err.Error())
		}
		if strings.Contains(info.Name(), "_test.go") || !strings.Contains(info.Name(), ".go") {
			return nil
		}
		dumpFileWithoutPackage(path, imports, cw)
		return nil
	})

	dumpFileWithoutPackage(path.Join("solution", "solution.go"), imports, cw)

	if err := os.Mkdir("output", os.ModeDir|os.ModePerm); err != nil && !os.IsExist(err) {
		panic(fmt.Sprintf("Failed to create output directory: %v", err))
	}

	f, err := os.Create(path.Join("output", "output.go"))
	if err != nil {
		panic(fmt.Sprintf("Failed to open file: %v", err))
	}
	output := f

	output.Write([]byte("package main\n\n"))
	output.Write([]byte("import (\n"))

	importsSorted := []string{}
	for k := range imports {
		importsSorted = append(importsSorted, fmt.Sprintf("\t%v\n", k))
	}

	sort.Strings(importsSorted)

	for _, i := range importsSorted {
		output.Write([]byte(i))
	}

	output.Write([]byte(")\n"))
	output.Write(cw.Bytes())
	output.Write([]byte(
		`
func main() {
	Solution(SkipBOM(os.Stdin), os.Stdout)
}
`,
	))
}
