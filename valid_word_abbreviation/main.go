package main

import (
	"fmt"
	"strconv"
)

func main() {
	//res := validWordAbbreviation("internationalization", "i12iz4n")
	//res := validWordAbbreviation("a", "2")
	//res := validWordAbbreviation("internationalization", "i5a11o1")

	//res := validWordAbbreviation("aa", "a1")
	res := validWordAbbreviation("hi", "hi1")

	fmt.Println(res)
}

/**
  isNumber | currentNumber

*/

func validWordAbbreviation(word string, abbr string) bool {
	if word == abbr {
		return true
	}

	wordI := 0
	abbrI := 0
	currentNumber := 0

	for wordI < len(word) && abbrI < len(abbr) {
		v, err := strconv.Atoi(string(abbr[abbrI]))
		isNumber := err == nil

		if isNumber && currentNumber == 0 && v == 0 {
			return false
		}

		if !isNumber && currentNumber == 0 { // previous not number and current items are equal
			if word[wordI] != abbr[abbrI] {
				return false
			} else {
				wordI++
				abbrI++
				continue
			}
		}

		if isNumber && currentNumber == 0 {
			currentNumber += v
			abbrI++

			if len(abbr) == abbrI {
				wordI += currentNumber
				currentNumber = 0
				if wordI > len(word) {
					return false
				}
			}

			continue
		}

		if isNumber && currentNumber != 0 { // current is number and previous number
			currentNumber *= 10
			currentNumber += v
			abbrI++

			if len(abbr) == abbrI {
				wordI += currentNumber
				currentNumber = 0
				if wordI > len(word) {
					return false
				}
			}

			continue
		}

		/**
		internationalization
		i12iz4n
		*/
		if !isNumber && currentNumber != 0 { // current is not number but previous was number
			wordI += currentNumber
			currentNumber = 0
			if wordI > len(word) {
				return false
			}
		}
	}

	return wordI == len(word) && abbrI == len(abbr)
}
