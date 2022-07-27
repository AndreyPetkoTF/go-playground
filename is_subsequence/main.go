package main

import "fmt"

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
}

// abc
// aahbgdc

func isSubsequence(s string, t string) bool {
	for i := 0; i < len(t); i++ {
		if t[i] == s[0] {
			s = s[1:]
		}
	}

	return len(s) == 0
}
