package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, result int
	fmt.Scan(&n)
	arr := make([]int16, 0, 3)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		switch len(arr) {
		case 2:
			if arr[0] > arr[1] {
				result++
			}
		case 3:
			arr = arr[1:]
		}
		x, _ := strconv.Atoi(scanner.Text())
		arr = append(arr, int16(x))
		if len(arr) == 3 {
			if arr[1] > arr[0] && arr[1] > arr[2] {
				result++
			}
		}
	}
	switch len(arr) {
	case 1:
		result++
	case 2:
		if arr[1] > arr[0] {
			result++
		}
	case 3:
		if arr[2] > arr[1] {
			result++
		}
	}
	fmt.Print(result)
}
