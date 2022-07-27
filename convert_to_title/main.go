package main

import (
	"fmt"
)

func main() {
	fmt.Println(convertToTitle(701))
}

func convertToTitle(number int) string {
	result := ""

	for {
		if number <= 26 {
			result = string(number+64) + result
			break
		} else {
			result = string((number-1)%26+1+64) + result
			number = (number - 1) / 26
		}
	}

	return result
}

/**

1-26 -> A-Z

each 26 -> A,B second (AA, AB, ...)
each 26*26 ->

*/
