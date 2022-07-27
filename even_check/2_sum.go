package main

//func main() {
//	fmt.Println(twoSum([]int{2, 3, 4}, 6))
//}

func twoSum(nums []int, target int) []int {
	leftPointer := 0
	rightPointer := len(nums) - 1

	for leftPointer < rightPointer {
		if nums[leftPointer]+nums[rightPointer] == target {
			return []int{leftPointer, rightPointer}
		}

		if nums[leftPointer]+nums[rightPointer] > target {
			rightPointer--
		} else {
			leftPointer++
		}
	}

	return []int{0, 0}
}
