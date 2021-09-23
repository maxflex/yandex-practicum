package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	items []int
	i, j  int
}

func (q *Queue) init(size int) {
	q.items = make([]int, size)
	q.i, q.j = -1, -1
}

func (q *Queue) enqueue(x int) bool {
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

func (q *Queue) dequeue() (int, error) {
	if q.isEmpty() {
		return 0, errors.New("queue empty")
	}
	popped := q.items[q.i]
	if q.i == q.j {
		q.i, q.j = -1, -1
	} else {
		q.i = (q.i + 1) % cap(q.items)
	}
	return popped, nil
}

func (q *Queue) peek() (int, error) {
	if q.isEmpty() {
		return 0, errors.New("queue empty")
	}
	return q.items[q.i], nil
}

func solve(n int, k uint8) int {
	if n < 2 {
		return 1
	}
	var (
		q      Queue
		i      int
		result int
	)
	mod := pow(10, k)
	q.init(2)
	q.enqueue(1)
	q.enqueue(1)
	for i = 2; i <= n; i++ {
		popped, _ := q.dequeue()
		peek, _ := q.peek()
		result = (popped + peek) % mod
		if i != n {
			q.enqueue(result)
		}
	}
	return result
}

func pow(base int, exponent uint8) int {
	if exponent == 0 {
		return 1
	}
	var i uint8
	result := base
	for i = 1; i < exponent; i++ {
		result *= base
	}
	return result
}

func main() {
	var (
		n int
		k uint8
	)
	fmt.Scanf("%d %d", &n, &k)
	fmt.Print(solve(n, k))
}
