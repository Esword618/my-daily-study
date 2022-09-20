package main

import (
	"image"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/fogleman/ease"
	"github.com/fogleman/gg"
)

const P = 32 / 2
const Px = 256 / 2
const W = 1024/2 + Px*1
const H = 0 + P*2
const N = 1024 / 2

func Easy() fyne.CanvasObject {
	var img1 = canvas.NewImageFromImage(nil)
	// img1.FillMode = canvas.ImageFillOriginal
	img2 := canvas.NewImageFromImage(nil)
	// img2.FillMode = canvas.ImageFillOriginal
	img3 := canvas.NewImageFromImage(nil)
	// img3.FillMode = canvas.ImageFillOriginal
	img4 := canvas.NewImageFromImage(nil)
	// img4.FillMode = canvas.ImageFillOriginal
	img5 := canvas.NewImageFromImage(nil)
	// img5.FillMode = canvas.ImageFillOriginal
	img6 := canvas.NewImageFromImage(nil)
	// img6.FillMode = canvas.ImageFillOriginal
	img7 := canvas.NewImageFromImage(nil)
	// img7.FillMode = canvas.ImageFillOriginal
	img8 := canvas.NewImageFromImage(nil)
	// img8.FillMode = canvas.ImageFillOriginal
	img9 := canvas.NewImageFromImage(nil)
	// img9.FillMode = canvas.ImageFillOriginal
	img10 := canvas.NewImageFromImage(nil)
	// img10.FillMode = canvas.ImageFillOriginal
	img11 := canvas.NewImageFromImage(nil)
	// img11.FillMode = canvas.ImageFillOriginal
	img12 := canvas.NewImageFromImage(nil)
	// img12.FillMode = canvas.ImageFillOriginal
	img13 := canvas.NewImageFromImage(nil)
	// img13.FillMode = canvas.ImageFillOriginal
	img14 := canvas.NewImageFromImage(nil)
	// img14.FillMode = canvas.ImageFillOriginal
	img15 := canvas.NewImageFromImage(nil)
	// img15.FillMode = canvas.ImageFillOriginal
	img16 := canvas.NewImageFromImage(nil)
	// img16.FillMode = canvas.ImageFillOriginal
	img17 := canvas.NewImageFromImage(nil)
	// img17.FillMode = canvas.ImageFillOriginal
	img18 := canvas.NewImageFromImage(nil)
	// img18.FillMode = canvas.ImageFillOriginal
	img19 := canvas.NewImageFromImage(nil)
	// img19.FillMode = canvas.ImageFillOriginal
	img20 := canvas.NewImageFromImage(nil)
	// img20.FillMode = canvas.ImageFillOriginal
	img21 := canvas.NewImageFromImage(nil)
	// img21.FillMode = canvas.ImageFillOriginal
	img22 := canvas.NewImageFromImage(nil)
	// img22.FillMode = canvas.ImageFillOriginal
	img23 := canvas.NewImageFromImage(nil)
	// img23.FillMode = canvas.ImageFillOriginal
	img24 := canvas.NewImageFromImage(nil)
	// img24.FillMode = canvas.ImageFillOriginal
	img25 := canvas.NewImageFromImage(nil)
	// img25.FillMode = canvas.ImageFillOriginal
	img26 := canvas.NewImageFromImage(nil)
	// img26.FillMode = canvas.ImageFillOriginal
	img27 := canvas.NewImageFromImage(nil)
	// img27.FillMode = canvas.ImageFillOriginal
	img28 := canvas.NewImageFromImage(nil)
	// img28.FillMode = canvas.ImageFillOriginal
	img29 := canvas.NewImageFromImage(nil)
	// img29.FillMode = canvas.ImageFillOriginal
	img30 := canvas.NewImageFromImage(nil)
	// img30.FillMode = canvas.ImageFillOriginal
	img31 := canvas.NewImageFromImage(nil)
	// img31.FillMode = canvas.ImageFillOriginal
	const n = 100
	btnDo := widget.NewButton("Do", func() {
		for i := 0; i < n; i++ {
			t := float64(i) / (n - 1)
			img1.Image = render(ease.Linear, t)
			img1.Refresh()
			img2.Image = render(ease.InQuad, t)
			img2.Refresh()
			img3.Image = render(ease.InCubic, t)
			img3.Refresh()
			img4.Image = render(ease.InQuart, t)
			img4.Refresh()
			img5.Image = render(ease.InQuint, t)
			img5.Refresh()
			img6.Image = render(ease.InSine, t)
			img6.Refresh()
			img7.Image = render(ease.InExpo, t)
			img7.Refresh()
			img8.Image = render(ease.InCirc, t)
			img8.Refresh()
			img9.Image = render(ease.InElastic, t)
			img9.Refresh()
			img10.Image = render(ease.InBack, t)
			img10.Refresh()
			img11.Image = render(ease.InBounce, t)
			img11.Refresh()
			img12.Image = render(ease.OutQuad, t)
			img12.Refresh()
			img13.Image = render(ease.OutCubic, t)
			img13.Refresh()
			img14.Image = render(ease.OutQuart, t)
			img14.Refresh()
			img15.Image = render(ease.OutQuint, t)
			img15.Refresh()
			img16.Image = render(ease.OutSine, t)
			img16.Refresh()
			img17.Image = render(ease.OutExpo, t)
			img17.Refresh()
			img18.Image = render(ease.OutCirc, t)
			img18.Refresh()
			img19.Image = render(ease.OutElastic, t)
			img19.Refresh()
			img20.Image = render(ease.OutBack, t)
			img20.Refresh()
			img21.Image = render(ease.OutBounce, t)
			img21.Refresh()
			img22.Image = render(ease.InOutQuad, t)
			img22.Refresh()
			img23.Image = render(ease.InOutCubic, t)
			img23.Refresh()
			img24.Image = render(ease.InOutQuart, t)
			img24.Refresh()
			img25.Image = render(ease.InOutQuint, t)
			img25.Refresh()
			img26.Image = render(ease.InOutSine, t)
			img26.Refresh()
			img27.Image = render(ease.InOutExpo, t)
			img27.Refresh()
			img28.Image = render(ease.InOutCirc, t)
			img28.Refresh()
			img29.Image = render(ease.InOutElastic, t)
			img29.Refresh()
			img30.Image = render(ease.InOutBack, t)
			img30.Refresh()
			img31.Image = render(ease.InOutBounce, t)
			img31.Refresh()

			time.Sleep(time.Millisecond * 50)
		}
	})
	return container.NewAppTabs(container.NewTabItem(
		"Gif", container.NewVScroll(container.
			NewGridWithColumns(3, btnDo, img1, img2, img3, img4, img5, img6, img7, img8, img9,
				img10, img11, img12, img13, img14, img15, img16,
				img17, img18, img19, img20, img21, img22,
				img23, img24, img25, img26, img27, img28, img29, img30, img31))),
		container.NewTabItem("Axis", Easy1()))

}
func render(function ease.Function, t float64) image.Image {
	dc := gg.NewContext(W-50, H)
	dc.InvertY()
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.Translate(Px, P)
	dc.Scale(W-Px*2, 1)
	// draw minor grid
	for i := 1; i <= 10; i++ {
		x := float64(i) / 10
		dc.MoveTo(x, -P/2)
		dc.LineTo(x, P/2)
	}
	dc.SetRGBA(0, 0, 0, 0.25)
	dc.SetLineWidth(2)
	dc.Stroke()
	// draw axes
	dc.MoveTo(0, 0)
	dc.LineTo(1, 0)
	dc.MoveTo(0, -P/2)
	dc.LineTo(0, P/2)
	dc.MoveTo(1, -P/2)
	dc.LineTo(1, P/2)
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(3)
	dc.Stroke()
	// draw points
	x := function(t)
	dc.DrawPoint(x, 0, 16)
	dc.SetRGB(0, 0, 0)
	dc.Fill()
	dc.DrawPoint(x, 0, 12)
	dc.SetRGB(1, 0, 0)
	dc.Fill()
	return dc.Image()
}
func Easy1() fyne.CanvasObject {
	img1 := canvas.NewImageFromImage(render1(ease.Linear, "ease.Linear(t)"))
	img2 := canvas.NewImageFromImage(render1(ease.InQuad, "ease.InQuad(t)"))
	img3 := canvas.NewImageFromImage(render1(ease.OutQuad, "ease.OutQuad(t)"))
	img4 := canvas.NewImageFromImage(render1(ease.InOutQuad, "ease.InOutQuad(t)"))
	img5 := canvas.NewImageFromImage(render1(ease.InCubic, "ease.InCubic(t)"))
	img6 := canvas.NewImageFromImage(render1(ease.OutCubic, "ease.OutCubic(t)"))
	img7 := canvas.NewImageFromImage(render1(ease.InOutCubic, "ease.InOutCubic(t)"))
	img8 := canvas.NewImageFromImage(render1(ease.InQuart, "ease.InQuart(t)"))
	img9 := canvas.NewImageFromImage(render1(ease.OutQuart, "ease.OutQuart(t)"))
	img10 := canvas.NewImageFromImage(render1(ease.InOutQuart, "ease.InOutQuart(t)"))
	img11 := canvas.NewImageFromImage(render1(ease.InQuint, "ease.InQuint(t)"))
	img12 := canvas.NewImageFromImage(render1(ease.OutQuint, "ease.OutQuint(t)"))
	img13 := canvas.NewImageFromImage(render1(ease.InOutQuint, "ease.InOutQuint(t)"))
	img14 := canvas.NewImageFromImage(render1(ease.InSine, "ease.InSine(t)"))
	img15 := canvas.NewImageFromImage(render1(ease.OutSine, "ease.OutSine(t)"))
	img16 := canvas.NewImageFromImage(render1(ease.InOutSine, "ease.InOutSine(t)"))
	img17 := canvas.NewImageFromImage(render1(ease.InExpo, "ease.InExpo(t)"))
	img18 := canvas.NewImageFromImage(render1(ease.OutExpo, "ease.OutExpo(t)"))
	img19 := canvas.NewImageFromImage(render1(ease.InOutExpo, "ease.InOutExpo(t)"))
	img20 := canvas.NewImageFromImage(render1(ease.InCirc, "ease.InCirc(t)"))
	img21 := canvas.NewImageFromImage(render1(ease.OutCirc, "ease.OutCirc(t)"))
	img22 := canvas.NewImageFromImage(render1(ease.InOutCirc, "ease.InOutCirc(t)"))
	img23 := canvas.NewImageFromImage(render1(ease.InElastic, "ease.InElastic(t)"))
	img24 := canvas.NewImageFromImage(render1(ease.OutElastic, "ease.OutElastic(t)"))
	img25 := canvas.NewImageFromImage(render1(ease.InOutElastic, "ease.InOutElastic(t)"))
	img26 := canvas.NewImageFromImage(render1(ease.InBack, "ease.InBack(t)"))
	img27 := canvas.NewImageFromImage(render1(ease.OutBack, "ease.OutBack(t)"))
	img28 := canvas.NewImageFromImage(render1(ease.InOutBack, "ease.InOutBack(t)"))
	img29 := canvas.NewImageFromImage(render1(ease.InBounce, "ease.InBounce(t)"))
	img30 := canvas.NewImageFromImage(render1(ease.OutBounce, "ease.OutBounce(t)"))
	img31 := canvas.NewImageFromImage(render1(ease.InOutBounce, "ease.InOutBounce(t)"))
	img32 := canvas.NewImageFromImage(render1(ease.InOutBounce, "ease.InOutBounce(t)"))
	img32.FillMode = canvas.ImageFillOriginal
	return container.NewVScroll(container.NewGridWithColumns(4,
		img1, img2, img3, img4, img5, img6, img7, img8, img9,
		img10, img11, img12, img13, img14, img15, img16,
		img17, img18, img19, img20, img21, img22,
		img23, img24, img25, img26, img27, img28, img29, img30, img31, img32))
}
func render1(function ease.Function, title string) image.Image {
	const P = 256 / 4
	const S = 1024/4 + P*2/4
	const N = 1024 / 4
	// const S = W
	dc := gg.NewContext(S, S)
	dc.InvertY()
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.Translate(P, P)
	dc.Scale(S-P*2, S-P*2)
	// draw minor grid
	for i := 1; i <= 10; i++ {
		x := float64(i) / 10
		dc.MoveTo(x, 0)
		dc.LineTo(x, 1)
		dc.MoveTo(0, x)
		dc.LineTo(1, x)
	}
	dc.SetRGBA(0, 0, 0, 0.25)
	dc.SetLineWidth(2)
	dc.Stroke()
	// draw axes
	dc.MoveTo(0, 0)
	dc.LineTo(1, 0)
	dc.LineTo(1, 1)
	dc.LineTo(0, 1)
	dc.ClosePath()
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(3)
	dc.Stroke()
	// draw points
	for i := 0; i < N; i++ {
		x := float64(i) / (N - 1)
		y := function(x)
		dc.LineTo(x, y)
	}
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(12)
	dc.StrokePreserve()
	dc.SetRGB(1, 0, 0)
	dc.SetLineWidth(6)
	dc.Stroke()
	// title
	dc.Identity()
	dc.SetRGB(0, 0, 0)
	// if err := dc.LoadFontFace("/Users/fogleman/Library/Fonts/DejaVuSansMono.ttf", 48); err != nil {
	// 	panic(err)
	// }
	dc.DrawStringAnchored(title, S/2, P/4, 0.5, 0.5)
	return dc.Image()
	// dc.SavePNG(path)
}
