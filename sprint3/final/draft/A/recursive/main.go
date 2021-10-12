package main_recursive

import "fmt"

func brokenSearch(arr []int, k int) int {
	n := len(arr) - 1
	offset := getOffset(arr, 0, n)
	return find(arr, k, 0, n, offset)
}

func find(arr []int, k, l, r, offset int) int {
	mid := (l + r) / 2
	offsetMid := (mid + offset) % len(arr)
	switch {
	case arr[offsetMid] == k:
		return offsetMid
	case k < arr[offsetMid]:
		return find(arr, k, l, mid-1, offset)
	case k > arr[offsetMid]:
		return find(arr, k, mid+1, r, offset)
	default:
		return -1
	}
}

func getOffset(arr []int, l, r int) int {
	mid := (l + r) / 2
	if arr[(mid+1)%len(arr)] < arr[mid] {
		return (mid + 1) % len(arr)
	}
	if r > l {
		left := getOffset(arr, l, mid-1)
		if left != -1 {
			return left
		}
		return getOffset(arr, mid+1, r)
	}
	return -1
}

func main() {
	arr := []int{19, 21, 100, 101, 1, 4, 5, 7, 12}
	offset := getOffset(arr, 0, len(arr)-1)
	result := find(arr, 5, 0, len(arr)-1, offset)
	fmt.Println(result)
}
