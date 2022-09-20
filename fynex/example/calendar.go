package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	xwdg "gitee.com/y2h/fynex/widget"
	_ "gitee.com/y2h/goutil"
)

func calendar() fyne.CanvasObject {
	cc := container.NewVBox(xwdg.NewCalendar())
	return cc
}
