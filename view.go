package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func createWindow() *(pixelgl.Window) {
	cfg := pixelgl.WindowConfig{
		Title:  "Game Of Life",
		Bounds: pixel.R(0, 0, 750, 750),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	return win
}
