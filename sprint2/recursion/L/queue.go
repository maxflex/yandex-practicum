package main

import (
	"errors"
	"math/big"
)

type Queue struct {
	items []*big.Int
	i, j  int
}

func (q *Queue) init(size int) {
	q.items = make([]*big.Int, size)
	q.i, q.j = -1, -1
}

func (q *Queue) enqueue(x *big.Int) bool {
	if q.isFull() {
		return false
	}
	if q.j == -1 {
		q.j, q.i = 0, 0
	} else {
		q.j = (q.j + 1) % cap(q.items)
	}
	q.items[q.j] = x
	return true
}

func (q *Queue) isEmpty() bool {
	return q.i == -1
}

func (q *Queue) isFull() bool {
	return (q.j+1)%cap(q.items) == q.i
}

func (q *Queue) dequeue() (*big.Int, error) {
	if q.isEmpty() {
		return big.NewInt(0), errors.New("queue empty")
	}
	popped := q.items[q.i]
	if q.i == q.j {
		q.i, q.j = -1, -1
	} else {
		q.i = (q.i + 1) % cap(q.items)
	}
	return popped, nil
}

func (q *Queue) peek() (*big.Int, error) {
	if q.isEmpty() {
		return big.NewInt(0), errors.New("queue empty")
	}
	return q.items[q.i], nil
}
