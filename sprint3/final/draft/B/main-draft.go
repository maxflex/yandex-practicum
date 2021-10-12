package main_sketch

import (
	"fmt"
	"math/rand"
	"time"
)

type Participant struct {
	login   string
	score   int
	penalty int
}

func less(a Participant, b Participant) bool {
	if a.score != b.score {
		return a.score < b.score
	}
	if a.penalty != b.penalty {
		return a.penalty > b.penalty
	}
	return a.login > b.login
}

func partition(arr []Participant, i int, j int) int {
	// случайный индекс
	pivotIndex := i + rand.Intn(j-i)
	pivot := arr[pivotIndex]
	// fmt.Println("pivot", pivot, "i", i, "j", j)
	for j >= i {
		for less(pivot, arr[i]) {
			i++
			// fmt.Println("i++", i)
		}
		for less(arr[j], pivot) {
			j--
			// fmt.Println("j--", j)
		}
		if j >= i {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}
	return i
}

func quickSort(arr []Participant, i int, j int) {
	if j-i <= 0 {
		return
	}
	mid := partition(arr, i, j)
	quickSort(arr, i, mid-1)
	quickSort(arr, mid, j)
}

// func insertionSort(arr []Participant) {
// 	for i := 1; i < len(arr); i++ {
// 		for j := i; j >= 1; j-- {
// 			if compare(arr[j], arr[j-1]) > 0 {
// 				arr[j], arr[j-1] = arr[j-1], arr[j]
// 			} else {
// 				break
// 			}
// 		}
// 	}
// 	fmt.Println(arr)
// }

func main() {
	rand.Seed(time.Now().Unix())
	// arr := []Participant{
	// 	{"maxflex", 10, 2},
	// 	{"anton", 10, 3},
	// 	{"xavier", 20, 5},
	// }
	// [{anton 10 3} {maxflex 10 2} {xavier 20 5}]
	arr := []Participant{
		{"alla", 4, 100},
		{"gena", 6, 1000},
		{"gosha", 2, 90},
		{"rita", 2, 90},
		{"timofey", 4, 80},
	}
	// fmt.Println(greater(arr[4], arr[0]))
	// arr := []Participant{
	// 	{"alla", 0, 0},
	// 	{"gena", 0, 0},
	// 	{"gosha", 0, 0},
	// 	{"rita", 0, 0},
	// 	{"timofey", 0, 0},
	// }
	// gena
	// timofey
	// alla
	// gosha
	// rita
	// insertionSort(arr)
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
