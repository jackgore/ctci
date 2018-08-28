package stacks

import (
	"errors"
)

type Stack struct {
	count int
	stack *StackNode
}

type StackNode struct {
	value int
	next  *StackNode
}

// New creates a new Stack.
func NewStack() (*Stack, error) {
	s := &Stack{
		count: 0,
	}

	return s, nil
}

// Push pushes the provided value onto the given stack.
func (s *Stack) Push(value int) error {
	node := &StackNode{
		value: value,
		next:  s.stack,
	}

	s.stack = node
	s.count += 1

	return nil
}

// Pop pops the first value off the given stack.
func (s *Stack) Pop() (int, error) {
	if s.count == 0 {
		return -1, errors.New("attempted to pop empty stack")
	}

	top := s.stack

	s.stack = top.next

	return top.value, nil
}

// Peek returns the value of the top of the stack.
func (s *Stack) Peek() (int, error) {
	if s.count == 0 {
		return -1, errors.New("attempted to peek empty stack")
	}

	return s.stack.value, nil
}
