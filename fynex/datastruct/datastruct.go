package datastruct

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func DataStruct() fyne.CanvasObject {

	btn1 := widget.NewButton("celltree", func() {
		celltreee()
	})

	grid := container.NewGridWithColumns(6, btn1)

	return grid

}
