package main

import (
	"fmt"
)

func main() {
	fmt.Println(numIslands([][]byte{
		{49, 49, 49, 49, 48},
		{49, 49, 48, 49, 48},
		{49, 49, 48, 48, 48},
		{48, 48, 48, 48, 48},
	}))
}

func numIslands(grid [][]byte) int {
	var counter int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 49 {
				counter++
				removeIsland(grid, i, j)
			}
		}
	}

	return counter
}

func removeIsland(grid [][]byte, i int, j int) {
	if (i < 0 || i > len(grid)-1) || (j < 0 || j > len(grid[i])-1) || grid[i][j] == 48 {
		return
	}

	if grid[i][j] == 49 {
		grid[i][j] = 48
	}

	removeIsland(grid, i, j+1)
	removeIsland(grid, i, j-1)
	removeIsland(grid, i-1, j)
	removeIsland(grid, i+1, j)
}
