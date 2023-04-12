package main

import (
	"testing"
)

func TestString(t *testing.T) {
	got := HelloWorld()
	want := "Hello World!"

	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}

func TestSum(t *testing.T) {
	t.Run("Sum of numbers in array", func(t *testing.T) {
		numbers := []int{5, 4, 3, 2, 1}

		got := Sum(numbers)
		want := 16

		if got != want {
			t.Errorf("got %d, want %d, given %d", got, want, numbers)
		}
	})
}
