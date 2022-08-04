package main

import "fmt"

func main() {
	res := longestPalindrome("abbccd")

	fmt.Println(res)
}

func longestPalindrome(s string) int {
	vToCount := make(map[string]int)

	for _, v := range s {
		vToCount[string(v)]++
	}

	sum := 0
	evenAdded := false

	for _, v := range vToCount {
		if v%2 == 0 {
			sum += v
		} else {
			if evenAdded == false {
				sum += v
				evenAdded = true
			} else {
				sum += v - 1
			}
		}
	}

	return sum
}