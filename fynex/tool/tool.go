package tool

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Tool() fyne.CanvasObject {
	btn1 := widget.NewButton("Html to txt", func() {
		html2Txt()
	})
	btn2 := widget.NewButton("Fast Zip", func() {
		fastZip()
	})
	btn3 := widget.NewButton("gopsutil", func() {
		goPsUtil()
	})
	btn4 := widget.NewButton("caire", func() {
		// cairee()
	})

	btn5 := widget.NewButton("pigo", func() {
		pigoo()
	})

	btn6 := widget.NewButton("icons", func() {
		iconScreen()
	})

	btn7 := widget.NewButton("svg editor", func() {
		svge()
	})

	btn8 := widget.NewButton("Dominant Color", func() {
		dominantcolorUI()
	})

	btn9 := widget.NewButton("Json to struct", func() {
		json2Struct()
	})

	btn10 := widget.NewButton("sjson", func() {
		sjsonn()
	})

	btn11 := widget.NewButton("gjson", func() {
		gjsonn()
	})

	btn12 := widget.NewButton("Mapbox Vector Tiles", func() {
		mvtt()
	})

	btn13 := widget.NewButton("fc", func() {
		fcui()
	})
	btn14 := widget.NewButton("GenSvg", func() {
		gensvgui()
	})
	btn15 := widget.NewButton("OCR", func() {
		ocr()
	})
	btn16 := widget.NewButton("jacalz Eval", func() {
		jacalzEval()
	})

	btn17 := widget.NewButton("tdewolff/canvas", func() {
		tdewolff_canvas()
	})
	btn18 := widget.NewButton("theriault/maths", func() {
		theriault_maths()
	})
	grid := container.NewGridWithColumns(6, btn1, btn2, btn3, btn4, btn5, btn6,
		btn7, btn8, btn9, btn10, btn11, btn12, btn13, btn14, btn15, btn16,
		btn17, btn18)

	return grid
}
