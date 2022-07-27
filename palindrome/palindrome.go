package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(romanToInt("MCMXCIV")) // 1994
}

var romanMap = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	var res int

	for i := 0; i < len(s); i++ {
		if string(s[i]) == "I" && i != len(s)-1 {
			if string(s[i+1]) == "V" || string(s[i+1]) == "X" {
				res -= 2
			}
		}

		// 10 + 50 = 40 || 10 + 100 = 90

		if string(s[i]) == "X" && i != len(s)-1 {
			if string(s[i+1]) == "L" || string(s[i+1]) == "C" {
				res -= 20
			}
		}

		//
		if string(s[i]) == "C" && i != len(s)-1 {
			if string(s[i+1]) == "D" || string(s[i+1]) == "M" {
				res -= 200
			}
		}

		if number, ok := romanMap[string(s[i])]; ok {
			res += number
		}
	}

	return res
}

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)

	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}

	return true
}
