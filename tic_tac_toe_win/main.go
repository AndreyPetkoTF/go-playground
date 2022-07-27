package main

import "fmt"

func main() {
	fmt.Println(tictactoe([][]int{
		{0, 0}, {1, 1}, {2, 0}, {1, 0}, {1, 2}, {2, 1}, {0, 1}, {0, 2}, {2, 2},
	}))
}

func tictactoe(moves [][]int) string {
	var movesA [][]int
	var movesB [][]int

	for i, v := range moves {
		if i%2 == 0 {
			movesA = append(movesA, v)
		} else {
			movesB = append(movesB, v)
		}
	}

	if isWin(movesA) {
		return "A"
	}

	if isWin(movesB) {
		return "B"
	}

	if len(moves) == 9 {
		return "Draw"
	}

	return "Pending"
}

func isWin(moves [][]int) bool {
	sameH := make(map[int]int)
	sameW := make(map[int]int)
	boardSize := 3

	sameHW := 0
	anotherDiag := 0

	for _, v := range moves {
		sameH[v[0]]++
		sameW[v[1]]++

		if v[0] == v[1] {
			sameHW++

			if sameHW == boardSize {
				return true
			}
		}

		if v[0]+v[1] == boardSize-1 {
			anotherDiag++

			if anotherDiag == boardSize {
				return true
			}
		}

		if sameH[v[0]] == boardSize || sameW[v[1]] == boardSize {
			return true
		}
	}

	return false
}
