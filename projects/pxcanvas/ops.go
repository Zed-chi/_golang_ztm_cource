package pxcanvas

import "fyne.io/fyne/v2"

func (pxcanvas *PxCanvas) scale(direction int) {
	switch {
	case direction > 0:
		pxcanvas.PxSize += 1
	case direction < 0:
		if pxcanvas.PxSize > 2 {
			pxcanvas.PxSize -= 1
		}
	default:
		pxcanvas.PxSize = 10
	}
}

func (pxcanvas *PxCanvas) Pan(prevCoord, curCoord fyne.PointEvent) {
	xDelta := curCoord.Position.X - prevCoord.Position.X
	yDelta := curCoord.Position.Y - prevCoord.Position.Y

	pxcanvas.CanvasOffset.X += xDelta
	pxcanvas.CanvasOffset.Y += yDelta
	pxcanvas.Refresh()
}
