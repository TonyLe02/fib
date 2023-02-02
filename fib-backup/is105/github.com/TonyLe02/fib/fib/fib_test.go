package fib

import "testing"

func TestFib(t *testing.T) {
	// Her går testen
	type test struct {
		input int
		want  int
	}

	var tests = []test{
		{10, 1},
		// her kan vi bare legge på nye test cases
	}

	// tc - test case
	for _, tc := range tests {
		got := Fib(tc.input)
		if got != tc.want {
			t.Errorf("Fib(%d): want %d, got %d", tc.input, tc.want, got)
		}
	}
}
