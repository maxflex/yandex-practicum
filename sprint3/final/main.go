package main

import "fmt"

func brokenSearch(arr []int, k int) int {
	l := 0
	r := len(arr) - 1
	offset := getOffset(arr)
	mid := (l + r) / 2
	offsetMid := (mid + offset) % len(arr)
	for arr[offsetMid] != k {
		if k > arr[offsetMid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
		if l > r {
			return -1
		}
		mid = (l + r) / 2
		offsetMid = (mid + offset) % len(arr)
	}
	return offsetMid
}

func getOffset(arr []int) int {
	l := 0
	r := len(arr) - 1
	for {
		mid := (l + r) / 2
		next := (mid + 1) % len(arr)
		if arr[mid] > arr[next] {
			return next
		}
		if arr[l] > arr[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
		if l > r {
			return -1
		}
	}
}

func main() {
	// arr := []int{19, 21, 100, 101, 1, 4, 5, 7, 12}
	arr := []int{18, 19, 20, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
	result := brokenSearch(arr, 19)
	fmt.Println(result)
}
