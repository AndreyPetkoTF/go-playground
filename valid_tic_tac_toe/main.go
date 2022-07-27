package main

import "fmt"

func main() {
	fmt.Println(validTicTacToe([]string{"O  ", "   ", "   "})) // false
}

func validTicTacToe(board []string) bool {
	var movesX [][]int
	var movesO [][]int

	boardSize := 3

	for i, v := range board {
		for j := 0; j < boardSize; j++ {
			value := string(v[j])

			if value == "X" {
				movesX = append(movesX, []int{i, j})
			}

			if value == "O" {
				movesO = append(movesO, []int{i, j})
			}
		}
	}

	if len(movesX) < len(movesO) || len(movesX)-len(movesO) > 1 {
		return false
	}

	if isWin(movesX) && isWin(movesO) {
		return false
	}

	return true
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

/**
X   X   X
X   O   O
O   O   _



*/
