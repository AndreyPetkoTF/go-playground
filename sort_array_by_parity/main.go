package main

func main() {

}

func sortArrayByParity(nums []int) []int {
	var odd []int
	var even []int

	for _, v := range nums {
		if v%2 == 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}

	return append(odd, even...)
}
