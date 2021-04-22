package utils

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestSkipBOM(t *testing.T) {
	for _, tc := range []struct {
		name      string
		input     string
		expOutput string
	}{{
		name:      "no_bom",
		input:     "foo_bar_baz",
		expOutput: "foo_bar_baz",
	}, {
		name:      "with_bom",
		input:     "\uFEFFfoo_bar_baz",
		expOutput: "foo_bar_baz",
	}} {
		t.Run(tc.name, func(t *testing.T) {
			r := SkipBOM(bytes.NewBufferString(tc.input))
			b, err := ioutil.ReadAll(r)
			if err != nil {
				t.Fatalf("Failed to read from reader: %v", err)
			}
			if res := string(b); res != tc.expOutput {
				t.Errorf("Output different than expected, want: %v got: %v", tc.expOutput, res)
			}
		})
	}
}

func TestMustReadInt(t *testing.T) {
	for _, tc := range []struct {
		name        string
		input       string
		expOutput   int
		shouldPanic bool
	}{{
		name:      "val_1",
		input:     "43 ",
		expOutput: 43,
	}, {
		name:      "val_2",
		input:     "51\n",
		expOutput: 51,
	}, {
		name:        "no_int",
		input:       "a",
		shouldPanic: true,
	}} {
		t.Run(tc.name, func(t *testing.T) {
			if tc.shouldPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic")
					}
				}()
			}

			r := bytes.NewBufferString(tc.input)
			if res := MustReadInt(r); res != tc.expOutput {
				t.Errorf("Wrong output, want: %v got: %v", tc.expOutput, res)
			}
		})
	}
}

func TestMustReadNewline(t *testing.T) {
	for _, tc := range []struct {
		name        string
		input       string
		shouldPanic bool
	}{{
		name:  "newline",
		input: "\n",
	}, {
		name:        "no_newline",
		input:       "a",
		shouldPanic: true,
	}} {
		t.Run(tc.name, func(t *testing.T) {
			if tc.shouldPanic {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic")
					}
				}()
			}

			r := bytes.NewBufferString(tc.input)
			MustReadNewLine(r)
		})
	}
}
