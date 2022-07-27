package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(repeatedSubstringPattern("abab"))         // true
	fmt.Println(repeatedSubstringPattern("aba"))          // false
	fmt.Println(repeatedSubstringPattern("abcabcabcabc")) // true
}

func repeatedSubstringPattern(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		res := strings.Split(s, s[0:i+1])

		if len(res) > 1 && emptyItems(res) {
			return true
		}
	}

	return false
}

func emptyItems(res []string) bool {
	for _, v := range res {
		if v != "" {
			return false
		}
	}

	return true
}
