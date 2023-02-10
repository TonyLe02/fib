package fib

import "testing"

func TestFib(t *testing.T) {
	// Her går testen
	type test struct {
		input int
		want  int
	}

	var tests = []test{
		{3, 1},
		{4, 1},
		{5, 1},
		{6, 1},
		{7, 1},
		{8, 1},
		{9, 1},
		{10, 1},
		{11, 1},
		{12, 1},
		{13, 1},
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
