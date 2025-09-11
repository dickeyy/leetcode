package main

/*
 *   Stats:
 *   Runtime: 0ms (beats 100.00% of Go solutions)
 *   Memory: 5.22mb (beats 74.75% of Go solutions)
 */

func solveNQueens(n int) [][]string {
	result := [][]string{}
	board := make([][]byte, n)
	for i := range board {
		board[i] = make([]byte, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}

	backtrack(&result, board, 0, n)
	return result
}

func backtrack(result *[][]string, board [][]byte, row, n int) {
	// base case: if we have placed all queens
	if row == n {
		sol := make([]string, n)
		for i := range board {
			sol[i] = string(board[i])
		}
		*result = append(*result, sol)
		return
	}

	// try placing a queen in each col of curr row
	for col := 0; col < n; col++ {
		if isSafe(board, row, col, n) {
			// place queen
			board[row][col] = 'Q'

			// recursively place queens
			backtrack(result, board, row+1, n)

			// backtrack remove queen
			board[row][col] = '.'
		}
	}
}

func isSafe(board [][]byte, row, col, n int) bool {
	// check column
	for i := 0; i < row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}

	// check upper-left diagonal
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	// check upper-right diagonal
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	return true
}
