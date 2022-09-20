package component

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func CyyList(data []string) *widget.List {
	return widget.NewList(func() int {
		return len(data)

	}, func() fyne.CanvasObject {
		return widget.NewLabel("i am a list")

	}, func(id int, obj fyne.CanvasObject) {
		c := obj.(*widget.Label)
		c.SetText(data[id])
	})
}
