package main

import "fmt"

func main() {
	moveZeroes([]int{0, 1, 0})
}

func moveZeroes(nums []int) {
	//zerosFound := 0

	res := make([]int, len(nums))

	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			res[j] = nums[i]
			j++
		}
	}

	for i := 0; i < len(nums); i++ {
		nums[i] = res[i]
	}

	fmt.Println(nums)
}
