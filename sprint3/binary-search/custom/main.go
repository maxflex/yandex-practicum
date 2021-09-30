package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
Основная идея: если искомое меньше, то делаем:
arr = arr[:mid]
если искомое больше, делаем:
arr = arr[mid+1:]

Но проблема в том, когда так модифицируем arr, то
теряются иднексы. Надо как-то сохранить индекс
Запутанная логика в функции в основном связана с тем,
как сохранить индекс, когда исходный массив так меняется
*/
func search(arr []int, x int) int {
	mid := (len(arr) - 1) / 2
	index := mid
	for arr[mid] != x {
		less := false
		if x > arr[mid] {
			arr = arr[mid+1:]
			index++
		} else {
			arr = arr[:mid]
			index--
			less = true
		}
		if len(arr) == 0 {
			return -1
		}
		mid = (len(arr) - 1) / 2
		if less {
			index -= (len(arr) - 1 - mid)
		} else {
			index += mid
		}
	}
	return index
}

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, 0, n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	words := strings.Fields(scanner.Text())
	for _, v := range words {
		x, _ := strconv.Atoi(v)
		arr = append(arr, x)
	}
	scanner.Scan()
	price, _ := strconv.Atoi(scanner.Text())
	fmt.Println(arr, price)
	fmt.Println(search(arr, price))
}
