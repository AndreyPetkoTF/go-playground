package main

import "fmt"

func main() {
	game := Constructor(2)

	fmt.Println(game.Move(0, 1, 2))
	fmt.Println(game.Move(1, 0, 1))
	fmt.Println(game.Move(1, 1, 2))

	fmt.Println(1)
	/**
	  	O
	    X O
	*/

}

type TicTacToe struct {
	Size    int
	sameHX  map[int]int
	sameWX  map[int]int
	sameHO  map[int]int
	sameWO  map[int]int
	sameDX  int
	sameDO  int
	sameD2X int
	sameD2O int
}

func Constructor(n int) TicTacToe {
	return TicTacToe{
		Size:   n,
		sameHX: make(map[int]int),
		sameWX: make(map[int]int),
		sameHO: make(map[int]int),
		sameWO: make(map[int]int),
	}
}

func (this *TicTacToe) Move(row int, col int, player int) int {
	var sameH map[int]int
	var sameW map[int]int

	if player == 1 {
		sameH = this.sameHX
		sameW = this.sameWX
	} else {
		sameH = this.sameHO
		sameW = this.sameWO
	}

	sameH[row]++
	sameW[col]++

	if sameH[row] == this.Size || sameW[col] == this.Size {
		return player
	}

	if row == col {
		if player == 1 {
			this.sameDX++

			if this.sameDX == this.Size {
				return player
			}
		} else {
			this.sameDO++

			if this.sameDO == this.Size {
				return player
			}
		}
	}

	if row+col == this.Size-1 {
		if player == 1 {
			this.sameD2X++

			if this.sameD2X == this.Size {
				return player
			}
		} else {
			this.sameD2O++

			if this.sameD2O == this.Size {
				return player
			}
		}
	}

	return 0
}
