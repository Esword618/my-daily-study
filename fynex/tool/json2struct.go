package tool

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/ChimeraCoder/gojson"
	"github.com/gogf/gf/v2/encoding/gjson"
)

func json2Struct() {
	w := fyne.CurrentApp().NewWindow("Json to struct")
	left := widget.NewMultiLineEntry()
	right := widget.NewMultiLineEntry()
	useUrl := widget.NewCheck("Use Url", nil)
	url := widget.NewEntry()
	url.SetText("https://api.github.com/repos/chimeracoder/gojson")
	structName := widget.NewEntry()
	structName.SetPlaceHolder("struct Name")
	pkgName := widget.NewEntry()
	pkgName.SetPlaceHolder("pkgName")
	tags := widget.NewEntry()
	tags.SetPlaceHolder("tags,split with ,")
	subStruct := widget.NewCheck("subStruct", nil)
	convertFloats := widget.NewCheck("ConvertFloats", nil)
	btn := widget.NewButton("Do", func() {
		if strings.TrimSpace(structName.Text) == "" {
			structName.SetText("NewStruct")
		}
		if strings.TrimSpace(pkgName.Text) == "" {
			pkgName.SetText("Main")
		}
		if useUrl.Checked {
			resp, err := http.Get(strings.TrimSpace(url.Text))
			if err != nil {
				right.SetText(err.Error())
				return
			}
			defer resp.Body.Close()

			bs, _ := io.ReadAll(resp.Body)
			js, _ := gjson.LoadContent(string(bs), true)
			jss, _ := js.ToJsonIndentString()
			left.SetText(jss)
		}
		if out, e := gojson.Generate(strings.NewReader(left.Text),
			gojson.ParseJson,
			structName.Text, pkgName.Text, strings.Split(tags.Text, ","),
			subStruct.Checked, convertFloats.Checked); e == nil {
			right.SetText(string(out))
		} else {
			right.SetText(fmt.Sprintf("JSON Parse Error: %v\n", e))
		}
	})
	split := container.NewHSplit(left, right)
	cc := container.NewBorder(container.NewVBox(url, btn), nil,
		container.NewVBox(useUrl, structName, pkgName,
			tags, subStruct, convertFloats),
		nil, split)
	w.SetContent(cc)
	w.Resize(fyne.NewSize(700,700))
	w.CenterOnScreen()
	w.Show()
}
