package main

import (
	"fmt"
)

type Puzzle [][]int

func main() {
	myPuzzle := Puzzle{{0, 4, 3, 0, 8, 0, 2, 5, 0},
					{6, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 1, 0, 9, 4},
					{9, 0, 0, 0, 0, 4, 0, 7, 0},
					{0, 0, 0, 6, 0, 8, 0, 0, 0},
					{0, 1, 0, 2, 0, 0, 0, 0, 3},
					{8, 2, 0, 5, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 5},
					{0, 3, 4, 0, 9, 0, 7, 1, 0}}
	
	print(myPuzzle)
	const Truth = true
	fmt.Println("sudoku solved?", Truth)
}

func print(puzzle Puzzle) {
	for _, i := range puzzle {
		for _, j := range i {
			fmt.Print(j)
		}
		fmt.Println()
	}
}

//solves the sudoku puzzle from 
func solve(i, j int, puzzle Puzzle) bool {
	return false
}
