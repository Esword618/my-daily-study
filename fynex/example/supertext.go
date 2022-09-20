package main

import (
	"bytes"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	wdg "gitee.com/y2h/fynex/widget"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gutil"
	"golang.org/x/image/colornames"
)

func supertext() fyne.CanvasObject {
	lblMsg := widget.NewLabel("Msg")
	buf := bytes.NewBuffer(nil)
	fn := func(t *wdg.Y2hText) {
		buf.Reset()
		g.DumpTo(buf, t.Data,
			gutil.DumpOption{WithType: true})
		lblMsg.SetText(t.Txt1.Text + "\n" + buf.String())
	}
	type p struct {
		name string
		age  int
	}
	p1 := p{"tom", 11}
	txt1 := wdg.NewY2hText().SText1("txt111").
		TextColor1(colornames.Red).
		TextColor2(colornames.White).
		BgFillColor(colornames.Blue).
		BgStrokeColor(colornames.Red).
		SData(p1).
		Click(fn).
		Do(func(t *wdg.Y2hText) {
			t.Txt1.Text = "newtext1"
			t.Txt1.Alignment = fyne.TextAlignCenter
			t.Txt2.Alignment = fyne.TextAlignTrailing
		})
	txt2 := wdg.NewY2hText().SText1("txt222").
		TextColor1(colornames.Red).
		TextSize(40).
		BgFillColor(colornames.Yellow).
		BgStrokeColor(colornames.Red).
		Style(fyne.TextStyle{Bold: true}).
		TextAlign(fyne.TextAlignCenter).
		SData("i am txt2").
		Click(fn).Do(func(t *wdg.Y2hText) {
		t.Txt1.Text = "newtext2"
		t.Txt2.Text = ""
	})

	txt3 := wdg.NewY2hText().
		TextColor1(colornames.Red).
		BgFillColor(colornames.Burlywood).
		BgStrokeColor(colornames.Red).
		Style(fyne.TextStyle{Italic: true}).
		TextAlign(fyne.TextAlignTrailing).
		SData(33333).
		Click(fn).Do(func(t *wdg.Y2hText) {
		t.Txt1.Text = "newtext3"
	})
	txt4 := wdg.NewY2hText().
		TextColor1(colornames.Red).
		BgFillColor(colornames.Aliceblue).
		BgStrokeColor(colornames.Red).
		Style(fyne.TextStyle{Monospace: true}).
		TextAlign(fyne.TextAlignTrailing).
		SData(4.4).Do(func(t *wdg.Y2hText) {
		t.Txt1.Text = "newtext4"
	}).Click(fn)
	txt5 := wdg.NewY2hText().
		TextColor1(colornames.Red).
		BgFillColor(colornames.Antiquewhite).
		BgStrokeColor(colornames.Black).
		Style(fyne.TextStyle{TabWidth: 10}).
		TextAlign(fyne.TextAlignTrailing).
		SData([]string{"5", "4", "3"}).Do(func(t *wdg.Y2hText) {
		t.Txt1.Text = "new text	5"
	}).
		Click(fn)

	return container.NewVBox(lblMsg, txt1, txt2, txt3, txt4, txt5)

}
