package main

import (
	"errors"
)

type Stack struct {
	items []int
	max   []int
}

func (s *Stack) push(x int) {
	var max int
	if len(s.max) == 0 || x > s.max[len(s.max)-1] {
		max = x
	} else {
		max = s.max[len(s.max)-1]
	}
	s.max = append(s.max, max)
	s.items = append(s.items, x)
}

func (s *Stack) pop() (int, error) {
	if len(s.items) == 0 {
		return -1, errors.New("empty pussy stack")
	}
	popped := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	s.max = s.max[:len(s.max)-1]
	return popped, nil
}

func (s *Stack) getMax() (int, error) {
	if len(s.max) == 0 {
		return -1, errors.New("empty pussy stack")
	}
	return s.max[len(s.max)-1], nil
}
