package main

import "fmt"

func main() {
	fmt.Println(addGenerics(1, 2))
	fmt.Println(addGenerics(1.2, 2.3))

	fmt.Println(addGenerics1(1.2, 2.3))

	fmt.Println(max(1.2, 2.3))
	fmt.Println(max("b", "a"))
}

type typeSet interface {
	int | float64 | string
}

func addGenerics[T int | float64](a, b T) T {
	return a + b
}

func max[T int | float64 | string](a, b T) T {
	if a > b {
		return a
	}

	return b
}

func addGenerics1[T typeSet](a, b T) T {
	return a + b
}
