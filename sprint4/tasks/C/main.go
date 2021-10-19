package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	s       string
	q, m    int
	precalc []int
)

// метод Гонера: ((s1*q + s2)*q + s3) % M
func h(l, r int) (h int) {
	for i := l; i <= r; i++ {
		h = (h*q + int(s[i]) - 94) % m
	}
	return
}

func pow(x, n int) int {
	switch n {
	case 0:
		return 1
	case 1:
		return x
	}
	return x * pow(x, n-1)
}

func main() {
	var l, r int
	fmt.Scan(&q, &m)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer([]byte{}, 1024*1024)
	scanner.Scan()
	s = scanner.Text()
	// for i := 0; i < len(s); i++ {
	// 	precalc = append(precalc, pow(2, 4))
	// }
	// fmt.Println("PRECALC", precalc)
	scanner.Scan()
	for scanner.Scan() {
		text := strings.Fields(scanner.Text())
		l, _ = strconv.Atoi(text[0])
		r, _ = strconv.Atoi(text[1])
		fmt.Printf("L: %d | R: %d | H: %d\n", l, r, h(l-1, r-1))
	}
}
