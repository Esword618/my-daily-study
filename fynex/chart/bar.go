package chart

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/image/colornames"

	xwdg "gitee.com/y2h/fynex/widget"
)

var _ fyne.Widget = (*Bar)(nil)

// Bar is a widget for displaying a separator with themeable color.
//
// Since: 1.4
type Bar struct {
	widget.BaseWidget
	W, H       float32
	barWidth   float32
	data       []float32
	barColor   color.Color
	X          *canvas.Line
	Y          *canvas.Line
	T1         *canvas.Text

	Center   fyne.Position
	ChartCtn *fyne.Container

	xgrid *fyne.Container
	ygrid *fyne.Container
}

// NewBar creates a new separator.
//
// Since: 1.4
func NewBar() *Bar {
	s := &Bar{}
	s.ExtendBaseWidget(s)
	s.T1 = canvas.NewText("", colornames.Blue)
	s.X = canvas.NewLine(colornames.Blue)
	s.X.StrokeWidth = 3
	s.Y = canvas.NewLine(colornames.Red)
	s.Y.StrokeWidth = 3
	s.xgrid = container.NewMax()
	s.ygrid = container.NewMax()
	s.ChartCtn = container.NewMax()
	for i := 0; i < 10; i++ {
		s.xgrid.Add(canvas.NewLine(color.Black))
		s.ygrid.Add(canvas.NewLine(color.Black))
	}
	return s
}

func (s *Bar) Data(data []float32) *Bar {
	s.data = data
	s.ChartCtn.Objects = make([]fyne.CanvasObject, len(data))
	for i, v := range data {
		line := xwdg.NewBar().StrokeColor(colornames.Chocolate).
			StrokeWidth(3).SText1(fmt.Sprintf("%.2f", v))

		s.ChartCtn.Objects[i] = line
	}
	return s
}
func (s *Bar) Do(fn func(*Bar)) *Bar {
	fn(s)
	return s
}
func (s *Bar) BarWidth(w float32) *Bar {
	s.barWidth = w
	return s
}
func (s *Bar) BarColor(c color.Color) *Bar {
	s.barColor = c
	return s
}
func (s *Bar) Tapped(e *fyne.PointEvent) {

}
func (s *Bar) MouseIn(*desktop.MouseEvent) {
}

var r1 float32 = 10

// MouseMoved is called when a desktop pointer hovers over the widget
func (s *Bar) MouseMoved(e *desktop.MouseEvent) {

}

// MouseOut is called when a desktop pointer exits the widget
func (s *Bar) MouseOut() {
}

// CreateRenderer returns a new renderer for the separator.
//
// Implements: fyne.Widget
func (s *Bar) CreateRenderer() fyne.WidgetRenderer {
	s.ExtendBaseWidget(s)
	ctn := container.NewMax(s.xgrid, s.ygrid, s.X, s.Y, s.T1,
		s.ChartCtn)
	return &barRenderer{
		WidgetRenderer: widget.NewSimpleRenderer(ctn),
		d:              s,
	}
}

// MinSize returns the minimal size of the separator.
//
// Implements: fyne.Widget
func (s *Bar) MinSize() fyne.Size {
	s.ExtendBaseWidget(s)
	return fyne.NewSize(600, 600)
}

var _ fyne.WidgetRenderer = (*barRenderer)(nil)

type barRenderer struct {
	fyne.WidgetRenderer
	d *Bar
}

func (r *barRenderer) Layout(s fyne.Size) {
	r.d.W = s.Width
	r.d.H = s.Height
	r.d.Center = fyne.NewPos(s.Width/2, s.Height/2)

	for i, obj := range r.d.xgrid.Objects {
		line := obj.(*canvas.Line)
		line.Position1 = fyne.NewPos(0, float32(i)*s.Height/10)
		line.Position2 = fyne.NewPos(s.Width, float32(i)*s.Height/10)
	}
	for i, obj := range r.d.ygrid.Objects {
		line := obj.(*canvas.Line)
		line.Position1 = fyne.NewPos(float32(i)*s.Width/10, 0)
		line.Position2 = fyne.NewPos(float32(i)*s.Width/10, s.Height)
	}

	step := r.d.W / float32(len(r.d.data))
	for i, obj := range r.d.ChartCtn.Objects {
		bar := obj.(*xwdg.Bar)
		bar.FillColor(r.d.barColor)

		if r.d.data[i] < 0 {
			bar.Resize(fyne.NewSize(r.d.barWidth, -r.d.data[i]))
			bar.Move(fyne.NewPos(step*float32(i), s.Height/2))
		} else {
			bar.Resize(fyne.NewSize(r.d.barWidth, r.d.data[i]))
			bar.Move(fyne.NewPos(step*float32(i), s.Height/2-r.d.data[i]))
		}
	}

	r.d.X.Position1 = fyne.NewPos(0, s.Height/2)
	r.d.X.Position2 = fyne.NewPos(s.Width, s.Height/2)
	r.d.Y.Position1 = fyne.NewPos(s.Width/2, 0)
	r.d.Y.Position2 = fyne.NewPos(s.Width/2, s.Height)
}
func (r *barRenderer) MinSize() fyne.Size {
	return r.d.MinSize()
}

func (r *barRenderer) Refresh() {
	canvas.Refresh(r.d)
}
