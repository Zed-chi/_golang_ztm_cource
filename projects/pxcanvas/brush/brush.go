package brush

import "test/apptype"
import "fyne.io/fyne/v2/driver/desktop"
import "fyne.io/fyne/v2"
import "image/color"
import "fyne.io/fyne/v2/canvas"

const (
	Pixel = iota
)

func TryPaintPixel(
	appState *apptype.State,
	canvas apptype.Brushable,
	ev *desktop.MouseEvent,
) bool {
	x, y := canvas.MouseToCanvas(ev)
	if x != nil && y != nil && ev.Button == desktop.MouseButtonPrimary {
		canvas.SetColor(appState.BrushColor, *x, *y)
		return true
	}
	return false
}

func TryBrush(
	appState *apptype.State,
	canvas apptype.Brushable,
	ev *desktop.MouseEvent,
) bool {
	switch {
	case appState.BrushType == Pixel:
		return TryPaintPixel(appState, canvas, ev)
	default:
		return false
	}
}

func Cursor(config apptype.PxCanvasConfig, brush apptype.BrushType, ev *desktop.MouseEvent, x, y int) []fyne.CanvasObject {
	if brush == Pixel {
		pxSize := float32(config.PxSize)
		xOrigin := (float32(x) * pxSize) + config.CanvasOffset.X
		yOrigin := (float32(y) * pxSize) + config.CanvasOffset.Y
		cursorColor := color.NRGBA{80, 80, 80, 255}
		left := canvas.NewLine(cursorColor)
		left.StrokeWidth = 3
		left.Position1 = fyne.NewPos(xOrigin, yOrigin)
		left.Position2 = fyne.NewPos(xOrigin, yOrigin+pxSize)

		top := canvas.NewLine(cursorColor)
		top.StrokeWidth = 3
		top.Position1 = fyne.NewPos(xOrigin, yOrigin)
		top.Position2 = fyne.NewPos(xOrigin+pxSize, yOrigin)

		right := canvas.NewLine(cursorColor)
		right.StrokeWidth = 3
		right.Position1 = fyne.NewPos(xOrigin+pxSize, yOrigin)
		right.Position2 = fyne.NewPos(xOrigin+pxSize, yOrigin+pxSize)

		bottom := canvas.NewLine(cursorColor)
		bottom.StrokeWidth = 3
		bottom.Position1 = fyne.NewPos(xOrigin, yOrigin+pxSize)
		bottom.Position2 = fyne.NewPos(xOrigin, yOrigin+pxSize)

		objects := []fyne.CanvasObject{left, top, right, bottom}
		return objects
	}
	return nil
}
