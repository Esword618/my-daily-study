package common

import (
	"image"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/tools/playground"
)

// Capture 对窗体的某个部分进行截图
// need to import "fyne.io/fyne/v2/tools/playground"
func Capture(obj fyne.CanvasObject, size fyne.Size, scale float32) image.Image {
	c := playground.NewSoftwareCanvas()
	c.SetContent(obj)
	c.Resize(size)
	c.SetScale(scale)
	return c.Capture()
}
