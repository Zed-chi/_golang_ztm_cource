package ui

import "github.com/lusingander/colorpicker"
import "image/color"
import "fyne.io/fyne/v2"
import "fyne.io/fyne/v2/container"

func SetupColorPicker(app *AppInit) *fyne.Container {
	picker := colorpicker.New(200, colorpicker.StyleHue)
	picker.SetOnChanged(func(clr color.Color) {
		app.State.BrushColor = clr
		app.Swatches[app.State.SwatchSelected].SetColor(clr)
	})
	return container.NewVBox(picker)
}
