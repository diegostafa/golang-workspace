package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"

	"golang.org/x/image/colornames"
)

func main() {
	pixelgl.Run(run)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "pixel art painter",
		Bounds: pixel.R(0, 0, 900, 900),
		//VSync:     true,
		//Resizable: true, // BROKE
	}

	win, _ := pixelgl.NewWindow(cfg)
	canvas := imdraw.New(nil)
	grid := newGrid(16, 16, win, colornames.Black, colornames.Aliceblue)

	for !win.Closed() {
		win.Clear(colornames.Firebrick)
		canvas.Clear()
		grid.renderCells(canvas)

		if win.MouseInsideWindow() {
			grid.highlightCell(canvas, win.MousePosition(), colornames.Lightgreen)
			if win.Pressed(pixelgl.MouseButtonLeft) {
				grid.cellAt(win.MousePosition()).color = colornames.Orange
			}
			if win.Pressed(pixelgl.MouseButtonRight) {
				grid.cellAt(win.MousePosition()).color = colornames.Black
			}
		}

		if win.JustPressed(pixelgl.KeyC) {
			grid.clear()
		} else if !win.Pressed(pixelgl.KeyEnter) {
			grid.renderBorders(canvas)
		}

		canvas.Draw(win)
		win.Update()
	}
}
