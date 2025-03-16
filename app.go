package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Display(i [][]bool) {
	for _, j := range i {
		for _, k := range j {
			if k {
				fmt.Print("██")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}

func CountNeighbors(board [][]bool, x, y int) int {
	directions := []struct{ x, y int }{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	count := 0
	for _, dir := range directions {
		nx, ny := x+dir.x, y+dir.y
		if nx >= 0 && ny >= 0 && ny < len(board) && nx < len(board[0]) && board[ny][nx] {
			count++
		}
	}
	return count
}

func Step(board [][]bool) [][]bool {
	next := make([][]bool, len(board))

	for y := range board {
		next[y] = make([]bool, len(board[y]))
		for x := range board[y] {
			n := CountNeighbors(board, x, y)
			if board[y][x] {
				next[y][x] = n == 2 || n == 3
			} else {
				next[y][x] = n == 3
			}
		}
	}
	return next
}

func main() {
	rand.Seed(time.Now().UnixNano())

	board := make([][]bool, 40)
	for i := range board {
		board[i] = make([]bool, 100)
	}

	for i := range board {
		for j := range board[i] {
			board[i][j] = rand.Float64() > 0.5
		}
	}

	for {
		Display(board)
		board = Step(board)
		time.Sleep(400 * time.Millisecond)
		fmt.Print("\033[H\033[2J")
	}
}
