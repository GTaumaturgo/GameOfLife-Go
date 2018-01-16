package main

import "testing"

func TestNewBoard(t *testing.T) {
	b := newBoard()
	if len(b) != len(b[0]) || len(b) != boardSize {
		t.Errorf("size of the board different than expected after creation")
	}
}

func TestToggleCell(t *testing.T) {
	b := newBoard()

	b.toggleCell(0, 0)
	if b[0][0] != alive {
		t.Errorf("toggling a dead cell is not making it alive")
	}
	b.toggleCell(0, 0)
	if b[0][0] != dead {
		t.Errorf("toggling an alive cell is not making it dead")
	}
}

func TestMakeCellAlive(t *testing.T) {
	b := newBoard()
	b.makeCellAlive(0, 0)
	if b[0][0] != alive {
		t.Errorf("make cell alive not working properly")
	}
}
func TestMakeCellDead(t *testing.T) {
	b := newBoard()
	b.makeCellAlive(0, 0)
	b.makeCellDead(0, 0)
	if b[0][0] != dead {
		t.Errorf("make cell dead not working properly")
	}
}

func TestNumberOfAliveNeighbours(t *testing.T) {
	b := newBoard()

	b.makeCellAlive(boardSize-1, 0)
	b.makeCellAlive(0, 0)
	b.makeCellAlive(0, boardSize-1)
	if b.numberOfAliveNeighbours(boardSize-1, boardSize-1) != 3 {
		t.Errorf("number of neighbours not being counted properly")
	}

}
