package main

import (
	"fmt"
	"io"
    "bufio"
    "os"
    "strconv"
    "strings"
)

type Puzzle [9][9]int

func main() {
	myPuzzle := loadFromFile("sudokutest.txt")
	fmt.Println("Sudoku puzzle is:")
	print(myPuzzle)

	isSolved, solved := solve(0, 0, myPuzzle)
	if isSolved {
		fmt.Println("Puzzle solved. The solution is:")
		print(solved)
	} else {
		fmt.Println("The puzzle could not be solved.")
	}
}

// check that there is no error, panic with error otherwise
func checkError (err error) {
    if err != nil{
        panic(err)
    }
}

// given a sudoku puzzle in a txt file,
// read in the values and store in a 2d array (matrix)
// if file can't be found, exit with an error
func loadFromFile(filename string) (sudokuBoard Puzzle) {
    file, err := os.Open(filename)
    checkError(err)
    reader := bufio.NewReader(file)

    i := 0
    for i<9 && err!=io.EOF {
        j := 0
        digit:= ""
        for j<8 {
        	// read each digit in the row except for the last digit
        	digit, _ = reader.ReadString(' ')
            sudokuBoard[i][j], _ = strconv.Atoi(strings.TrimSpace(digit))
            j++
        }
        // read last digit which should have a newline character
        digit, _ = reader.ReadString('\n')
        sudokuBoard[i][j], _ = strconv.Atoi(strings.TrimSpace(digit))
        i++
    }

    return sudokuBoard
}

// Prints a sudoku Puzzle.
func print(puzzle Puzzle) {
	for iIndex, i := range puzzle {
		for jIndex, j := range i {
			fmt.Print(j, " ")
			if jIndex == 2 || jIndex == 5 {
				fmt.Print("| ")
			}
		}
		
		fmt.Println()
		if iIndex == 2 || iIndex == 5 {
			fmt.Println("------+-------+------")
		}
	}
	fmt.Println()
}

//solves the sudoku puzzle from i, j and returns true if it is solved,
//false if it cannot be solved.
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
