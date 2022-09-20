package tool

import (
	"image/color"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/ajstarks/fc"
)

func fcui() {
	width, height := 500, 500
	r := height / 2
	rand.Seed(time.Now().Unix())
	black := color.RGBA{0, 0, 0, 255}
	white := color.RGBA{255, 255, 255, 255}
	blue := color.RGBA{0, 0, 255, 255}
	red := color.RGBA{255, 0, 0, 255}
	yellow := color.RGBA{255, 255, 0, 255}
	canvas := container.NewWithoutLayout()
	btn1 := widget.NewButton("Concentric", func() {
		canvas.RemoveAll()
		for g := uint8(0); g < 250; g += 50 {
			fc.AbsCircle(canvas, width/2, height/2, r, color.RGBA{g, g, g, 255})
			r -= 50
		}
	})
	btn2 := widget.NewButton("Mondrian", func() {
		canvas.RemoveAll()

		var border float32 = 6.0
		w3 := width / 3
		w6 := w3 / 2
		w23 := w3 * 2
		w36 := w3 - w6
		ww6 := width - w6
		hw6 := height - w6
		hw3 := height - w3
		b2 := border * 2

		fc.AbsCornerRect(canvas, 0, 0, w3, w3, white)       // upper left white square
		fc.AbsCornerRect(canvas, 0, w3, w3, w3, white)      // middle left white square
		fc.AbsCornerRect(canvas, 0, w23, w3, w3, blue)      // lower left blue square
		fc.AbsCornerRect(canvas, w3, 0, w23, w23, red)      // large red square
		fc.AbsCornerRect(canvas, w3, w23, w23, w3, white)   // lower-middle white rectangle
		fc.AbsCornerRect(canvas, ww6, hw3, w36, w6, white)  // lower right white square
		fc.AbsCornerRect(canvas, ww6, hw6, w36, w6, yellow) // lower right yellow square

		fc.AbsLine(canvas, 0, 0, 0, height, b2, black)          // left border
		fc.AbsLine(canvas, width, 0, width, height, b2, black)  // right border
		fc.AbsLine(canvas, 0, 0, width, 0, b2, black)           // top border
		fc.AbsLine(canvas, 0, height, width, height, b2, black) // botom border

		fc.AbsLine(canvas, 0, w3, w3, w3, border, black)         // bottom of upper left white square
		fc.AbsLine(canvas, w3, 0, w3, height, border, black)     // right border for left-hand squares
		fc.AbsLine(canvas, 0, w23, width, w23, border, black)    // two-thirds border
		fc.AbsLine(canvas, ww6, hw3, ww6, height, border, black) // left border for small squares
		fc.AbsLine(canvas, ww6, hw6, width, hw6, border, black)  // top/bottom of small squares
	})
	btn3 := widget.NewButton("mondrian", func() {
		canvas.RemoveAll()
		earthsize := 4
		sunsize := earthsize * 109
		fc.AbsCircle(canvas, 150, 50, earthsize, blue)
		fc.AbsCircle(canvas, width, height, sunsize, yellow)

		for i := 0; i < width; i++ {
			r := uint8(rand.Intn(255))
			c := color.RGBA{r, r, r, 255}
			fc.AbsLine(canvas, i, 0, rand.Intn(width), height, 5, c)
		}
	})

	btn4 := widget.NewButton("Sun+Earth", func() {
		canvas.RemoveAll()
		for i := 0; i < width; i++ {
			x, y := rand.Intn(width), rand.Intn(height)
			fc.AbsLine(canvas, x, y, x, y+1, 0.4, white)
		}
	})
	top := container.NewHBox(btn1, btn2, btn3, btn4)
	cc := container.NewBorder(top, nil, nil, nil, canvas)

	win := fyne.CurrentApp().NewWindow("fc - high-level canvas for the fyne toolkit")
	win.SetContent(cc)
	win.CenterOnScreen()
	win.Resize(fyne.NewSize(700, 700))
	win.Show()
}
