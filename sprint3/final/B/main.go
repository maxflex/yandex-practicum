package main

/**
ID успешной посылки: 54286437

Временная сложность
В среднем случае: O(n * log(n))
В худьшем случае: O(n^2)

Пространственная сложность: O(1)
*/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
Объявляем тип участника
*/
type Participant struct {
	login   string
	score   int
	penalty int
}

/**
Функция компаратор двух участников
*/
func less(a, b Participant) bool {
	if a.score != b.score {
		return a.score < b.score
	}
	if a.penalty != b.penalty {
		return a.penalty > b.penalty
	}
	return a.login > b.login
}

func partition(arr []Participant, i, j int) int {
	pivot := arr[(i+j)/2] // берём середину
	for j >= i {
		for less(pivot, arr[i]) {
			i++
		}
		for less(arr[j], pivot) {
			j--
		}
		if j >= i {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}
	return i
}

func quickSort(arr []Participant, i, j int) {
	if j-i <= 0 {
		return
	}
	mid := partition(arr, i, j)
	quickSort(arr, i, mid-1)
	quickSort(arr, mid, j)
}

func printResult(arr []Participant) {
	writer := bufio.NewWriter(os.Stdout)
	for _, v := range arr {
		writer.WriteString(v.login + "\n")
	}
	writer.Flush()
}

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]Participant, n)
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		scanner.Scan()
		fields := strings.Fields(scanner.Text())
		score, _ := strconv.Atoi(fields[1])
		penalty, _ := strconv.Atoi(fields[2])
		arr[i] = Participant{fields[0], score, penalty}
	}
	quickSort(arr, 0, len(arr)-1)
	printResult(arr)
}
