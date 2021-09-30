package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func search(arr []int, x, i, j int) int {
	if i > j {
		return -1
	}
	mid := (i + j) / 2
	if arr[mid] == x {
		return mid
	}
	if x > arr[mid] {
		return search(arr, x, mid+1, j)
	} else {
		return search(arr, x, i, mid-1)
	}
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
	fmt.Println(search(arr, price, 0, len(arr)-1))
}
