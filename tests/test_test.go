package main

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// USING ONLY STRUCTS

// var tests = []struct {
// 	name  string
// 	input string
// 	sep   string
// 	want  []string
// }{
// 	{name: "Test against right Answer", input: "a/b/c", sep: "/", want: []string{"a", "b", "c"}},
// 	{name: "Test against wrong sep", input: "a/b/c", sep: ",", want: []string{"a/b/c"}},
// 	{name: "Test against trailling seps", input: "a/b/c/", sep: "/", want: []string{"a", "b", "c"}},
// }

// USING STRUCTS AND A MAP

var tests = map[string]struct {
	input string
	sep   string
	want  []string
}{
	"Test against right Answer": {input: "a/b/c", sep: "/", want: []string{"a", "b", "c"}},
	"Test against wrong sep":    {input: "a/b/c", sep: ",", want: []string{"a/b/c"}},
	"trailling":                 {input: "a/b/c/", sep: "/", want: []string{"a", "b", "c"}},
}

func TestFinal(t *testing.T) {
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := Split(test.input, test.sep)
			diff := cmp.Diff(test.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func TestSplit(t *testing.T) {

	got := Split("a/b/c", "/")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Expected %v but got %v", want, got)
	}
}

func TestSplitWrong(t *testing.T) {
	got := Split("a/b/c", ",")
	want := []string{"a/b/c"}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Expected %v but got %v", want, got)
	}
}
