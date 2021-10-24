package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	kids := make([]int, 0, n)

	scanner.Scan()
	words := strings.Fields(scanner.Text())
	for _, word := range words {
		v, _ := strconv.Atoi(word)
		kids = append(kids, v)
	}
	sort.Ints(kids)

	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	cookies := make([]int, 0, m)

	scanner.Scan()
	words = strings.Fields(scanner.Text())
	for _, word := range words {
		v, _ := strconv.Atoi(word)
		cookies = append(cookies, v)
	}
	sort.Ints(cookies)

	var happy, j int
	for i := 0; i < len(kids) && j < len(cookies); i++ {
		for j < len(cookies) && cookies[j] < kids[i] {
			j++
		}
		if j < len(cookies) {
			happy++
			j++
		}
	}

	fmt.Print(happy)
}
