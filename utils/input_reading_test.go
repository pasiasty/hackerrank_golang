package utils

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"reflect"
	"strings"
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
			r := SkipBOM(bufio.NewReader(bytes.NewBufferString(tc.input)))
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

func TestMustReadFloat(t *testing.T) {
	for _, tc := range []struct {
		name        string
		input       string
		expOutput   float64
		shouldPanic bool
	}{{
		name:      "val_1",
		input:     "43.123 ",
		expOutput: 43.123,
	}, {
		name:      "val_2",
		input:     "51.7265\n",
		expOutput: 51.7265,
	}, {
		name:        "no_float",
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
			if res := MustReadFloat(r); res != tc.expOutput {
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

func TestMustReadLineOfInts(t *testing.T) {
	for _, tc := range []struct {
		name             string
		input            string
		expOutput        []int
		wantNumOfResults int
		shouldPanic      bool
	}{{
		name:      "dont_check_num_of_results",
		input:     "1 3 15\n17 16",
		expOutput: []int{1, 3, 15},
	}, {
		name:      "single_line",
		input:     "1 3 15",
		expOutput: []int{1, 3, 15},
	}, {
		name:             "check_num_of_results",
		input:            "1 3 15\n17 16",
		expOutput:        []int{1, 3, 15},
		wantNumOfResults: 3,
	}, {
		name:             "wrong_num_of_results",
		input:            "1 3 15\n17 16",
		wantNumOfResults: 4,
		shouldPanic:      true,
	}, {
		name:        "wrong_format",
		input:       "1 3 15a\n17 16",
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

			r := strings.NewReader(tc.input)
			if res := MustReadLineOfInts(bufio.NewReader(r), tc.wantNumOfResults); !reflect.DeepEqual(tc.expOutput, res) {
				t.Errorf("Wrong result, want: %v got: %v", tc.expOutput, res)
			}
		})
	}
}

func TestMustReadLineOfFloats(t *testing.T) {
	for _, tc := range []struct {
		name             string
		input            string
		expOutput        []float64
		wantNumOfResults int
		shouldPanic      bool
	}{{
		name:      "dont_check_num_of_results",
		input:     "1.2 3.7 15.8\n17 16",
		expOutput: []float64{1.2, 3.7, 15.8},
	}, {
		name:      "single_line",
		input:     "1.2 3.7 15.8",
		expOutput: []float64{1.2, 3.7, 15.8},
	}, {
		name:             "check_num_of_results",
		input:            "1.2 3.7 15.8\n17 16",
		expOutput:        []float64{1.2, 3.7, 15.8},
		wantNumOfResults: 3,
	}, {
		name:             "wrong_num_of_results",
		input:            "1 3 15\n17 16",
		wantNumOfResults: 4,
		shouldPanic:      true,
	}, {
		name:        "wrong_format",
		input:       "1 3 15a\n17 16",
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

			r := strings.NewReader(tc.input)
			if res := MustReadLineOfFloats(bufio.NewReader(r), tc.wantNumOfResults); !reflect.DeepEqual(tc.expOutput, res) {
				t.Errorf("Wrong result, want: %v got: %v", tc.expOutput, res)
			}
		})
	}
}
