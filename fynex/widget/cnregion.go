package widget

import (
	"bytes"
	"embed"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	// "github.com/gogf/gf/v2/frame/g"

	// "github.com/gogf/gf/v2/util/gutil"
	"github.com/issue9/cnregion"
)

//go:embed regions.db
var regions embed.FS

func Cnregion() fyne.CanvasObject {
	v, err := cnregion.LoadFS(regions, "regions.db", "-", 2021)
	// v, err := cnregion.LoadFile("./regions.db", "-", 2020)
	if err != nil {
		// msg.SetText("load regions.db error " + err.Error())
		return nil
	}
	lblMsg := widget.NewEntry()
	lblMsg.PlaceHolder = "input search region(温州) here then click button Search"
	buf := bytes.NewBuffer(nil)
	btnSearch := widget.NewButton("Search", func() {
		buf.Reset()
		if lblMsg.Text==""{
			return
		}
		resultEntry := widget.NewMultiLineEntry()
		regionsSearched := v.Search(&cnregion.SearchOptions{Text: lblMsg.Text})
		for _, v := range regionsSearched {
			buf.WriteString(v.FullName() + "\n")
		}
		// g.DumpTo(buf, regionsSearched,gutil.DumpOption{})
		resultEntry.SetText(buf.String())
		win := fyne.CurrentApp().NewWindow("Searched Result")
		win.Resize(fyne.NewSize(600, 600))
		win.SetContent(resultEntry)
		win.CenterOnScreen()
		win.Show()
	})
	optsg := v.Provinces()
	selLast := NewSelectG(nil).Col(3).Change(func(sg *SelectG) {
		if sg.Selected == nil {
			return
		}
		lblMsg.SetText(sg.Selected.FullName())
	}).Do(func(sg *SelectG) {
		sg.PlaceHolder = "乡"
	})
	selStreet := NewSelectG(nil).Col(3).Change(func(sg *SelectG) {
		if sg.SelectedIndex() == -1 {
			return
		}
		selLast.SetOptions(sg.Options[sg.SelectedIndex()].Items())
		selLast.SetSelectedIndex(0)
		lblMsg.SetText(sg.Selected.FullName())
	}).Do(func(sg *SelectG) {
		sg.PlaceHolder = "镇"
	})
	selTown := NewSelectG(nil).Col(3).Change(func(sg *SelectG) {
		lblMsg.SetText(sg.Selected.FullName())
		if len(sg.Options[sg.SelectedIndex()].Items()) == 0 {
			selStreet.ClearSelected()
			selStreet.Options = nil
			selStreet.PlaceHolder = "no data"
			selLast.ClearSelected()
			selLast.Options = nil
			selLast.PlaceHolder = "no data"
			return
		}
		selStreet.SetOptions(sg.Options[sg.SelectedIndex()].Items())
		selStreet.SetSelectedIndex(0)
	}).Do(func(sg *SelectG) {
		sg.PlaceHolder = "县"
	})
	selCity := NewSelectG(nil).Col(3).Change(func(sg *SelectG) {
		selTown.SetOptions(sg.Options[sg.SelectedIndex()].Items())
		selTown.SetSelectedIndex(0)
		lblMsg.SetText(sg.Selected.FullName())
	}).Do(func(sg *SelectG) {
		sg.PlaceHolder = "市"
	})
	selProvince := NewSelectG(optsg).Col(2).Change(func(sg *SelectG) {
		selCity.SetOptions(optsg[sg.SelectedIndex()].Items())
		selCity.SetSelectedIndex(0)
		lblMsg.SetText(sg.Selected.FullName())
	}).Do(func(sg *SelectG) {
		sg.PlaceHolder = "省"
	})
	return container.NewVBox(container.NewBorder(nil, nil, nil,
		btnSearch, lblMsg),
		container.NewGridWithColumns(5, selProvince,
			selCity, selTown, selStreet, selLast))
}
