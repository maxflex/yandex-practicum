package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var k, score byte
	keyboard := make(map[byte]byte)
	reader := bufio.NewReader(os.Stdin)
	fmt.Scan(&k)
	k *= 2
	for {
		byte, err := reader.ReadByte()
		if err != nil {
			break
		}
		if byte == '\n' || byte == '.' {
			continue
		}
		keyboard[byte] += 1
	}
	for _, v := range keyboard {
		if v <= k {
			score += 1
		}
	}
	fmt.Print(score)
}
