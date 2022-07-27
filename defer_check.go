package main

import (
	"fmt"
)

func sayName(name string) {
	fmt.Println(name)
}

func someFunc() {
	name := "Hello"
	defer sayName(name)
	name = "World"
}

func main() {
	name := "Hello"

	defer func(name string) {
		fmt.Println(name)
	}(name)

	name = "world"
}
