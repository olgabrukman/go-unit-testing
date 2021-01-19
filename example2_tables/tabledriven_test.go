package example2_tables

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	type test struct {
		input string
		sep   string
		want  []string
	}

	tests := []test{
		{input: "a/b/c", sep: "/", want: []string{"a", "b", "c"}},
		{input: "a/b/c", sep: ",", want: []string{"a/b/c"}},
		{input: "abc/", sep: "/", want: []string{"abc", "/"}},
	}

	for _, testCase := range tests {
		got := Split(testCase.input, testCase.sep)
		if !reflect.DeepEqual(testCase.want, got) {
			t.Errorf("expected: %v, got: %v", testCase.want, got)
		}
	}
}

func TestSplitWithName(t *testing.T) {
	tests := []struct {
		name  string
		input string
		sep   string
		want  []string
	}{
		{name: "simple", input: "a/b/c", sep: "/", want: []string{"a", "b", "c"}},
		{name: "wrong sep", input: "a/b/c", sep: ",", want: []string{"a/b/c"}},
		{name: "trailing sep", input: "a/b/c/ ", sep: "/", want: []string{"a", "b", "c", "/"}},
		{name: "no sep", input: "abc", sep: "/", want: []string{"abc"}},
	}

	for _, testCase := range tests {
		got := Split(testCase.input, testCase.sep)
		if !reflect.DeepEqual(testCase.want, got) {
			//Fatalf to exit right away on first failure
			t.Errorf("%s: expected: %v, got: %v", testCase.name, testCase.want, got)
		}
	}
}

func TestSplitWithRun(t *testing.T) {
	tests := []struct {
		name  string
		input string
		sep   string
		want  []string
	}{
		{name: "simple", input: "a/b/c", sep: "/", want: []string{"a", "b", "c"}},
		{name: "wrong sep", input: "a/b/c", sep: ",", want: []string{"a/b/c"}},
		{name: "trailing sep", input: "a/b/c/", sep: "/", want: []string{"a", "b", "c", ""}},
		{name: "no sep", input: "abc", sep: "/", want: []string{"abc"}},
	}

	for _, testCase := range tests {
		tc := testCase // NOTE: https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
		// Run runs a testcase test as a subtest of t called name. It runs the testcase test in a separate goroutine
		// and blocks until the test returns or calls t.Parallel to become a parallel test.
		// Run reports whether the testcase succeeded (or at least did not fail before calling t.Parallel).
		t.Run(tc.name, func(*testing.T) {
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}
		})

	}
}
func TestSplitParallel(t *testing.T) {
	type test struct {
		name  string
		input string
		sep   string
		want  []string
	}

	tests := []test{
		{name: "simple", input: "a/b/c", sep: "/", want: []string{"a", "b", "c"}},
		{name: "wrong sep", input: "a/b/c", sep: ",", want: []string{"a/b/c"}},
		{name: "trailing sep", input: "a/b/c/", sep: "/", want: []string{"a", "b", "c", ""}},
		{name: "no sep", input: "abc", sep: "/", want: []string{"abc"}},
	}

	t.Parallel()
	for _, testCase := range tests {
		tc := testCase // NOTE: https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
		t.Run(tc.name, func(*testing.T) {
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
			}
		})
	}
}
