package main

type conway struct{}

func (conway) shouldBeAlive(b board, l int, c int) bool {

	n := b.numberOfAliveNeighbours(l, c)
	if b[l][c] == alive {
		return n == 2 || n == 3
	}
	return n == 3
}
