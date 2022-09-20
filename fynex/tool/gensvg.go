package tool

import (
	"bytes"
	"fmt"
	"image"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/ajstarks/gensvg"
	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
)

var width = 500.0
var height = 500.0

func gensvgui() {
	var d []byte

	buf := bytes.NewBuffer(nil)
	l := &widget.Label{
		Wrapping: fyne.TextTruncate,
	}
	v := &canvas.Raster{
		Generator: func(width, height int) image.Image {
			l.SetText("")
			output := image.NewNRGBA(image.Rect(0, 0, width, height))
			if len(d) == 0 {
				l.SetText("No Data")
				return output
			}
			icon, err := oksvg.ReadIconStream(bytes.NewReader(d))
			if err != nil {
				l.SetText(err.Error())
				return output
			}
			inputW, inputH := icon.ViewBox.W, icon.ViewBox.H
			iconAspect := inputW / inputH
			viewAspect := float64(width) / float64(height)
			outputW, outputH := width, height
			if viewAspect > iconAspect {
				outputW = int(float64(height) * iconAspect)
			} else if viewAspect < iconAspect {
				outputH = int(float64(width) / iconAspect)
			}
			scanner := rasterx.NewScannerGV(int(inputW), int(inputH), output, output.Bounds())
			raster := rasterx.NewDasher(width, height, scanner)
			icon.SetTarget(0, 0, float64(outputW), float64(outputH))
			icon.Draw(raster, 1)
			defer func() {
				if r := recover(); r != nil {
					l.SetText(fmt.Sprintf("Crash when rendering SVG: %v", r))
				}
			}()
			return output
		},
	}

	btn1 := widget.NewButton("Circle", func() {
		d = circle(buf)
		v.Refresh()
	})

	btn2 := widget.NewButton("Ellipse", func() {
		d = ellipse(buf)
		v.Refresh()
	})
	btn3 := widget.NewButton("Polygon", func() {
		d = polygon(buf)
		v.Refresh()
	})

	btn4 := widget.NewButton("Rect", func() {
		d = rect(buf)
		v.Refresh()
	})

	btn5 := widget.NewButton("Text", func() {
		d = text1(buf)
		v.Refresh()
	})
	btn6 := widget.NewButton("Grid", func() {
		d = grid(buf)
		v.Refresh()
	})
	w := fyne.CurrentApp().NewWindow("GenSvg https://hub.fastgit.xyz/ajstarks/gensvg")

	cc := container.NewBorder(nil, l, container.NewVBox(btn1, btn2, btn3, btn4, btn5,btn6), nil,
		container.NewScroll(v))
	w.SetContent(cc)
	w.CenterOnScreen()
	w.Resize(fyne.NewSize(700, 700))
	w.Show()
}

func circle(buf *bytes.Buffer) []byte {
	buf.Reset()
	canvas0 := gensvg.New(buf)
	// canvas0 := gensvg.New(os.Stdout)
	canvas0.Start(width, height)
	canvas0.Circle(width/2, height/2, 100)
	canvas0.Text(width/2, height/2, "Hello, SVG", "text-anchor:middle;font-size:30px;fill:white")
	canvas0.End()
	return buf.Bytes()
}
func ellipse(buf *bytes.Buffer) []byte {
	buf.Reset()
	canvas0 := gensvg.New(buf)
	canvas0.Start(width, height)
	canvas0.Ellipse(width/2, height/2, width/3, height/4, "fill:red",
		"stroke:blue")
	canvas0.End()
	return buf.Bytes()
}
func polygon(buf *bytes.Buffer) []byte {
	buf.Reset()
	canvas0 := gensvg.New(buf)
	canvas0.Start(width, height)
	canvas0.Polygon([]float64{350, 379, 469, 397, 423, 350, 277, 303, 231, 321},
		[]float64{75, 161, 161, 215, 301, 250, 301, 215, 161, 161}, "fill:red",
		"stroke:blue")

	canvas0.Polygon([]float64{150, 258, 258, 150, 42, 42},
		[]float64{75, 137.5, 262.5, 325, 262.6, 137.5}, "fill:red",
		"stroke:blue", "stroke-width:10")
	canvas0.End()
	return buf.Bytes()
}
func rect(buf *bytes.Buffer) []byte {
	buf.Reset()
	canvas0 := gensvg.New(buf)
	canvas0.Start(width, height)
	canvas0.Rect(10, 10, 100, 50, "fill:red", "stroke:blue")
	canvas0.CenterRect(30, 100, 20, 10, "fill:red", "stroke:blue")
	canvas0.Roundrect(10, 150, 50, 40, 10, 10, "fill:red", "stroke:blue")
	canvas0.End()
	return buf.Bytes()
}
func text1(buf *bytes.Buffer) []byte {
	buf.Reset()
	canvas0 := gensvg.New(buf)
	canvas0.Start(width, height)
	canvas0.Text(width/2, height/2, "Hello, out there", "fill:none", "stroke:blue", "stroke-width:10",
		"font-size:55", "font-family:Verdana")
	canvas0.End()
	return buf.Bytes()
}
func grid(buf *bytes.Buffer) []byte {
	buf.Reset()
	canvas0 := gensvg.New(buf)
	canvas0.Start(width, height)
	canvas0.Grid(10, 10, 400, 400, 40, "stroke:blue", "stroke-width:10")
	canvas0.End()
	return buf.Bytes()
}
