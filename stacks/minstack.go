package stacks

import (
	"errors"
)

type MinStack struct {
	count int
	stack *MinStackNode
}

type MinStackNode struct {
	value int
	min   int
	next  *MinStackNode
}

// New creates a new Stack.
func NewMinStack() (*MinStack, error) {
	s := &MinStack{
		count: 0,
	}

	return s, nil
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// Push pushes the provided value onto the given stack.
func (s *MinStack) Push(value int) error {
	var minVal int
	if s.count == 0 {
		minVal = value
	} else {
		minVal = min(s.stack.value, value)
	}

	node := &MinStackNode{
		value: value,
		min:   minVal,
		next:  s.stack,
	}

	s.stack = node
	s.count += 1

	return nil
}

func (s *MinStack) Min() (int, error) {
	if s.count == 0 {
		return -1, errors.New("attempted to get min of empty stack")
	}

	return s.stack.min, nil
}

// Pop pops the first value off the given stack.
func (s *MinStack) Pop() (int, error) {
	if s.count == 0 {
		return -1, errors.New("attempted to pop empty stack")
	}

	top := s.stack

	s.stack = top.next

	return top.value, nil
}

// Peek returns the value of the top of the stack.
func (s *MinStack) Peek() (int, error) {
	if s.count == 0 {
		return -1, errors.New("attempted to peek empty stack")
	}

	return s.stack.value, nil
}
