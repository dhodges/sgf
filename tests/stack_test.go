package tests

import (
	"testing"

	"github.com/dhodges/sgfinfo/parse"
)

func TestStack(t *testing.T) {
	stack := new(parse.Stack)

	stack.Push("H")
	stack.Push("R")
	stack.Push("Puffenstuff")

	if stack.Len() != 3 {
		t.Error("stack count is wrong")
	}

	item := stack.Peek()
	if item != "Puffenstuff" {
		t.Errorf("stack is wrong, expected 'Puffenstuff' but found %q", item)
	}

	item = stack.Pop()
	if item != "Puffenstuff" {
		t.Errorf("stack is wrong, expected 'Puffenstuff' but found %q", item)
	}

	item = stack.Pop()
	if item != "R" {
		t.Errorf("stack is wrong, expected 'R' but found %q", item)
	}

	item = stack.Pop()
	if item != "H" {
		t.Errorf("stack is wrong, expected 'H' but found %q", item)
	}
}
