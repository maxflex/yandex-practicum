package main

import (
	"errors"
)

type Stack struct {
	items []rune
}

func (s *Stack) push(x rune) {
	s.items = append(s.items, x)
}

func (s *Stack) pop() (rune, error) {
	if len(s.items) == 0 {
		return -1, errors.New("empty pussy stack")
	}
	popped := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return popped, nil
}

func (s *Stack) empty() bool {
	return len(s.items) == 0
}
