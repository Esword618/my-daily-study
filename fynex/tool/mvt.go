package tool

import (
	"bytes"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"github.com/tidwall/mvt"
)

func mvtt() {
	var tile mvt.Tile
	l := tile.AddLayer("triforce")
	f := l.AddFeature(mvt.Polygon)

	f.MoveTo(128, 96)
	f.LineTo(148, 128)
	f.LineTo(108, 128)
	f.LineTo(128, 96)
	f.ClosePath()

	f.MoveTo(148, 128)
	f.LineTo(168, 160)
	f.LineTo(128, 160)
	f.LineTo(148, 128)
	f.ClosePath()

	f.MoveTo(108, 128)
	f.LineTo(128, 160)
	f.LineTo(88, 160)
	f.LineTo(108, 128)
	f.ClosePath()

	data := tile.Render()

	buf := bytes.NewBuffer(data)
	// buf := bufio.NewReader(nil)
	// buf.Read(data)
	w := fyne.CurrentApp().NewWindow("Mapbox Vector Tiles")
	img := canvas.NewImageFromReader(buf, "newimage")
	w.SetContent(img)
	w.CenterOnScreen()
	w.Resize(fyne.NewSize(700, 700))
	w.Show()
}
