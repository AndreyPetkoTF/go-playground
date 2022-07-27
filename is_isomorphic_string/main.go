package main

import (
	"fmt"
	"reflect"
)

func main() {
	res := isIsomorphicSecondWay("badc", "baba")

	//res := isIsomorphic("paper", "title")

	fmt.Println(res)
}

func isIsomorphicSecondWay(s string, t string) bool {
	charToCharS := make(map[rune]rune)
	charToCharT := make(map[rune]rune)
	sRunes := []rune(s)
	tRunes := []rune(t)

	for i, v := range sRunes {
		if _, ok := charToCharS[v]; !ok {
			charToCharS[v] = tRunes[i]
		} else {
			if tRunes[i] != charToCharS[v] {
				return false
			}
		}
	}

	for i, v := range tRunes {
		if _, ok := charToCharT[v]; !ok {
			charToCharT[v] = sRunes[i]
		} else {
			if sRunes[i] != charToCharT[v] {
				return false
			}
		}
	}

	return true
}

func isIsomorphic(s string, t string) bool {
	itemToIndexesS := make(map[rune][]int)

	for i, v := range s {
		itemToIndexesS[v] = append(itemToIndexesS[v], i)
	}

	itemToIndexesT := make(map[rune][]int)
	for i, v := range t {
		itemToIndexesT[v] = append(itemToIndexesT[v], i)
	}

	for singleChar, indexesS := range itemToIndexesS {
		if len(indexesS) == 1 {
			delete(itemToIndexesS, singleChar)
			continue
		}

		runesT := []rune(t)
		indexesT := itemToIndexesT[runesT[indexesS[0]]]

		if reflect.DeepEqual(indexesS, indexesT) {
			delete(itemToIndexesS, singleChar)
			delete(itemToIndexesT, runesT[indexesS[0]])
		}
	}

	for i, v := range itemToIndexesT {
		if len(v) == 1 {
			delete(itemToIndexesT, i)
		}
	}

	return len(itemToIndexesS) == 0 && len(itemToIndexesT) == 0
}
