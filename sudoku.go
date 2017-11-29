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

//solves the sudoku puzzle from i, j and returns true if it is solved, false if it cannot be solved.
func solve(i, j int, puzzle Puzzle) bool {
	if i == 9 {
		i = 0
		j++
		
		if j == 9 {
			return true
		}
	}

	// solve the blank cell
	if puzzle[i][j] == 0 {
		for testNum := 1; testNum <= 9; testNum++ {
			if valid(testNum, i, j) {
				puzzle[i][j] = testNum
				if (solve(i + 1, j, puzzle)) {
					return true
				}
			}
		}

		// Tried all 9 numbers. Reset it to 0 and backtrack.
		puzzle[i][j] = 0
		return false
	}

	// Space was already filled.
	return solve(i + 1, j, puzzle)
}

func valid(x, i, j int) bool {
	// NOT IMPLEMENTED
	return true;
}
