package main

import (
	"fmt"
	"math/rand"
	"time"
)

func partition(arr []int, i int, j int) int {
	l := i
	r := j
	// случайный индекс
	pivotIndex := i + rand.Intn(j-i)
	pivot := arr[pivotIndex]
	fmt.Println("pivot", pivot, "pivotIndex", pivotIndex)
	fmt.Println("i", i, "j", j)
	fmt.Println("START")
	fmt.Println(arr[i:j])
	for j >= i {
		for arr[i] < pivot {
			i++
			fmt.Println("i++", i)
		}
		for arr[j] > pivot {
			j--
			fmt.Println("j--", j)
		}
		if j >= i {
			fmt.Println("swap", arr[i], "and", arr[j])
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
			fmt.Println(arr[l:r])
			fmt.Println("i++", i, "j--", j)
		}
	}
	fmt.Println("FINISH")
	fmt.Println(arr[l:r])
	fmt.Println()
	return i
}

func quickSort(arr []int, i int, j int) {
	fmt.Println(i, j, j-i < 2)
	if j-i < 1 {
		return
	}
	mid := partition(arr, i, j)
	// fmt.Printf("%v | i: %v | j: %v | mid: %v | sort left: %v | sort right: %v\n", j-i < 2, i, j, mid, i < mid-1, mid < j)
	// if i < mid-1 {
	quickSort(arr, i, mid-1)
	// }
	// if mid < j {
	quickSort(arr, mid, j)
	// }
}

func main() {
	rand.Seed(time.Now().Unix())
	arr := []int{-5, -10, 2, 3, 1, 3}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
