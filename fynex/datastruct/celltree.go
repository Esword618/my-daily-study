package datastruct

import (
	"bytes"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/tidwall/celltree"
)

func celltreee() {
	input := widget.NewEntry()
	input.SetText("10")
	input.PlaceHolder = "input a number to insert or delete"
	output := widget.NewEntry()
	lblCount := widget.NewLabel("")
	var tr celltree.Tree
	tr.Insert(10, nil)
	tr.Insert(5, nil)
	tr.Insert(31, nil)
	tr.Insert(16, nil)
	tr.Insert(9, nil)

	buf := bytes.NewBuffer(nil)
	view := func() {
		buf.Reset()
		tr.Scan(func(cell uint64, data interface{}) bool {
			buf.WriteString(gconv.String(cell) + "\n")
			return true
		})
		output.SetText(buf.String())
		lblCount.SetText("Count of the celltree: "+gconv.String(tr.Count()))
	}
	btnInsert := widget.NewButton("Insert", func() {
		tr.Insert(gconv.Uint64(input.Text), nil)
		view()
	})
	btnView := widget.NewButton("View", func() {
		view()
	})
	btnDelete := widget.NewButton("Delete", func() {
		tr.Delete(gconv.Uint64(input.Text), nil)
		view()
	})
	view()
	cc := container.NewBorder(container.NewVBox(input,
		container.NewHBox(btnInsert, btnDelete, btnView, lblCount)),
		nil, nil, nil, output)
	w := fyne.CurrentApp().NewWindow("CellTree")

	w.SetContent(cc)
	w.CenterOnScreen()
	w.Resize(fyne.NewSize(700, 700))
	w.Show()
}
