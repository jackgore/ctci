package stacks

import (
	"testing"
)

func TestMinStack(t *testing.T) {
	s, _ := NewMinStack()

	s.Push(5)

	if m, _ := s.Min(); m != 5 {
		t.Errorf("expected min to be 5 found %v", m)
	}

	s.Push(2)
	if m, _ := s.Min(); m != 2 {
		t.Errorf("expected min to be 2 found %v", m)
	}
	s.Push(7)
	if m, _ := s.Min(); m != 2 {
		t.Errorf("expected min to be 2 found %v", m)
	}
	s.Push(1)
	if m, _ := s.Min(); m != 1 {
		t.Errorf("expected min to be 1 found %v", m)
	}

	s.Pop()
	if m, _ := s.Min(); m != 2 {
		t.Errorf("expected min to be 2 found %v", m)
	}
}
