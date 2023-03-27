package main

import (
	"reflect"
	"testing"
)

func TestString(t *testing.T) {
	got := HelloWorld()
	want := "Hello World!."

	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}

func TestSum(t *testing.T) {
	t.Run("Sum of numbers in array", func(t *testing.T) {
		numbers := []int{5, 4, 3, 2, 1}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d, want %d, given %d", got, want, numbers)
		}
	})

	t.Run("Sum of numbers in slice", func(t *testing.T) {
		numbers := []int{5, 4, 3, 2, 1}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d, want %d, given %d", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := "[]int{3, 9}"

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
