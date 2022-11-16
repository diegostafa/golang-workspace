package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type canvasScreen struct {
	window   *sdl.Window
	renderer *sdl.Renderer
	//windowSurface *sdl.Surface
	//sheetSurface  *sdl.Surface

	isRunning bool
	isDrawing bool
	currLine  *line

	controller *controller
}

func (screen *canvasScreen) ClearScreen() {
	screen.renderer.SetDrawColor(20, 0, 0, 255)
	screen.renderer.Clear()
	screen.renderer.Present()
	screen.renderer.SetDrawColor(20, 0, 0, 255)
}

func (screen *canvasScreen) Init(sheet *sheet, pen *pen) {
	screen.window = CreateWindow("Painter", sheet.size.width, sheet.size.height)

	screen.renderer = CreateRenderer(screen.window)
	screen.ClearScreen()
	screen.renderer.SetDrawColor(255, 0, 0, 255)

	screen.isRunning = true
	screen.currLine = new(line)
	screen.MainLoop()
}

func (screen *canvasScreen) MainLoop() {
	for screen.isRunning {
		screen.PollEvents()
	}
}
func (screen *canvasScreen) Close() {
	screen.window.Destroy()
	screen.renderer.Destroy()
	sdl.Quit()
}

func (screen *canvasScreen) PollEvents() {
	for pendingEvent := sdl.PollEvent(); pendingEvent != nil; pendingEvent = sdl.PollEvent() {
		switch pendingEvent.GetType() {
		case sdl.QUIT:
			screen.isRunning = false
		case sdl.WINDOWEVENT_RESIZED:
			screen.ClearScreen()
		case sdl.MOUSEBUTTONDOWN:
			clickEvent := pendingEvent.(*sdl.MouseButtonEvent)
			switch clickEvent.Button {
			case sdl.BUTTON_LEFT:
				screen.isDrawing = true
			case sdl.BUTTON_RIGHT:
				screen.controller.Undo()
			}
		case sdl.MOUSEBUTTONUP:
			clickEvent := pendingEvent.(*sdl.MouseButtonEvent)
			switch clickEvent.Button {
			case sdl.BUTTON_LEFT:
				screen.isDrawing = false
				screen.AddLine()

			}
		case sdl.MOUSEMOTION:
			if screen.isDrawing {
				motionEvent := pendingEvent.(*sdl.MouseMotionEvent)
				screen.AddPoint(motionEvent.X, motionEvent.Y)
				screen.RenderCurrentLine()
			}
		}
	}
}

func (screen *canvasScreen) AddLine() {
	if len(screen.currLine.points) > 0 {
		screen.controller.SaveLine(screen.currLine)
		screen.currLine = new(line)
	}
}

func (screen *canvasScreen) RenderCurrentLine() {
	if screen.isDrawing && len(screen.currLine.points) > 1 {
		screen.renderer.DrawLine(
			screen.currLine.points[len(screen.currLine.points)-2].x,
			screen.currLine.points[len(screen.currLine.points)-2].y,
			screen.currLine.points[len(screen.currLine.points)-1].x,
			screen.currLine.points[len(screen.currLine.points)-1].y)
		screen.renderer.Present()
	}
}

func (screen *canvasScreen) RenderSheet(sheet *sheet) {
	screen.ClearScreen()
	screen.renderer.SetDrawColor(255, 0, 0, 255)
	for i := 0; i < len(sheet.shapes); i++ {
		currLine := sheet.shapes[i]
		for j := 0; j < len(currLine.points)-1; j++ {
			screen.renderer.DrawLine(
				currLine.points[j].x,
				currLine.points[j].y,
				currLine.points[j+1].x,
				currLine.points[j+1].y)
		}
	}
	screen.renderer.Present()
}

func (screen *canvasScreen) AddPoint(x int32, y int32) {
	screen.currLine.points = append(screen.currLine.points, point{x: x, y: y})
}
