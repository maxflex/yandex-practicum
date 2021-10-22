package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CacheKey struct {
	l, r int
}

type Cache map[CacheKey]int

var (
	s     string
	q, m  int
	cache Cache
)

// метод Гонера: ((s1*q + s2)*q + s3) % M
func h(l, r int) (h int) {
	key := CacheKey{l, r}
	if val, ok := cache[key]; ok {
		return val
	}
	for i := l; i <= r; i++ {
		h = (h*q + int(s[i])) % m
	}
	cache[key] = h
	return
}

func main() {
	var l, r int
	cache = make(Cache)
	fmt.Scan(&q, &m)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer([]byte{}, 1024*1024)
	scanner.Scan()
	s = scanner.Text()
	scanner.Scan()
	for scanner.Scan() {
		text := strings.Fields(scanner.Text())
		l, _ = strconv.Atoi(text[0])
		r, _ = strconv.Atoi(text[1])
		fmt.Println(h(l-1, r-1))
	}
}
