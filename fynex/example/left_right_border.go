package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	xlyt "gitee.com/y2h/fynex/layout"
)

func lrborder() fyne.CanvasObject {

	leftborder := xlyt.NewLeftBorder(widget.NewLabel("Left"), widget.NewEntry())
	rightborder := xlyt.NewRightBorder(widget.NewLabel("Right"), widget.NewEntry())
	return container.NewGridWithRows(2, leftborder, rightborder)
}
