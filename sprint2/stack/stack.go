package main

import (
	"errors"
)

type Stack struct {
	items []int
	max   *int
}

func (s *Stack) push(x int) {
	if s.max == nil || x > *s.max {
		s.max = &x
	}
	s.items = append(s.items, x)
}

func (s *Stack) pop() (int, error) {
	if len(s.items) == 0 {
		return -1, errors.New("empty pussy stack")
	}
	popped := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	if len(s.items) == 0 {
		s.max = nil
	} else if popped == *s.max {
		s.recalcMax()
	}
	return popped, nil
}

func (s *Stack) recalcMax() {
	max := s.items[0]
	for _, v := range s.items[1:] {
		if v > max {
			max = v
		}
	}
	s.max = &max
}

func (s *Stack) getMax() (int, error) {
	if s.max == nil {
		return -1, errors.New("cant get max")
	}
	return *s.max, nil
}
