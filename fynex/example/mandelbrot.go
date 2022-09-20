// https://github.com/soniakeys/mandelbrot/blob/master/mandelbrot.go
package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math/cmplx"
	"os"
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"gitee.com/y2h/fynex/sched"
	xwdg "gitee.com/y2h/fynex/widget"
	"golang.org/x/image/colornames"
)

const (
	// maxEsc = 100
	// rMin  = -2.
	// rMax  = .5
	// iMin  = -1.
	// iMax  = 1.
	width = 750
	// red   = 230
	// green = 235
	// blue  = 255
)

type mandelbrotS struct {
	sched *sched.Pool
}

var (
	maxEsc float64 = 100.
	rMin           = -2.
	rMax           = .5

	iMin = -1.
	iMax = 1.
	red  = 100.
	green = 149.
	blue  = 237.
	b    *image.NRGBA
	me   = mandelbrotS{}
)

func mandelbrot(a complex128) float64 {
	i := 0
	n := int(maxEsc)
	for z := a; cmplx.Abs(z) < 2 && i < n; i++ {
		z = z*z + a
	}
	return (maxEsc - float64(i)) / maxEsc
}

func mandelbrotUI() fyne.CanvasObject {
	me.sched = sched.New(sched.Workers(runtime.NumCPU()))
	runtime.SetFinalizer(&me, func(r *mandelbrotS) {
		r.sched.Release()
	})
	img := canvas.NewImageFromImage(nil)
	slider := xwdg.NewSlider(10, 200).Do(func(s *xwdg.Slider) {
		s.Step = 10
		s.Caption = "maxEsc"
		s.ColorTrack(colornames.Aliceblue)
		s.ColorActive(colornames.Beige)
		s.OnChanged = func(f float64) {
			maxEsc = f
		}
		s.SetValue(100)
	})
	sliderRMin := xwdg.NewSlider(-10, 10).Do(func(s *xwdg.Slider) {
		s.Step = 1
		s.Caption = "rMin"
		s.ColorTrack(colornames.Antiquewhite)
		s.ColorActive(colornames.Bisque)
		s.OnChanged = func(f float64) {
			rMin = f
			// refresh(img)
		}
		s.SetValue(-2)
	})
	sliderRMax := xwdg.NewSlider(0.1, 1).Do(func(s *xwdg.Slider) {
		s.Step = 0.1
		s.Caption = "rMax"
		s.ColorTrack(colornames.Aqua)
		s.ColorActive(colornames.Burlywood)
		s.OnChanged = func(f float64) {
			rMax = f
			// refresh(img)
		}
		s.SetValue(.5)
	})
	sliderIMin := xwdg.NewSlider(-10, 0).Do(func(s *xwdg.Slider) {
		s.Step = 1
		s.Caption = "iMin"
		s.ColorTrack(colornames.Antiquewhite)
		s.ColorActive(colornames.Bisque)
		s.OnChanged = func(f float64) {
			iMin = f
			// refresh(img)
		}
		s.SetValue(-1)
	})
	sliderIMax := xwdg.NewSlider(1, 10).Do(func(s *xwdg.Slider) {
		s.Step = 1
		s.Caption = "iMax"
		s.ColorTrack(colornames.Aqua)
		s.ColorActive(colornames.Burlywood)
		s.OnChanged = func(f float64) {
			iMax = f
			// refresh(img)
		}
		s.SetValue(1)
	})
	sliderRed := xwdg.NewSlider(0, 255).Do(func(s *xwdg.Slider) {
		s.Step = 25
		s.Caption = "Red"
		s.ColorTrack(colornames.Aqua)
		s.ColorActive(colornames.Burlywood)
		s.OnChanged = func(f float64) {
			red = f
			// refresh(img)
		}
		s.SetValue(100)
	})
	sliderGreen := xwdg.NewSlider(0, 255).Do(func(s *xwdg.Slider) {
		s.Step = 25
		s.Caption = "Green"
		s.ColorTrack(colornames.Aqua)
		s.ColorActive(colornames.Burlywood)
		s.OnChanged = func(f float64) {
			green = f
			// refresh(img)
		}
		s.SetValue(149)
	})
	sliderBlue := xwdg.NewSlider(0, 255).Do(func(s *xwdg.Slider) {
		s.Step = 25
		s.Caption = "Blue"
		s.ColorTrack(colornames.Aqua)
		s.ColorActive(colornames.Burlywood)
		s.OnChanged = func(f float64) {
			blue = f
			// refresh(img)
		}
		s.SetValue(237)
	})
	lblMsg := widget.NewLabel("")
	btnDo := widget.NewButton("Generate mandelbrot image", func() {
		refresh1(img)
		refresh(img)
	})
	btnSave := widget.NewButton("Save", func() {
		f, err := os.Create("mandelbrot.png")
		if err != nil {
			lblMsg.SetText(err.Error())
			return
		}
		if err = png.Encode(f, b); err != nil {
			lblMsg.SetText(err.Error())
		}
		if err = f.Close(); err != nil {
			lblMsg.SetText(err.Error())
		}
	})

	btnDefault := widget.NewButton("Default", func() {
		slider.SetValue(100)
		sliderRMin.SetValue(-2)
		sliderRMax.SetValue(.5)
		sliderIMin.SetValue(-1)
		sliderIMax.SetValue(1.)
		sliderRed.SetValue(230.)

	})
	return container.NewBorder(container.NewVBox(sliderRed,sliderGreen,
		sliderBlue,
		slider, sliderRMin, sliderRMax, sliderIMin, sliderIMax,
		container.NewGridWithColumns(3, btnDo, btnDefault, btnSave),
	),
		lblMsg, nil, nil, img)
}
func refresh(img *canvas.Image) {
	scale := width / (rMax - rMin)
	height := int(scale * (iMax - iMin))
	bounds := image.Rect(0, 0, width, height)
	b = image.NewNRGBA(bounds)
	draw.Draw(b, bounds, image.NewUniform(color.Black), image.ZP, draw.Src)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			fEsc := mandelbrot(complex(
				float64(x)/scale+rMin,
				float64(y)/scale+iMin))
			b.Set(x, y, color.NRGBA{uint8(red * fEsc),
				uint8(green * fEsc), uint8(blue * fEsc), 255})

		}
	}

	img.Image = b
	img.Refresh()
}
func refresh1(img *canvas.Image) {
	img.Image = nil
	// sched := sched.New(sched.Workers(runtime.NumCPU()))
	// defer sched.Release()

	scale := width / (rMax - rMin)
	height := int(scale * (iMax - iMin))
	bounds := image.Rect(0, 0, width, height)
	b = image.NewNRGBA(bounds)
	draw.Draw(b, bounds, image.NewUniform(color.Black), image.ZP, draw.Src)

	me.sched.Add(width)
	for x := 0; x < width; x++ {
		me.sched.Run(func() {
			for y := 0; y < height; y++ {
				fEsc := mandelbrot(complex(
					float64(x)/scale+rMin,
					float64(y)/scale+iMin))
				b.Set(x, y, color.NRGBA{uint8(red * fEsc),
					uint8(green * fEsc), uint8(blue * fEsc), 255})

			}
		})
	}
	me.sched.Wait()

	img.Image = b
	img.Refresh()
}
