package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Stack(t *testing.T) {
	s := New()
	s.Push(0)

	assert.Equal(t, 0, s.curr, "curr == 0")

	s.Push(1)
	s.Push(2)

	value, _ := s.Pop()

	assert.Equal(t, 2, value, "pop == 2")
	assert.Equal(t, 1, s.curr, "curr = 1")
}
