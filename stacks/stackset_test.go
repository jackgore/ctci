package stacks

import (
	"testing"
)

func TestStackSet(t *testing.T) {
	s := NewStackSet(3)

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)

	val5, _ := s.Pop()
	val4, _ := s.Pop()
	val3, _ := s.Pop()
	val2, _ := s.Pop()
	val1, _ := s.Pop()

	if val5 != 5 {
		t.Errorf("expected 5 got: %v", val5)
	}

	if val4 != 4 {
		t.Errorf("expected 4 got: %v", val4)
	}

	if val3 != 3 {
		t.Errorf("expected 3 got: %v", val3)
	}

	if val2 != 2 {
		t.Errorf("expected 2 got: %v", val2)
	}

	if val1 != 1 {
		t.Errorf("expected 1 got: %v", val1)
	}

	_, err := s.Pop()
	if err == nil {
		t.Errorf("Expected error received nil")
	}
}
