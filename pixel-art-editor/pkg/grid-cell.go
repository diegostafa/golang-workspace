package main

import (
	"image/color"
)

type gridCell struct {
	color color.RGBA
}

func newCell(color color.RGBA) *gridCell {
	newCell := new(gridCell)
	newCell.color = color
	return newCell
}
