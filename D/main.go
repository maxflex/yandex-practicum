package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	scanner := bufio.NewScanner(os.Stdin)
	//adjust the capacity to your need (max characters in line)
	const maxCapacity = 600 * 1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	scanner.Scan()
	strArr := strings.Fields(scanner.Text())
	for i, v := range strArr {
		arr[i], _ = strconv.Atoi(v)
	}
	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())
	a, b := solve(arr, k)
	if a == -1 {
		fmt.Print("None")
	} else {
		fmt.Printf("%v %v", arr[a], arr[b])
	}
}

func solve(arr []int, k int) (int, int) {
	i := 0
	j := len(arr) - 1
	for {
		sum := arr[i] + arr[j]
		switch {
		case sum == k:
			return i, j
		case sum < k:
			i++
		case sum > k:
			j--
		}
		if i == j {
			return -1, -1
		}
	}
}
