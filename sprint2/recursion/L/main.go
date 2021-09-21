package main

import (
	"fmt"
	"math/big"
)

func solve(n int) *big.Int {
	if n < 2 {
		return big.NewInt(1)
	}
	var q Queue
	var i int
	result := new(big.Int)
	q.init(2)
	q.enqueue(big.NewInt(1))
	q.enqueue(big.NewInt(1))
	for i = 2; i <= n; i++ {
		popped, _ := q.dequeue()
		peek, _ := q.peek()
		result = result.Add(popped, peek)
		if i != n {
			q.enqueue(result)
		}
	}
	return result
}

func pow(base int64, exponent uint8) int64 {
	if exponent == 0 {
		return 1
	}
	var i uint8
	result := base
	for i = 1; i < exponent; i++ {
		result *= base
	}
	return result
}

func main() {

	var (
		n int
		k uint8
	)
	fmt.Scanf("%d %d", &n, &k)
	// fmt.Printf("n: %v | k: %v | solve: %v | pow: %v\n", n, k, solve(n), pow(10, k))
	result := solve(n)
	fmt.Println(result.Mod(result, big.NewInt(pow(10, k))))
}
