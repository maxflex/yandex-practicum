package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Item struct {
	a, b, c, d int
}

func (item Item) String() string {
	return fmt.Sprintf("%d %d %d %d", item.a, item.b, item.c, item.d)
}

type ResultSet struct {
	items []Item
}

func (r *ResultSet) push(item Item) {
	if r.exists(item) {
		return
	}
	r.items = append(r.items, item)
}

func (r ResultSet) exists(item Item) bool {
	for _, v := range r.items {
		if v.a == item.a && v.b == item.b && v.c == item.c && v.d == item.d {
			return true
		}
	}
	return false
}

func (r ResultSet) String() (result string) {
	result += fmt.Sprintf("%v\n", len(r.items))
	for _, v := range r.items {
		result += fmt.Sprintf("%v\n", v)
	}
	return
}

func partition(arr []int, i, j int) int {
	pivot := arr[(i+j)/2] // берём середину
	for j >= i {
		for arr[i] < pivot {
			i++
		}
		for arr[j] > pivot {
			j--
		}
		if j >= i {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}
	return i
}

func quickSort(arr []int, i, j int) {
	if j-i <= 0 {
		return
	}
	mid := partition(arr, i, j)
	quickSort(arr, i, mid-1)
	quickSort(arr, mid, j)
}

func main() {
	var n, x int
	fmt.Scan(&n, &x)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer([]byte{}, 1024*1024)
	scanner.Scan()
	words := strings.Fields(scanner.Text())
	arr := make([]int, n)
	for i, word := range words {
		v, _ := strconv.Atoi(word)
		arr[i] = v
	}
	quickSort(arr, 0, n-1)
	var res ResultSet
	for i := 0; i < n-1; i++ {
		for j := n - 1; j > i; j-- {
			l := i + 1
			r := j - 1
			for l < r {
				sum := arr[l] + arr[r] + arr[i] + arr[j]
				// fmt.Println(i, l, r, j, "\t", sum)
				if sum < x {
					l++
				} else if sum > x {
					r--
				} else {
					item := Item{arr[i], arr[l], arr[r], arr[j]}
					res.push(item)
					l++
				}
			}
		}
	}
	fmt.Println(res)
}
