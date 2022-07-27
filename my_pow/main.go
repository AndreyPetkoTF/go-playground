package main

import "fmt"

func main() {
	//fmt.Println(myPow(2.00000, -1))
	fmt.Println(fastPowMain(2, 4))
}

func myPow(x float64, n int) float64 {
	var res float64 = 1

	for i := 0; i < abs(n); i++ {
		res *= x
	}

	if n >= 0 {
		return res
	} else {
		return float64(1) / res
	}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func fastPowMain(x float64, n int) float64 {
	if n < 0 {
		n = -n
		x = 1 / x
	}

	return fastPow(x, n)
}

func fastPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	half := fastPow(x, n/2)
	if n%2 == 0 {
		return half * half
	} else {
		return half * half * x
	}
}