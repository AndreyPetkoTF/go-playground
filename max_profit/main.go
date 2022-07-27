package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProfit1([]int{2, 2, 5})) // 4
}

func maxProfit(prices []int) int {
	maxDiff := 0
	min := -1
	currentMaxIndex := 0
	currentMax := 0

	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] && min == -1 {
			min = prices[i]
			currentMax = prices[i+1]

			for j := i + 1; j < len(prices); j++ {
				if prices[j] >= currentMax {
					currentMax = prices[j]
					currentMaxIndex = j
				}
			}

			maxDiff = currentMax - prices[i]
		}

		if prices[i] < min {
			if i < currentMaxIndex {
				maxDiff = currentMax - prices[i]
			} else {
				newMax := 0
				for j := i + 1; j < len(prices); j++ {
					if prices[j] > newMax {
						newMax = prices[j]
					}
				}

				newDiff := newMax - prices[i]

				if maxDiff < newDiff {
					maxDiff = newDiff
				}
			}

			min = prices[i]
		}
	}

	return maxDiff
}

// 2 2 5
func maxProfit1(prices []int) int {
	minPrice := math.MaxInt
	maxProfitInt := 0

	for i := 0; i < len(prices); i++ {
		if minPrice > prices[i] {
			minPrice = prices[i]
		} else {
			if maxProfitInt < prices[i]-minPrice {
				maxProfitInt = prices[i] - minPrice
			}
		}
	}

	return maxProfitInt
}
