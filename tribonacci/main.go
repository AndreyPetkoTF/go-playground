package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(tribonacci(4))
	fmt.Println(time.Since(now).Seconds())
}

// t(n) = t(n-1) + t(n-2) + t(n-3)

//4 -> 4
// 3 -> 3
func tribonacci(n int) int {
	cache := make(map[int]int)
	cache[0] = 0
	cache[1] = 1
	cache[2] = 1

	for i := 3; i <= n; i++ {
		// 2 + 1 + 0 = 3
		// 4 = 3 + 2 + 1

		cache[i] = cache[i-1] + cache[i-2] + cache[i-3]
	}

	return cache[n]
}
