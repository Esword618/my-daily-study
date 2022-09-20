package main

import (
	"fmt"
	"image/color"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/ncruces/zenity"
)

func Zenity() fyne.CanvasObject {
	lblMsg := widget.NewMultiLineEntry()
	ctxt := canvas.NewText("i am the color", color.Black)
	btn1 := widget.NewButton("SelectFile", func() {
		spath, _ := zenity.SelectFile(
			zenity.FileFilters{
				{"Image files", []string{"*.png", "*.gif", "*.ico", "*.jpg", "*.webp"}},
			},
			zenity.Title("Please choose an image"),

			// zenity.Directory(),
		)
		lblMsg.SetText(spath)
	})
	btn2 := widget.NewButton("SelectFileMutiple", func() {
		spath, _ := zenity.SelectFileMutiple(
			zenity.FileFilters{
				{"Image files", []string{"*.png", "*.gif", "*.ico", "*.jpg", "*.webp"}},
			},
			zenity.Title("SelectFileMutiple"),
			// zenity.Directory(),
		)
		lblMsg.SetText(strings.Join(spath, "\n"))
	})
	btn3 := widget.NewButton("SelectFileSave", func() {
		spath, _ := zenity.SelectFileSave(
			zenity.FileFilters{
				{"Image files", []string{"*.png", "*.gif", "*.ico", "*.jpg", "*.webp"}},
			},
			zenity.Title("SelectFileSave"),
			// zenity.Directory(),
		)
		lblMsg.SetText(spath)
	})

	btn4 := widget.NewButton("SelectColor", func() {
		c, _ := zenity.SelectColor(
			zenity.Title("SelectColor"),
		)
		ctxt.Color = c
		ctxt.Refresh()

	})
	btn5 := widget.NewButton("Calendar", func() {
		t, _ := zenity.Calendar("Selcet date",
			zenity.Title("Select date"),
		)
		lblMsg.SetText(t.String())
	})
	btn6 := widget.NewButton("Entry", func() {
		t, _ := zenity.Entry("entry",
			zenity.Title("Select date"),
		)
		lblMsg.SetText(t)
	})
	btn7 := widget.NewButton("List", func() {
		t, _ := zenity.List("list", []string{"list1", "list2", "list3", "list4"},
			zenity.Title("Select date"),
		)
		lblMsg.SetText(t)
	})
	btn8 := widget.NewButton("ListItems", func() {
		t, _ := zenity.ListItems("list", "list1", "list2", "list3", "list4")
		lblMsg.SetText(t)
	})
	btn9 := widget.NewButton("ListMultipleItems", func() {
		t, _ := zenity.ListMultipleItems("list", "list1", "list2", "list3", "list4")
		lblMsg.SetText(strings.Join(t, "\n"))
	})
	btn10 := widget.NewButton("ListMultiple", func() {
		t, _ := zenity.ListMultiple("list", []string{"list1", "list2", "list3", "list4"})
		lblMsg.SetText(strings.Join(t, "\n"))
	})
	btn11 := widget.NewButton("Progress", func() {
		t, _ := zenity.Progress()
		lblMsg.SetText(fmt.Sprint(t.MaxValue()))
	})
	btn12 := widget.NewButton("Info", func() {
		zenity.Info("Info")
	})

	btn13 := widget.NewButton("Notify", func() {
		zenity.Notify("i am a notify")
	})
	btn14 := widget.NewButton("Password", func() {
		nm, pd, _ := zenity.Password()
		lblMsg.SetText(nm + "\n" + pd)
	})
	btn15 := widget.NewButton("Question", func() {
		zenity.Question("i am a Question")
	})

	btn16 := widget.NewButton("Warning", func() {
		zenity.Warning("i am a Warning")
	})
	btn17 := widget.NewButton("Error", func() {
		zenity.Error("i am a Error")
	})
	return container.NewVBox(lblMsg, ctxt, btn1, btn2, btn3, btn4, btn5, btn6, btn7,
		btn8, btn9, btn10, btn11, btn12, btn13,btn14,btn15,btn16,btn17)

}
