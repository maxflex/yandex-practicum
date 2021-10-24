package main

import "fmt"

func main() {
	var n uint16
	var result string
	fmt.Scan(&n)
	for n != 0 {
		result += fmt.Sprint(n % 2)
		n = n / 2
	}
	for i := len(result) - 1; i >= 0; i-- {
		fmt.Print(string(result[i]))
	}
}
