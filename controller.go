package main

import (
	"fmt"
	"math"
	"os"
	"time"

	"github.com/faiface/pixel"

	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var menu = []string{"What would you like to do?", "1 - Toggle one cell", "2 - Compute next generation", "3 - Compute several generations", "4 - change game engine", "0 - Exit"}
var engineList = []string{"1 - Conway", "2 - Anneal"}

func initialize() {
	// ge := gameEngine(conway{})

	w := createWindow()
	mainGameLoop(w)
	os.Exit(0)

}

func mainGameLoop(w *pixelgl.Window) {
	wh := imdraw.New(nil)
	bl := imdraw.New(nil)
	b := newBoard()
	ge := conway{}
	run := false
	c := make(chan int)
	for !w.Closed() {
		w.Clear(colornames.Gray)
		bl.Clear()
		wh.Clear()
		wh.Color = colornames.White
		bl.Color = colornames.Black

		if w.JustPressed(pixelgl.MouseButtonLeft) {
			pos := w.MousePosition()
			b.toggleCell(int(pos.Y/cellSize), int(math.Floor(pos.X)/cellSize))
		}
		if w.JustPressed(pixelgl.KeySpace) {
			run = !run
			fmt.Println(run)
		}

		if run {
			b = nextGeneration(ge, b)
			go wait(c)
			<-c
		}
		fillImd(wh, bl, b)

		wh.Draw(w)
		bl.Draw(w)
		w.Update()

	}
}

func wait(c chan int) {
	time.Sleep(100 * time.Millisecond)
	c <- 0
}

func fillImd(wh, bl *imdraw.IMDraw, b board) {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {

			if b[i][j] == alive {
				bl.Push(pixel.V(float64(cellSize*j), float64((cellSize*i)+cellSize)))
				bl.Push(pixel.V(float64((cellSize*j)+cellSize), float64(cellSize*i)))
				bl.Rectangle(0)
			} else {
				wh.Push(pixel.V(float64(cellSize*j), float64((cellSize*i)+cellSize)))
				wh.Push(pixel.V(float64((cellSize*j)+cellSize), float64(cellSize*i)))
				wh.Rectangle(1)
			}
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

func printList(l []string) {
	for _, line := range l {
		fmt.Println(line)
	}
	fmt.Println()
}
