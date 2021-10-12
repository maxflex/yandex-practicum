package main

import "fmt"

var step int

func find(arr []int, x int, l int, r int, offset int) int {
	step++
	mid := (l + r) / 2
	offsetMid := getOffsetIndex(arr, mid, offset)
	fmt.Printf("step: %v | l: %v | r: %v \nmid: %v | omid: %v \nvalue: %v\n\n", step, l, r, mid, offsetMid, arr[offsetMid])
	if step > 10 {
		panic("error")
	}
	if arr[offsetMid] == x {
		return offsetMid
	}
	if x < arr[offsetMid] {
		return find(arr, x, l, mid-1, offset)
	} else {
		return find(arr, x, mid+1, r, offset)
	}
}

func getOffset(arr []int, l int, r int) int {
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

func getOffsetIndex(arr []int, brokenIndex int, offset int) int {
	return (brokenIndex + offset) % len(arr)
}

func main() {
	// arr := []int{18, 19, 20, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
	// 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20
	// offset = 3
	// x = 19
	// l = 0; r = 19
	// mid = 9
	//
	// step 2
	// l = 10; r = 19
	// mid = 15; cmid = 18
	//
	// step 3
	// l = 16; r = 19
	// mid = 17; cmid = 0
	//
	// step 4
	// l = 18; r = 19
	// mid = 18; cmid = 1

	// ------- l = 3 | r = 2

	// arr := []int{19, 21, 100, 101, 1, 4, 5, 7, 12}
	// 1 4 5 7 12 19 21 100 101
	// OFFSET = 4; X=5; LEN(ARR) = 9
	//
	// step 1
	// l = 0; r = 8
	// mid = 4; omid = 8
	//
	// step 2
	// l = 0; r = 3
	// mid = 1; omid = 5
	//
	// step 3
	// l = 2; r = 3
	// mid = 2; omid = 6

	arr := []int{5, 1}

	offset := getOffset(arr, 0, len(arr)-1)
	fmt.Println("OFFSET", offset)
	// fmt.Println("getIndex(0)", getIndex(0, offset, arr))
	// fmt.Println("getIndex(LAST)", getIndex(len(arr)-1, offset, arr))
	// result := find(arr, 5, offset, getOffsetIndex(arr, len(arr)-1, offset), offset)
	result := find(arr, 5, 0, len(arr)-1, offset)
	fmt.Println(result)
}
