package fib

import "testing"

func TestFib(t *testing.T) {
	// Her går testen
	type test struct {
		input int
		want  int
	}

	var tests = []test{
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{7, 13},
		{8, 21},
		{9, 34},
		{10, 55},
		{11, 89},
		{12, 144},
		{13, 233},
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
