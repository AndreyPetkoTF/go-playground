package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	now := time.Now()
	//fmt.Println(oddEvenJumps2({]int{5, 1, 3, 4, 2}))
	//fmt.Println(oddEvenJumps2({]int{10, 13, 12, 14, 15})) // 2

	fmt.Println(isValidIndex(4, []int{10, 13, 12, 14, 15}))

	fmt.Println(time.Since(now).Seconds())
}

func oddEvenJumps(arr []int) int {
	var num int

	// reverse main loop and create cache for indexes
	for i, _ := range arr {
		if isValidIndex(i, arr) {
			num++
		}
	}

	return num
}

func isValidIndex(startIndex int, arr []int) bool {
	oddJump := true
	currentIndex := startIndex
	stepPossible := true

	for currentIndex != len(arr)-1 && stepPossible == true {
		if oddJump == true {
			currentIndex, stepPossible = jumpOdd(currentIndex, arr)
		} else {
			currentIndex, stepPossible = jumpEven(currentIndex, arr)
		}

		oddJump = !oddJump
	}

	return currentIndex == len(arr)-1
}

func jumpOdd(currentIndex int, arr []int) (int, bool) {
	minValueAboveCurrent := math.MaxInt
	var newIndex int

	for i := currentIndex + 1; i < len(arr); i++ {
		val := arr[i]
		if arr[currentIndex] <= val && val < minValueAboveCurrent {
			minValueAboveCurrent = val
			newIndex = i
		}
	}

	if newIndex != 0 {
		return newIndex, true
	}

	return -1, false
}

func jumpEven(currentIndex int, arr []int) (int, bool) {
	maxValueBelowCurrent := 0
	var newIndex int

	for i := currentIndex + 1; i < len(arr); i++ {
		val := arr[i]
		if arr[currentIndex] >= val && val > maxValueBelowCurrent {
			maxValueBelowCurrent = val
			newIndex = i
		}
	}

	if newIndex != 0 {
		return newIndex, true
	}

	return 0, false
}
