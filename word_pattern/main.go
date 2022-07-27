package main

import (
	"fmt"
	"strings"
)

func main() {
	//fmt.Println(wordPattern("abba", "dog dog dog dog"))
	fmt.Println(wordPattern("aaa", "aa aa"))
}

func wordPattern(pattern string, s string) bool {
	chars := strings.Split(pattern, "")
	words := strings.Split(s, " ")

	if len(chars) != len(words) {
		return false
	}

	charToWord := make(map[string]string)
	wordToChar := make(map[string]string)

	/**
	  [a] => dog
	  [b] => cat

	 if (in map a) {
		get value of map -> dog
		if partsS[i] == a then partsPattern == dog
		if no return false
	}

	return true

	*/

	for i, char := range chars {
		if _, ok := charToWord[char]; !ok {
			charToWord[char] = words[i]
		} else {
			expectedWord := charToWord[char]

			if words[i] != expectedWord {
				return false
			}
		}
	}

	for i, word := range words {
		if _, ok := wordToChar[word]; !ok {
			wordToChar[word] = chars[i]
		} else {
			expectedChar := wordToChar[word]

			if chars[i] != expectedChar {
				return false
			}
		}
	}

	return true
}
