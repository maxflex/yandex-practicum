package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Create("output.txt")
	defer func() {
		f.Close()
	}()
	n := 1000000
	f.Write([]byte(fmt.Sprint(n)))
	f.Write([]byte{'\n'})
	for i := 0; i < n; i++ {
		f.Write([]byte("999900"))
		if i+1 < n {
			f.Write([]byte{' '})
		} else {
			f.Write([]byte{'\n'})
		}
	}
	f.Write([]byte("999900"))
}
