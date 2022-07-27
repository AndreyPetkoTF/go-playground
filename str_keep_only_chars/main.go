package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama")) // true
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)

	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")

	processedString := reg.ReplaceAllString(s, "")

	for i := 0; i <= len(processedString)/2; i++ {
		if processedString[i] != processedString[len(processedString)-1-i] {
			return false
		}
	}

	return true
}
