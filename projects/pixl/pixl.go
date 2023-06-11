package main

import "test/apptype"
import "test/swatch"
import "test/ui"
import "test/pxcanvas"
import "image/color"
import "fyne.io/fyne/v2/app"
import "fyne.io/fyne/v2"

func main() {
	pixlApp := app.New()
	pixlWindow := pixlApp.NewWindow("pixl")

	state := apptype.State{
		BrushColor:     color.NRGBA{120, 255, 150, 100},
		SwatchSelected: 0,
	}

	pixlCanvasConfig := apptype.PxCanvasConfig{
		DrawingArea:  fyne.NewSize(300, 300),
		CanvasOffset: fyne.NewPos(0, 0),
		PxRows:       10,
		PxCols:       10,
		PxSize:       30,
	}

	pixlCanvas := pxcanvas.NewPxCanvas(&state, pixlCanvasConfig)

	appInit := ui.AppInit{
		PixlCanvas: pixlCanvas,
		PixlWindow: pixlWindow,
		State:      &state,
		Swatches:   make([]*swatch.Swatch, 0, 64),
	}
	ui.Setup(&appInit)
	appInit.PixlWindow.ShowAndRun()
}
