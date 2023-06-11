package swatch

import "fyne.io/fyne/v2/widget"
import "fyne.io/fyne/v2"
import "fyne.io/fyne/v2/canvas"

import "test/apptype"
import "image/color"

type Swatch struct {
	widget.BaseWidget
	Selected     bool
	Color        color.Color
	SwatchIndex  int
	clickHandler func(s *Swatch)
}

func (swatch *Swatch) SetColor(c color.Color) {
	swatch.Color = c
	swatch.Refresh()
}

func NewSwatch(
	state *apptype.State, color color.Color,
	swatch_id int, click_handler func(s *Swatch),
) *Swatch {
	s := &Swatch{
		Selected:     false,
		Color:        color,
		SwatchIndex:  swatch_id,
		clickHandler: click_handler,
	}
	s.ExtendBaseWidget(s)
	return s
}

func (swatch *Swatch) CreateRenderer() fyne.WidgetRenderer {
	square := canvas.NewRectangle(swatch.Color)
	objects := []fyne.CanvasObject{square}
	return &SwatchRenderer{
		square:  *square,
		objects: objects,
		parent:  swatch,
	}
}
