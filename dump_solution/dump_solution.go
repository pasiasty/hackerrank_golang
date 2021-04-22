package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

var (
	fileRegexp = regexp.MustCompile(`import \((?P<imports>[\s\S]+?)\)(?P<rest>[\s\S]+)`)
)

func dumpFileWithoutPackage(filename string, imports map[string]interface{}, cw io.Writer) {
	ub, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read %s: %v", filename, err)
	}

	matches := fileRegexp.FindAllStringSubmatch(string(ub), -1)

	for _, l := range strings.Split(matches[0][1], "\n") {
		if trimmed := strings.Trim(l, "\t "); trimmed != "" {
			imports[trimmed] = new(interface{})
		}
	}
	cw.Write([]byte(matches[0][2]))
}

func main() {
	cw := bytes.NewBuffer(nil)
	imports := make(map[string]interface{})

	dumpFileWithoutPackage("utils/utils.go", imports, cw)
	dumpFileWithoutPackage("solution/solution.go", imports, cw)

	f, err := os.Create("output.go")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	output := f

	output.Write([]byte("package main\n\n"))
	output.Write([]byte("import (\n"))
	for k, _ := range imports {
		output.Write([]byte(fmt.Sprintf("\t%v\n", k)))
	}
	output.Write([]byte(")"))
	output.Write(cw.Bytes())
	output.Write([]byte(
		`
func main() {
	Solution(SkipBOM(os.Stdin), os.Stdout)
}
`,
	))
}
