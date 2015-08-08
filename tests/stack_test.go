package tests

import (
	"testing"

	"github.com/dhodges/sgfinfo/parse"
  "github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	stack := new(parse.Stack)

	stack.Push("H")
	stack.Push("R")
	stack.Push("Puffenstuff")

  assert.Equal(t, stack.Len(),  3, "stack count is wrong")

  assert.Equal(t, stack.Peek(), "Puffenstuff", "stack is wrong")
  assert.Equal(t, stack.Pop(),  "Puffenstuff", "stack is wrong")
  assert.Equal(t, stack.Pop(),  "R",           "stack is wrong")
  assert.Equal(t, stack.Pop(),  "H",           "stack is wrong")
}
