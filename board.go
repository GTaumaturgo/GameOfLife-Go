package main

import (
	"fmt"
)

type board [][]cell

const boardSize = 50
const cellSize = 750 / boardSize

func newBoard() board {
	b := board(make([][]cell, boardSize))

	for index := 0; index < boardSize; index++ {
		b[index] = make([]cell, boardSize)
	}

	return b
}

func (b board) print() {
	for i, line := range b {
		s := make([]byte, boardSize)

		for j := range line {
			if b[i][j] == false {
				s[j] = '.'
			} else {
				s[j] = 'o'
			}
		}
		fmt.Println(string(s))
	}
	fmt.Println()
}

func (b board) makeCellAlive(l int, c int) {
	b[l][c] = alive
}
func (b board) makeCellDead(l int, c int) {
	b[l][c] = dead
}
func (b board) toggleCell(l int, c int) {
	b[l][c] = !(b[l][c])
}

func (b board) numberOfAliveNeighbours(l int, c int) int {
	result := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if b[(l+i+boardSize)%boardSize][(c+j+boardSize)%boardSize] {
				result++
			}
		}
	}
	return result
}
