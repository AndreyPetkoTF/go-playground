package main

import "fmt"

func main() {
	variable := 5

	checkType(variable)
}

func checkType(variable interface{}) {
	switch res := variable.(type) {
	case int:
		fmt.Printf("%T", res)
	default:
		fmt.Printf("%T", res)
	}
}
