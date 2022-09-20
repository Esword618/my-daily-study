package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	xwdg "gitee.com/y2h/fynex/widget"
)

func selectPage() fyne.CanvasObject {
	lblMsg := widget.NewLabel("Msg")
	opts := numOpts(30)
	sel1 := xwdg.NewSelect(opts,  func(s string) {
		lblMsg.SetText(s)
	}).Col(2).Row(2).CurPage(5)
	sel2 := xwdg.NewSelect(opts, func(s string) {
		lblMsg.SetText(s)
	}).Col(3).Row(3)
	return container.NewVBox(lblMsg, sel1, sel2)
}

func numOpts(i int) (s []string) {
	for j:=0; j < i; j++ {
		s = append(s, strconv.Itoa(j))
	}
	return
}
