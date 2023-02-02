package fib

import "testing"

func TestFib(t *testing.T) {
	// Her går testen
	input = 1 // legg merke til operatøren =
	got = Fib(input)
	want = 1
	if got != want {
		t.Errorf("Fib(%d): want %d, got %d", input, want, got)
	}
}
