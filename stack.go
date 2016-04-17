package stack

import (
	"errors"
)

const (
	DEPTH = 200
)

var (
	ErrOverDepth = errors.New("over the depth 200")
	ErrEmpth     = errors.New("the stack is empty")
)

type Stack struct {
	curr int
	src  []int
}

func New() *Stack {
	return &Stack{
		curr: -1,
		src:  make([]int, DEPTH),
	}
}

func (s *Stack) Init() {
	if s.src == nil {
		s.src = make([]int, DEPTH)
	}
	s.curr = -1
}

func (s *Stack) Push(value int) error {
	if s.curr == 200-1 {
		return ErrOverDepth
	}
	s.curr++
	s.src[s.curr] = value

	return nil
}

func (s *Stack) Pop() (int, error) {
	if s.curr < 0 {
		return 0, ErrEmpth
	}

	value := s.src[s.curr]
	s.curr--

	return value, nil
}

func (s *Stack) Depth() int {
	return s.curr + 1
}
