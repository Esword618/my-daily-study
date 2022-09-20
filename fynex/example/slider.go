package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"gitee.com/y2h/goutil"

	xwdg "gitee.com/y2h/fynex/widget"
)

func slider() fyne.CanvasObject {
	slider := xwdg.NewSlider(1, 1000)
	slider.CircleScale(1.5).ColorValue(color.RGBA{255, 0, 0, 255}).
		ColorMin(goutil.HexToRGBA("999900")).
		ColorActive(goutil.HexToRGBA("#CC9933")).
		ColorTrack(goutil.HexToRGBA("#336666")).
		ColorCircle(color.RGBA{255, 255, 0, 255}).
		Do(func(s *xwdg.Slider) {
			s.Refresh()
		})

	slider1 := xwdg.NewSlider(1, 1000)
	slider.Orientation = xwdg.Vertical
	slider1.CircleScale(1.5).ColorValue(color.RGBA{255, 0, 0, 255}).
		ColorMin(goutil.HexToRGBA("999900")).
		ColorActive(goutil.HexToRGBA("#CC9933")).
		ColorTrack(goutil.HexToRGBA("#336666")).
		ColorCircle(color.RGBA{255, 255, 0, 255}).
		Do(func(s *xwdg.Slider) {
			s.Refresh()
		})
	return container.NewGridWithColumns(2, slider, slider1)
}
