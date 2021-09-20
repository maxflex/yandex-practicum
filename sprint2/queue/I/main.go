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
	items []int
	i, j  int
}

func (q *Queue) init(size int) {
	q.items = make([]int, size)
	fmt.Println("Creted queue of size", size, cap(q.items))
	q.i, q.j = 0, 0
}

func (q *Queue) push(x int) bool {
	fmt.Println("Will push at index", q.j, "Q.i", q.i, q.isFull())
	if q.isFull() {
		return false
	}
	q.items[q.j] = x
	q.j = (q.j + 1) % cap(q.items)
	return true
}

func (q *Queue) isEmpty() bool {
	return q.j == q.i
}

func (q *Queue) isFull() bool {
	return q.i == (q.j+1)%cap(q.items)
}

func (q *Queue) pop() (int, error) {
	if q.isEmpty() {
		return -1, errors.New("queue empty")
	}
	popped := q.items[q.i]
	q.i = (q.i + 1) % cap(q.items)
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
			if !q.push(x) {
				fmt.Fprintln(writer, "error")
			}
		case "pop":
			if popped, err := q.pop(); err != nil {
				fmt.Fprintln(writer, "None")
			} else {
				fmt.Fprintln(writer, popped)
			}
		case "size":
			fmt.Fprintln(writer, cap(q.items))
		}
	}
	writer.Flush()
}
