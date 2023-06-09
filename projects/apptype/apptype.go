package apptype

import (
	"image/color"

	"fyne.io/fyne/v2"
)

type BrushType = int

type PxCanvasConfig struct {
	draving_area     fyne.Size
	canvas_offset    fyne.Position
	px_rows, px_cols int
	px_size          int
}

type State struct {
	brush_color     color.Color
	brush_type      int
	swatch_selected int
	file_path       string
}

func (state *State) set_file_path(path string) {
	state.file_path = path
}
