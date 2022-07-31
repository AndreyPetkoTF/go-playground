package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(addGenerics(1, 2))
	fmt.Println(addGenerics(1.2, 2.3))

	fmt.Println(addGenerics1(1.2, 2.3))

	fmt.Println(max(1.2, 2.3))
	fmt.Println(max("b", "a"))

	data := make(map[int]int)
	now := time.Now()

	for i := 0; i < 10000; i++ {
		data[i] = i
		sum(data)
		sumIntOrFloat[int, int](data)
	}

	fmt.Println(time.Since(now).Seconds())

	PrintName(func() *User {
		return &User{"Hello"}
	})

	gooses := []Goose{
		{1, "goose1"},
		{2, "goose1"},
	}

	res := ToMap[Goose, int](gooses, func(g Goose) int { return g.ID })
	fmt.Println(res)
}

type Goose struct {
	ID   int
	name string
}

type NameGetter interface {
	Name() string
}

type User struct {
	name string
}

func (u User) Name() string {
	return u.name
}

func PrintName[T NameGetter](lazy func() T) {
	fmt.Println(lazy().Name())
}

func sum(data map[int]int) int {
	sum := 0
	for _, item := range data {
		sum += item
	}

	return sum
}

type Number interface {
	int | float64
}

func sumIntOrFloat[K comparable, V Number](data map[K]V) V {
	var sum V

	for _, item := range data {
		sum += item
	}

	return sum
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

func ToMap[T any, K comparable](items []T, keySelector func(T) K) map[K]T {
	var result = make(map[K]T, len(items))

	for _, item := range items {
		result[keySelector(item)] = item
	}

	return result
}
