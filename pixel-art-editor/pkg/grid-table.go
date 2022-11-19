package main

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

type gridTable struct {
	rows        int
	cols        int
	borderSize  int
	cellColor   color.RGBA
	borderColor color.RGBA
	cells       [][]*gridCell
	surface     *pixelgl.Window
}

func newGrid(rows int, cols int, surface *pixelgl.Window, cellColor color.RGBA, borderColor color.RGBA) *gridTable {
	grid := new(gridTable)
	grid.rows = rows
	grid.cols = cols
	grid.surface = surface
	grid.borderSize = 1
	grid.cellColor = cellColor
	grid.borderColor = borderColor
	grid.cells = make([][]*gridCell, cols)

	for i := range grid.cells {
		grid.cells[i] = make([]*gridCell, rows)
	}

	for j := 0; j < rows; j++ {
		for i := 0; i < cols; i++ {
			newCell := newCell(cellColor)
			grid.cells[i][j] = newCell
		}
	}

	return grid
}

func (g *gridTable) renderCells(canvas *imdraw.IMDraw) {
	for i := 0; i < g.rows; i++ {
		for j := 0; j < g.cols; j++ {
			topLeft, bottomRight := g.cellToScreenCoord(i, j)
			canvas.Color = g.cells[i][j].color
			canvas.Push(topLeft, bottomRight)
			canvas.Rectangle(0)
			canvas.EndShape = imdraw.SharpEndShape
		}
	}
}

func (g *gridTable) renderBorders(canvas *imdraw.IMDraw) {
	width, height := g.cellSize()
	for i := 0; i < g.rows; i++ {
		for j := 0; j < g.cols; j++ {
			canvas.Color = g.borderColor

			hLineStart := pixel.V(0, height*float64(i))
			hLineEnd := pixel.V(g.surface.Bounds().W(), height*float64(i))
			canvas.Push(hLineStart)
			canvas.Push(hLineEnd)
			canvas.Line(float64(g.borderSize))
			canvas.EndShape = imdraw.SharpEndShape

			vLineStart := pixel.V(width*float64(j), 0)
			vLineEnd := pixel.V(width*float64(j), g.surface.Bounds().H())
			canvas.Push(vLineStart)
			canvas.Push(vLineEnd)
			canvas.Line(float64(g.borderSize))
			canvas.EndShape = imdraw.SharpEndShape

		}
	}
}

func (g *gridTable) highlightCell(canvas *imdraw.IMDraw, mousePos pixel.Vec, color color.RGBA) {
	row, col := g.screenToCellCoord(mousePos.X, mousePos.Y)
	topLeft, bottomRight := g.cellToScreenCoord(row, col)
	canvas.Color = color
	canvas.Push(topLeft, bottomRight)
	canvas.Rectangle(0)
	canvas.EndShape = imdraw.SharpEndShape
}

func (g *gridTable) cellAt(mousePos pixel.Vec) *gridCell {
	row, col := g.screenToCellCoord(mousePos.X, mousePos.Y)
	return g.cells[row][col]
}

func (g *gridTable) clear() {
	for i := 0; i < g.rows; i++ {
		for j := 0; j < g.cols; j++ {
			g.cells[i][j].color = g.cellColor
		}
	}
}

func (g *gridTable) cellSize() (float64, float64) {
	cellWidth := g.surface.Bounds().W() / float64(g.rows)
	cellheight := g.surface.Bounds().H() / float64(g.cols)

	return cellWidth, cellheight
}

func (g *gridTable) screenToCellCoord(posX float64, posY float64) (int, int) {
	cellWidth, cellHeight := g.cellSize()
	row := posX / cellHeight
	col := posY / cellWidth
	return int(row) % g.rows, int(col) % g.cols
}

func (g *gridTable) cellToScreenCoord(row int, col int) (pixel.Vec, pixel.Vec) {
	cellWidth, cellHeight := g.cellSize()
	topLeft := pixel.V(float64(row)*cellWidth, float64(col)*cellHeight)
	bottomRight := pixel.V(topLeft.X+cellWidth, topLeft.Y+cellHeight)
	return topLeft, bottomRight
}
