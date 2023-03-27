package stack_test

import (
	"testing"

	stack "github.com/mohammad-quanit/go-tdd/Stack"
)

func TestEmpty(t *testing.T) {
	s := stack.NewStack()
	if s.Empty() != true {
		t.Error("Stack was not empty!")
	}
}

func TestNotEmpty(t *testing.T) {
	s := stack.NewStack()
	s.Add("Quanit")
	if s.Empty() != false {
		t.Error("Stack was empty")
	}
}
