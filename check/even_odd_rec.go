package main

import "strconv"

var cache = make(map[string]bool)

func oddEvenJumps2(arr []int) int {
	var num int

	for i, _ := range arr {
		if isValidIndex2(i, arr, true) {
			num++
		}
	}

	cache = make(map[string]bool)
	return num
}

func isValidIndex2(i int, arr []int, oddJump bool) bool {
	if res, ok := cache[strconv.Itoa(i)+strconv.FormatBool(oddJump)]; ok {
		return res
	}
	if i == len(arr)-1 {
		return true
	}

	var nextIndex int
	var stepPossible bool

	if oddJump == true {
		nextIndex, stepPossible = jumpOdd(i, arr)
	} else {
		nextIndex, stepPossible = jumpEven(i, arr)
	}

	if !stepPossible {
		return false
	}

	oddJump = !oddJump

	res := isValidIndex2(nextIndex, arr, oddJump)
	cache[strconv.Itoa(nextIndex)+strconv.FormatBool(oddJump)] = res

	return res
}
