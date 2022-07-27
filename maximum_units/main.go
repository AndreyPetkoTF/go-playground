package main

import (
	"fmt"
	"sort"
)

func main() {
	//res := maximumUnitsCustomSort([][]int{
	//	{1, 3},
	//	{5, 5},
	//	{2, 5},
	//	{4, 2},
	//	{4, 1},
	//	{3, 1},
	//	{2, 2},
	//	{1, 3},
	//	{2, 5},
	//	{3, 2},
	//}, 35)

	res := maximumUnitsCustomSort([][]int{
		{1, 3},
		{2, 2},
		{3, 1},
	}, 4)

	fmt.Println(res)
}

func maximumUnitsCustomSort(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})

	res := 0
	for _, v := range boxTypes {
		availableAmount := v[0]

		for availableAmount != 0 && truckSize != 0 {
			res += v[1]
			availableAmount--
			truckSize--

			if truckSize == 0 {
				break
			}
		}
	}

	return res
}

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sizeToAmount := make(map[int]int)

	var sizes []int
	for _, v := range boxTypes {
		if _, ok := sizeToAmount[v[1]]; !ok {
			sizes = append(sizes, v[1])
			sizeToAmount[v[1]] = v[0]
		} else {
			sizeToAmount[v[1]] += v[0]
		}
	}

	sort.Ints(sizes)

	res := 0

	for i := len(sizes) - 1; i >= 0; i-- {
		size := sizes[i]
		availableAmount := sizeToAmount[size]

		for availableAmount != 0 && truckSize != 0 {
			res += size
			availableAmount--
			truckSize--

			if truckSize == 0 {
				break
			}
		}
	}

	return res
}
