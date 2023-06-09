package pxcanvas

import "fyne.io/fyne/v2"
import "test/pxcanvas/brush"
import "fyne.io/fyne/v2/driver/desktop"

func (pxCanvas *PxCanvas) Scrolled(ev *fyne.ScrollEvent) {
	pxCanvas.scale(int(ev.Scrolled.DY))
	pxCanvas.Refresh()
}

func (pxCanvas *PxCanvas) MouseMoved(ev *desktop.MouseEvent) {
	if x, y := pxCanvas.MouseToCanvas(ev); x != nil && y != nil {
		brush.TryBrush(pxCanvas.appState, pxCanvas, ev)
		cursor := brush.Cursor(
			pxCanvas.PxCanvasConfig, pxCanvas.appState.BrushType, ev, *x, *y,
		)
		pxCanvas.renderer.SetCursor(cursor)
	} else {
		pxCanvas.renderer.SetCursor(make([]fyne.CanvasObject, 0))
	}
	pxCanvas.TryPan(pxCanvas.mouseState.previousCoord, ev)
	pxCanvas.Refresh()
	pxCanvas.mouseState.previousCoord = &ev.PointEvent
}

func (pxCanvas *PxCanvas) MouseIn(ev *desktop.MouseEvent) {

}

func (pxCanvas *PxCanvas) MouseOut(ev *desktop.MouseEvent) {

}

func (pxCanvas *PxCanvas) MouseDown(ev *desktop.MouseEvent) {
	brush.TryBrush(pxCanvas.appState, pxCanvas, ev)
}

func (pxCanvas *PxCanvas) MouseUp(ev *desktop.MouseEvent) {

}
