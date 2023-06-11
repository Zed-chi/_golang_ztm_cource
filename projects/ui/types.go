package ui

import "fyne.io/fyne/v2"
import "test/apptype"
import "test/swatch"
import "test/pxcanvas"

type AppInit struct {
	PixlCanvas *pxcanvas.PxCanvas
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
