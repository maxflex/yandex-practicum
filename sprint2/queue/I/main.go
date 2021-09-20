package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Queue struct {
	items      []int
	i, j, size int
}

func (q *Queue) init(size int) {
	q.items = make([]int, size)
	q.i, q.j = -1, -1
	q.size = 0
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
	q.size++
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
		return -1, errors.New("queue empty")
	}
	popped := q.items[q.i]
	q.size--
	if q.i == q.j {
		q.i, q.j = -1, -1
	} else {
		q.i = (q.i + 1) % cap(q.items)
	}
	return popped, nil
}

func (q *Queue) peek() (int, error) {
	if q.isEmpty() {
		return -1, errors.New("queue empty")
	}
	return q.items[q.i], nil
}

func main() {
	var size int
	var q Queue
	fmt.Scan(&size)
	fmt.Scan(&size)
	q.init(size)
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		// fmt.Println(line)
		switch line[0] {
		case "peek":
			if v, err := q.peek(); err != nil {
				fmt.Fprintln(writer, "None")
			} else {
				fmt.Fprintln(writer, v)
			}
		case "push":
			x, _ := strconv.Atoi(line[1])
			if !q.enqueue(x) {
				fmt.Fprintln(writer, "error")
			}
		case "pop":
			if popped, err := q.dequeue(); err != nil {
				fmt.Fprintln(writer, "None")
			} else {
				fmt.Fprintln(writer, popped)
			}
		case "size":
			fmt.Fprintln(writer, q.size)
		}
	}
	writer.Flush()
}
