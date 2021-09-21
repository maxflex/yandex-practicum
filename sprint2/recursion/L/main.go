package main

import "fmt"

func solve(n int) uint64 {
	if n < 2 {
		return 1
	}
	var (
		q      Queue
		i      int
		result uint64
	)
	q.init(2)
	q.enqueue(1)
	q.enqueue(1)
	for i = 2; i <= n; i++ {
		popped, _ := q.dequeue()
		peek, _ := q.peek()
		result = popped + peek
		if i != n {
			q.enqueue(result)
		}
	}
	return result
}

func pow(base uint64, exponent uint8) uint64 {
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
	for i := 68; i <= 98; i++ {
		solve := solve(i)
		fmt.Printf("i: %v\n%v\n====================\n", i, solve)
	}
	// var (
	// 	n int
	// 	k uint8
	// )
	// fmt.Scanf("%d %d", &n, &k)
	// fmt.Printf("n: %v | k: %v | solve: %v | pow: %v\n", n, k, solve(n), pow(10, k))
	// fmt.Println(solve(n) % pow(10, k))
}

/**
https://play.golang.org/p/s_y6TJByqZR
a := big.NewInt(4660046610375530309)
	b := big.NewInt(7540113804746346429)
	result := a.Add(a, b)
	fmt.Println(result)
	result = result.Add(result, b)
	fmt.Println(result)
	mod := big.NewInt(10000)
	fmt.Println(result.Mod(result, mod))
*/