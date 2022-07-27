package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProduct([]int{-2, 3, -4})) // 24
}

// [-2, 3, -4]

/**
-2, 3, -4
-2 -2


*/

// 3 || -6

func maxProduct(nums []int) int {
	minSoFar := nums[0]
	maxSoFar := nums[0]
	result := maxSoFar

	for i := 1; i < len(nums); i++ {
		curr := nums[i]

		tempMax := max(curr, max(maxSoFar*curr, minSoFar*curr))
		minSoFar = min(curr, min(maxSoFar*curr, minSoFar*curr))
		maxSoFar = tempMax
		result = max(maxSoFar, result)
	}

	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func maxBestSubArray(nums []int) int {
	return findBestSubarray(0, len(nums)-1, nums)
}

func findBestSubarray(left, right int, nums []int) int {
	if left > right {
		return math.MinInt
	}

	mid := (left + right) / 2
	curr := 0
	bestLeftSum := 0
	bestRightSum := 0

	for i := mid - 1; i >= left; i-- {
		curr += nums[i]
		bestLeftSum = max(bestLeftSum, curr)
	}

	curr = 0
	for i := mid + 1; i <= right; i++ {
		curr += nums[i]
		bestRightSum = max(bestRightSum, curr)
	}

	bestCombinedSum := nums[mid] + bestLeftSum + bestRightSum

	leftHalf := findBestSubarray(left, mid-1, nums)
	rightHalf := findBestSubarray(mid+1, right, nums)

	return max(bestCombinedSum, max(leftHalf, rightHalf))
}
