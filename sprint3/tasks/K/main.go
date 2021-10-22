package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func merge(arr []int, left int, mid int, right int) (result []int) {
	arrLeft := arr[left:mid]
	arrRight := arr[mid:right]
	i, j := 0, 0
	for i < len(arrLeft) && j < len(arrRight) {
		if arrLeft[i] <= arrRight[j] {
			result = append(result, arrLeft[i])
			i++
		} else {
			result = append(result, arrRight[j])
			j++
		}
	}
	for i < len(arrLeft) {
		result = append(result, arrLeft[i])
		i++
	}
	for j < len(arrRight) {
		result = append(result, arrRight[j])
		j++
	}
	return result
}

func merge_sort(arr []int, lf int, rg int) {
	if lf == rg-1 {
		return
	}
	mid := (lf + rg) / 2
	merge_sort(arr, lf, mid)
	merge_sort(arr, mid, rg)
	result := merge(arr, lf, mid, rg)
	for i := 0; i < len(result); i++ {
		arr[lf+i] = result[i]
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	words := strings.Fields(scanner.Text())
	items := make([]int, len(words))
	for i, word := range words {
		v, _ := strconv.Atoi(word)
		items[i] = v
	}
	fmt.Println(items)
	merge_sort(items, 0, len(items))
	fmt.Println(items)
}
