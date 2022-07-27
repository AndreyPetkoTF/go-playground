package main

import (
	"fmt"
	"sort"
)

func main() {
	res := kWeakestRowsGoSort([][]int{
		{1, 1, 0, 0, 0}, {1, 1, 1, 1, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 0, 0}, {1, 1, 1, 1, 1},
	}, 3)

	fmt.Println(res)
}

type entry struct {
	Index    int
	Soldiers int
}

func kWeakestRowsGoSort(mat [][]int, k int) []int {
	items := make([]entry, len(mat), len(mat))

	for i, row := range mat {
		soldiers := 0
		for _, v := range row {
			if v == 1 {
				soldiers++
			}
		}

		items[i] = entry{i, soldiers}
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].Soldiers < items[j].Soldiers
	})

	res := make([]int, k, k)
	for i := 0; i < k; i++ {
		res[i] = items[i].Index
	}

	return res
}

func kWeakestRows(mat [][]int, k int) []int {
	currentCol := 0
	maxColValue := len(mat[0]) - 1
	var indexes []int

	for {
		for i := 0; i < len(mat); i++ {
			if mat[i][currentCol] == 0 && !inSlice(indexes, i) {
				indexes = append(indexes, i)

				if len(indexes) == k {
					return indexesToRes(indexes, len(mat)-1, k)
				}
			}
		}

		if maxColValue == currentCol {
			return indexesToRes(indexes, len(mat)-1, k)
		}

		currentCol++
	}
}

func indexesToRes(indexes []int, maxRowIndex int, k int) []int {
	if len(indexes) != k {
		for i := 0; i <= maxRowIndex; i++ {
			if !inSlice(indexes, i) {
				indexes = append(indexes, i)

				if len(indexes) == k {
					return indexes
				}
			}
		}
	}

	return indexes
}

func inSlice(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}

	return false
}
