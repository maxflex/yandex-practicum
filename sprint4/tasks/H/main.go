package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer([]byte{}, 1024*1024)
	scanner.Scan()
	s1 := scanner.Text()
	scanner.Scan()
	s2 := scanner.Text()
	if len(s1) != len(s2) {
		fmt.Print("NO")
		return
	}
	m := make(map[byte]byte)
	for i := 0; i < len(s1); i++ {
		if val, ok := m[s1[i]]; ok {
			if val != s2[i] {
				fmt.Print("NO")
				return
			}
		} else {
			for _, v := range m {
				if v == s2[i] {
					fmt.Print("NO")
					return
				}
			}
			m[s1[i]] = s2[i]
		}
	}
	fmt.Print("YES")
}
