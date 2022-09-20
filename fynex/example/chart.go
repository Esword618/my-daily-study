package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"gitee.com/y2h/fynex/chart"
	"golang.org/x/image/colornames"
)

func Chart() fyne.CanvasObject {
	btn1 := widget.NewButton("Bar Chart", func() {
		bar()
	})
	btn2 := widget.NewButton("Line Chart", func() {
		line2()
	})
	btn5 := widget.NewButton("Vertical Line Chart", func() {
		line()
	})
	btn3 := widget.NewButton("PolyLine Chart", func() {
		polyline()
	})
	btn4 := widget.NewButton("Scatter Chart", func() {
		scatter()
	})
	grid := container.NewGridWithColumns(6, btn1, btn2, btn5, btn3, btn4)

	return grid
}

func bar() {
	w := fyne.CurrentApp().NewWindow("Bar Chart")
	bar1 := chart.NewBar().
		Data([]float32{120, -20, 30, 40, 60, 180, 240}).
		BarWidth(30).BarColor(colornames.Fuchsia)

	w.SetContent(bar1)
	w.CenterOnScreen()
	w.Resize(fyne.NewSize(700, 700))
	w.Show()
}
func line() {
	w := fyne.CurrentApp().NewWindow("Line Chart")
	line1 := chart.NewLine(true).Data(120, 20, 30, -40, 60, 180, 240).
		LineWidth(11).LineColor(colornames.Blueviolet)
	w.SetContent(line1)
	w.CenterOnScreen()
	w.Resize(fyne.NewSize(700, 700))
	w.Show()
}
func line2() {
	w := fyne.CurrentApp().NewWindow("Line Chart")
	line1 := chart.NewLine(false).Data(120, 20, 30, -40, 60, 180, 240).
		LineWidth(11).LineColor(colornames.Blueviolet)
	w.SetContent(line1)
	w.CenterOnScreen()
	w.Resize(fyne.NewSize(700, 700))
	w.Show()
}
func polyline() {
	w := fyne.CurrentApp().NewWindow("Line Chart")
	line1 := chart.NewPolyLine().Data(120, 20, 30, -40, 60, 180, 240).
		LineWidth(11).LineColor(colornames.Blueviolet)
	w.SetContent(line1)
	w.CenterOnScreen()
	w.Resize(fyne.NewSize(700, 700))
	w.Show()
}
func scatter() {
	w := fyne.CurrentApp().NewWindow("Scatter Chart")
	line1 := chart.NewScatter().Data(120, 20, 30, -40, 60, 180, 240)
	w.SetContent(line1)
	w.CenterOnScreen()
	w.Resize(fyne.NewSize(700, 700))
	w.Show()
}
