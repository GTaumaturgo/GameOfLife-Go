package main

type anneal struct{}

func (anneal) shouldBeAlive(b board, l int, c int) bool {

	n := b.numberOfAliveNeighbours(l, c)
	if b[l][c] == alive {
		return (n >= 5 && n <= 8) || n == 3
	}
	return n == 4 || n == 6 || n == 7 || n == 8
}
