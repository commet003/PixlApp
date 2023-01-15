package apptype

import (
	"image/color"
	"fyne.io/fyne/v2"
)

type BrushType = int

type PxCanvasconfig struct {
	DrawingArea    fyne.Size
	CanvasOffset   fyne.Position
	PxRows, PxCols int
	PxSize         int
}

type State struct {
	BrushColor      color.Color
	BrushType       int
	SwatchSeleceted int
	FilePath        string
}

func (state *State) SetFilePath(path string) {
	state.FilePath = path
}
