package main

import "fmt"

func main() {
	//fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	//fmt.Println(searchInsert([]int{1, 3, 4, 6}, 7)) // 1
	//fmt.Println(searchInsert([]int{1, 3, 4, 6}, 2)) // 1

	fmt.Println(searchInsert([]int{1, 3}, 0)) // 1
}

/**

t = 7
l   m r
1 3 5 6

l 0 m 2 r 3

l 0 m 1 r 3


*/

func searchInsert(nums []int, target int) int {
	mid := len(nums) / 2 // 2
	left := 0            //
	right := len(nums) - 1

	for left <= right {
		// find or not found
		if nums[mid] == target {
			return mid
		}

		if (mid == 0 || nums[mid-1] < target) && target < nums[mid] {
			return mid
		}

		if (mid == len(nums)-1 || target < nums[mid+1]) && nums[mid] < target {
			return mid + 1
		}

		if target < nums[mid] {
			right = mid - 1

			mid = (left+mid)/2 - 1

			if mid < 0 {
				mid = 0
			}
		} else {
			left = mid + 1
			mid = (right+mid)/2 + 1

			if mid > len(nums)-1 {
				mid = len(nums) - 1
			}
		}
	}

	return -1
}
