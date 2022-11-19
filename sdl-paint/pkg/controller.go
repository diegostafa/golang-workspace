package main

type controller struct {
	screen    *canvasScreen
	currSheet *sheet
	currPen   *pen
}

func (contr *controller) Exec(width int, height int) {
	contr.currSheet = CreateBlankSheet(size{width: int32(width), height: int32(height)}, color{hexARGB: 0xff0000})
	contr.screen = new(canvasScreen)
	contr.screen.controller = contr
	contr.screen.Init(contr.currSheet, contr.currPen)
}

func (contr *controller) SaveLine(line *line) {
	contr.currSheet.AddLine(line)

}
func (contr *controller) Undo() {
	contr.currSheet.Undo()
	contr.screen.RenderSheet(contr.currSheet)
}
