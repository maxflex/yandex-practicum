package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func search(arr []int, x int) int {
	i := 0
	j := len(arr) - 1
	mid := (i + j) / 2
	for {
		if i > j {
			return -1
		}
		if arr[mid] >= x && (mid == 0 || arr[mid-1] < x) {
			return mid + 1
		}
		if arr[mid] >= x {
			j = mid - 1
		} else {
			i = mid + 1
		}
		mid = (i + j) / 2
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, 0, n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		x, _ := strconv.Atoi(scanner.Text())
		if len(arr) == n {
			n = x
		} else {
			arr = append(arr, x)
		}
	}
	fmt.Printf("%d %d", search(arr, n), search(arr, n*2))
}
