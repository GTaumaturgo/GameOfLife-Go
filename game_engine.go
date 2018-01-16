package main

type gameEngine interface {
	shouldBeAlive(b board, l int, c int) bool
}

func nextGeneration(ge gameEngine, b board) board {
	result := newBoard()

	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if ge.shouldBeAlive(b, i, j) {
				result[i][j] = alive
			}
		}
	}
	return result
}
