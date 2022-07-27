package main

func main() {
	merge([]int{2, 0}, 1, []int{1}, 1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	if m == 0 {
		for i, _ := range nums2 {
			if i < n {
				nums1[i] = nums2[i]
			}
		}
	}

	cacheNums1 := cacheSlice(nums1)
	currentM := 0
	currentN := 0

	for currentM < m || currentN < n {
		if currentM < m && (currentN == n || cacheNums1[currentM] <= nums2[currentN]) {
			nums1[currentM+currentN] = cacheNums1[currentM]
			currentM++
		}

		if currentN < n && (currentM == m || cacheNums1[currentM] >= nums2[currentN]) {
			nums1[currentM+currentN] = nums2[currentN]
			currentN++
		}
	}
}

func cacheSlice(nums1 []int) []int {
	res := make([]int, len(nums1))
	for i, v := range nums1 {
		res[i] = v
	}

	return res
}
