package main

type pen struct {
}

type size struct {
	width, height int32
}

type point struct {
	x, y int32
}

type line struct {
	points []point
}

type color struct {
	hexARGB rune
}

type sheet struct {
	size   size
	color  color
	shapes []*line
}

func CreateBlankSheet(size size, color color) *sheet {
	s := new(sheet)
	s.size = size
	s.color = color
	return s
}

func (s *sheet) Undo() {
	if len(s.shapes) > 0 {
		s.shapes = s.shapes[:len(s.shapes)-1]
	}
}

func (s *sheet) AddLine(line *line) {
	s.shapes = append(s.shapes, line)
}
