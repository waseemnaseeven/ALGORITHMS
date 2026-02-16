package main

import (
	"fmt"
)

/*
	According to Wikipedia's article: "The Game of Life, also known simply as Life, is a cellular automaton devised by the British mathematician John Horton Conway in 1970."

	The board is made up of an m x n grid of cells, where each cell has an initial state: live (represented by a 1) or dead (represented by a 0). Each cell interacts with its eight neighbors (horizontal, vertical, diagonal) using the following four rules (taken from the above Wikipedia article):

	- Any live cell with fewer than two live neighbors dies as if caused by under-population.
	- Any live cell with two or three live neighbors lives on to the next generation.
	- Any live cell with more than three live neighbors dies, as if by over-population.
	- Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.

	The next state of the board is determined by applying the above rules simultaneously to every cell in the current state of the m x n grid board. In this process, births and deaths occur simultaneously.

	Given the current state of the board, update the board to reflect its next state.

	Note that you do not need to return anything.
*/

const (
    ALIVE   = 1
    DEAD    = 0
)

func countNeighbors(board [][]int, row, col int) int {
    rows := len(board)
	cols := len(board[0])
	
    var count int
    for i := row-1; i <= row+1; i++ {
        for j := col-1; j <= col+1; j++ {
            if i < 0 || i >=rows || j < 0 || j >= cols {
                continue
            }
            if i == row && j == col {
                continue
            }

            count += board[i][j]
        }
    }
    return count
}

func gameOfLife(board [][]int) {
    rows := len(board)
    if rows == 0 {
        return 
    }
    cols := len(board[0])

    newBoard := make([][]int, rows)
    for i := 0; i < rows; i++ {
        newBoard[i] = make([]int, cols)
        copy(newBoard[i], board[i])
    }

    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            newNeighbors := countNeighbors(board, i, j)
            if board[i][j] == ALIVE {
                if newNeighbors < 2 || newNeighbors > 3 {
                    newBoard[i][j] = DEAD
                }
            } else {
                if newNeighbors == 3 {
                    newBoard[i][j] = ALIVE
                }
            }
        }
    }

    for i := 0; i < rows; i++ {
        copy(board[i], newBoard[i])
    }
}

/*

	Write a program to solve a Sudoku puzzle by filling the empty cells.

	A sudoku solution must satisfy all of the following rules:

	Each of the digits 1-9 must occur exactly once in each row.
	Each of the digits 1-9 must occur exactly once in each column.
	Each of the digits 1-9 must occur exactly once in each of the 9 3x3 sub-boxes of the grid.
	The '.' character indicates empty cells.

	Input: board = [["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]

	Output: [["5","3","4","6","7","8","9","1","2"],["6","7","2","1","9","5","3","4","8"],["1","9","8","3","4","2","5","6","7"],["8","5","9","7","6","1","4","2","3"],["4","2","6","8","5","3","7","9","1"],["7","1","3","9","2","4","8","5","6"],["9","6","1","5","3","7","2","8","4"],["2","8","7","4","1","9","6","3","5"],["3","4","5","2","8","6","1","7","9"]]

*/

func isValidMove(board [][]byte, row, col int, d byte) bool {
    // Verify one lign and one column 9x9
    for i := 0; i < 9; i++ {
        if board[row][i] == d || board[i][col] == d {
            return false
        }
    }

    startRow := 3 * (row / 3)
    startCol := 3 * (col / 3)

    for i := startRow; i < startRow+3; i++ {
        for j := startCol; j < startCol+3; j++ {
            if board[i][j] == d {
                return false
            }
        }
    }
    return true
}

func backtracking(board [][]byte) bool {
    for row := 0; row < 9; row++ {
        for col := 0; col < 9; col++ {
            if board[row][col] == '.' {
                for d := byte('1'); d <= '9'; d++ {
                    isValidMove(board, row, col, d)
                }
                return false
            }
        }
    }
    return true
}

func solveSudoku(board [][]byte) {
    backtracking(board)
}


func main() {

	// GameOfLife
	board := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}

	fmt.Println(board)
	gameOfLife(board)
	fmt.Println(board)

	// Sudoku Solver
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	fmt.Println()
	solveSudoku(board)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Println("%c ", board[i][j])
		}
		fmt.Println()
	}
}