package main

func main() {
	generate(5)
}

func generate(numRows int) [][]int {
	if numRows == 1 {
		return [][]int{{1}}
	}

	if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}

	res := generate(numRows - 1)
	lastItem := res[len(res)-1]
	var newItem []int

	newItem = append(newItem, 1)

	for i := 0; i < len(lastItem)-1; i++ {
		newItem = append(newItem, lastItem[i]+lastItem[i+1])
	}

	newItem = append(newItem, 1)

	return append(res, newItem)
}
