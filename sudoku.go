package main

import (
	"fmt"
)

type Puzzle [9][9]int

func main() {
	myPuzzle := Puzzle{{0, 4, 3, 0, 8, 0, 2, 5, 0},
					{6, 0, 0, 0, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 1, 0, 9, 4},
					{9, 0, 0, 0, 0, 4, 0, 7, 0},
					{0, 0, 0, 6, 0, 8, 0, 0, 0},
					{0, 1, 0, 2, 0, 0, 0, 0, 3},
					{8, 2, 0, 5, 0, 0, 0, 0, 0},
					{0, 0, 0, 0, 0, 0, 0, 0, 5},
					{0, 3, 4, 0, 9, 0, 7, 1, 0},}
	
	isSolved, solved := solve(0, 0, myPuzzle)
	print(solved)
	fmt.Println("sudoku solved?", isSolved)
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
func solve(i, j int, puzzle Puzzle) (bool, Puzzle) {
	if i == 9 {
		i = 0
		j++
		
		if j == 9 {
			return true, puzzle
		}
	}

	// solve the blank cell
	if puzzle[i][j] == 0 {
		for testNum := 1; testNum <= 9; testNum++ {
			if valid(testNum, i, j, puzzle) {
				puzzle[i][j] = testNum
				isSolved, solved := solve(i + 1, j, puzzle)
				if isSolved {
					return true, solved
				}
			}
		}
		// Tried all 9 numbers. Reset it to 0 and backtrack.
		puzzle[i][j] = 0
		return false , puzzle
	}

	// Space was already filled.
	return solve(i + 1, j, puzzle)
}

// Returns true if x is a valid value to be placed at i, j; false otherwise
func valid(x, i, j int, puzzle Puzzle) bool {
	// check rows and columns
	for k := 0; k < 9; k++ {
		if x == puzzle[k][j] || x == puzzle[i][k] {
			return false
		}
	}

	// check boxes
	for row := 0; row < 3; row ++ {
		for col := 0; col < 3; col ++ {
			if x == puzzle[(i / 3) * 3 + row][(j / 3) * 3 + col] {
				return false
			}
		}
	}

	return true;
}
