package example_extra_benchmark

import "testing"

/*
Reference: https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go
Run regular test and benchmark tests: go test -bench=.
Run matching regular tests and benchmark tests: go test -run=XXX -bench=.
The above pattern does not match any regular tests. Thus, only benchmark tests will run
*/

type fibTest struct {
	n        int
	expected int
}

var fibTests = []fibTest{
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{7, 13},
}

func TestFib(t *testing.T) {
	for _, tt := range fibTests {
		actual := Fib(tt.n)
		if actual != tt.expected {
			t.Errorf("Fib(%d): expected %d, actual %d", tt.n, tt.expected, actual)
		}
	}
}

func BenchmarkFib10(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		Fib(10)
	}
}
