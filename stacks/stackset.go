package stacks

import (
	"errors"
)

type StackList struct {
	value *Stack
	next  *StackList
}

type StackSet struct {
	maxCapacity int
	itemCount   int
	stackCount  int
	stacks      *StackList
}

// New creates a new stack set with the provided maxCapacity.
func NewStackSet(max int) *StackSet {
	stack, _ := NewStack()

	s := &StackSet{
		stacks: &StackList{
			value: stack,
		},
		stackCount:  1,
		maxCapacity: max,
	}

	return s
}

// Push pushes a new value onto our list of stacks.
func (s *StackSet) Push(value int) error {
	// First we need to determine if we need to allocate a new stack.
	// Our total capacity is maxCapacity * stackCount
	if (s.maxCapacity * s.stackCount) == s.itemCount {
		// Create the new Stack to use
		stack, _ := NewStack()
		stack.Push(value)

		// Append to front of our stack list
		list := &StackList{
			value: stack,
			next:  s.stacks,
		}

		s.stacks = list
		s.stackCount += 1
	} else {
		// Otherwise we just push to head of the stack list and increase our count.
		s.stacks.value.Push(value)
	}

	s.itemCount += 1

	return nil
}

// Pop removes the item from the stack
func (s *StackSet) Pop() (int, error) {
	if s.itemCount == 0 {
		return -1, errors.New("attempted to pop from empty stackset")
	}

	// We just pop from the first stack and remove it if needed.
	value, _ := s.stacks.value.Pop()
	s.itemCount -= 1

	if s.itemCount == ((s.stackCount - 1) * s.maxCapacity) {
		// Removes the top stack
		s.stacks = s.stacks.next
		s.stackCount -= 1
	}

	return value, nil
}
