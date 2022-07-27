package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

func NewItem(a, b, c int) Item {
	sorted := []int{a, b, c}
	sort.Ints(sorted)

	return Item{sorted[0], sorted[1], sorted[2]}
}

type Item struct {
	a, b, c int
}

func threeSum(nums []int) [][]int {
	resSet := make(map[Item]struct{})
	seen := make(map[int]int)

	for i, v := range nums {

		for _, v2 := range nums[i+1:] {
			needed := -v - v2

			if k, ok := seen[needed]; ok && k == i {
				resSet[NewItem(v, v2, needed)] = struct{}{}
			}

			seen[v2] = i
		}
	}

	var res [][]int
	for k, _ := range resSet {
		res = append(res, []int{k.a, k.b, k.c})
	}

	return res
}

//func threeSum(nums []int) [][]int {
//	sort.Ints(nums)
//
//	var res [][]int
//
//	for i := 0; i < len(nums); i++ {
//		if i == 0 || nums[i-1] != nums[i] {
//			res = twoSumII(nums, i, res)
//		}
//	}
//
//	return res
//}

//func twoSumII(nums []int, i int, res [][]int) [][]int {
//	leftPointer := i + 1
//	rightPointer := len(nums) - 1
//
//	for leftPointer < rightPointer {
//		sum := nums[i] + nums[leftPointer] + nums[rightPointer]
//
//		if sum < 0 {
//			leftPointer++
//		} else if sum > 0 {
//			rightPointer--
//		} else {
//			res = append(res, []int{nums[i], nums[leftPointer], nums[rightPointer]})
//			leftPointer++
//			rightPointer--
//
//			for leftPointer < rightPointer && nums[leftPointer] == nums[leftPointer-1] {
//				leftPointer++
//			}
//		}
//	}
//
//	return res
//}
