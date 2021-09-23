package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func slidingWindow(n uint, k uint, arr []uint16) {
	writer := bufio.NewWriter(os.Stdin)
	res := make([]float64, 0, n-k+1)
	var val uint
	for _, v := range arr[0:k] {
		val += uint(v)
	}

	res = append(res, float64(val)/float64(k))

	for i := 1; i < cap(res); i++ {
		val = val - uint(arr[i-1]) + uint(arr[i+int(k)-1])
		res = append(res, float64(val)/float64(k))
	}

	for _, v := range res {
		fmt.Fprintf(writer, "%v ", v)
	}

	writer.Flush()
}

func main() {
	var n, k uint
	reader := bufio.NewReaderSize(os.Stdin, 4)

	fmt.Scan(&n)

	arr := make([]uint16, 0, n)
	var buf string
	for {
		byte, _ := reader.ReadByte()
		if byte == '\n' || byte == ' ' {
			i, _ := strconv.Atoi(buf)
			arr = append(arr, uint16(i))
			if byte == ' ' {
				buf = ""
			} else {
				break
			}
		} else {
			buf += string(byte)
		}
	}
	fmt.Scan(&k)
	slidingWindow(n, k, arr)
}
