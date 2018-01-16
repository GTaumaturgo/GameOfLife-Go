package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

var menu = []string{"What would you like to do?", "1 - Toggle one cell", "2 - Compute next generation", "3 - Compute several generations", "4 - change game engine", "0 - Exit"}
var engineList = []string{"1 - Conway", "2 - Anneal"}

func main() {
	ge := gameEngine(conway{})

	b := newBoard()

	for {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		b.print()
		printList(menu)
		var inp int
		fmt.Scanf("%d", &inp)
		switch inp {
		case 1:
			toggleCell(b)
		case 2:
			b = nextGeneration(ge, b)
		case 3:
			b = multipleGenerations(ge, b)
		case 4:
			ge = changeGameEngine()
		case 0:
			os.Exit(0)
		}

	}
}

func changeGameEngine() gameEngine {
	fmt.Println("Select one of the engines")
	printList(engineList)
	var e int
	fmt.Scanf("%d", &e)
	switch e {
	case 1:
		return gameEngine(conway{})
	case 2:
		return gameEngine(anneal{})
	default:

		return gameEngine(conway{})
	}

}

func multipleGenerations(ge gameEngine, b board) board {
	var n int
	fmt.Println("How many generations should be computed?")
	fmt.Scanf("%d", &n)
	for index := 0; index < n; index++ {
		b = nextGeneration(ge, b)
		b.print()
		printList(menu)
		time.Sleep(250 * time.Millisecond)
	}
	return b
}

func toggleCell(b board) {
	var l int
	var c int
	fmt.Println("Digit the coordinates of the cell to toggle (", 0, "-", boardSize-1, ")")
	fmt.Scanf("%d %d", &l, &c)
	b.toggleCell(l, c)
}

func printList(l []string) {
	for _, line := range l {
		fmt.Println(line)
	}
	fmt.Println()
}
