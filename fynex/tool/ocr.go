package tool

import (
	"context"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/gogf/gf/v2/text/gstr"

	. "gitee.com/y2h/fynex/widget"
)

func ocr() {
	w := fyne.CurrentApp().NewWindow("OCR tesseract")
	exepath := widget.NewEntry()
	msg := widget.NewLabel("")
	right := widget.NewMultiLineEntry()
	right.Wrapping = fyne.TextWrapBreak
	exepath.SetText(`D:\chenhao\soft\Tesseract-OCR\tesseract.exe`)
	img := NewClickImage()
	img.Img.FillMode = canvas.ImageFillContain

	btnDo := widget.NewButton("Ocr", func() {
		exe := exepath.Text + ` ` + img.Img.File + ` - -l chi_sim`
		msg.SetText(exe)
		r, err := gproc.ShellExec(context.Background(), exe)
		if err != nil {
			right.SetText(err.Error())
			return
		}
		r = gstr.TrimAll(r, " ")
		right.SetText(r)
	})
	top := container.NewVBox(exepath, btnDo)
	split := container.NewHSplit(img, right)
	w.SetContent(container.NewBorder(top, msg, nil, nil, split))
	w.Resize(fyne.NewSize(800, 700))
	w.Show()
}
