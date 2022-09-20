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
)

var _ fyne.Widget = (*Scatter)(nil)

// Scatter is a widget for displaying a separator with themeable color.
//
// Since: 1.4
type Scatter struct {
	widget.BaseWidget
	W, H      float32
	lineWidth float32
	data      []float32
	lineColor color.Color
	X         *canvas.Line
	Y         *canvas.Line
	T1        *canvas.Text

	Center   fyne.Position
	ChartCtn *fyne.Container

	txtCtn *fyne.Container
	xgrid  *fyne.Container
	ygrid  *fyne.Container
}

// NewScatter creates a new separator.
//
// Since: 1.4
func NewScatter() *Scatter {
	s := &Scatter{}
	s.ExtendBaseWidget(s)
	s.T1 = canvas.NewText("", colornames.Blue)
	s.X = canvas.NewLine(colornames.Blue)
	s.X.StrokeWidth = 3
	s.Y = canvas.NewLine(colornames.Red)
	s.Y.StrokeWidth = 3
	s.xgrid = container.NewMax()
	s.ygrid = container.NewMax()
	s.ChartCtn = container.NewMax()
	s.txtCtn = container.NewMax()
	for i := 0; i < 10; i++ {
		s.xgrid.Add(canvas.NewLine(color.Black))
		s.ygrid.Add(canvas.NewLine(color.Black))
	}
	return s
}

func (s *Scatter) Data(data ...float32) *Scatter {
	s.data = data
	s.ChartCtn.Objects = make([]fyne.CanvasObject, len(data))
	s.txtCtn.Objects = make([]fyne.CanvasObject, len(data))
	for i, v := range data {
		s.txtCtn.Objects[i] = canvas.NewText(fmt.Sprintf("%.2f", v),
			colornames.Red)
	}
	for i := range data {
		line := canvas.NewCircle(colornames.Chocolate)
		s.ChartCtn.Objects[i] = line
	}
	return s
}

func (s *Scatter) Do(fn func(*Scatter)) *Scatter {
	fn(s)
	return s
}
func (s *Scatter) ScatterWidth(w float32) *Scatter {
	s.lineWidth = w
	return s
}
func (s *Scatter) ScatterColor(c color.Color) *Scatter {
	s.lineColor = c
	return s
}
func (s *Scatter) Tapped(e *fyne.PointEvent) {

}
func (s *Scatter) MouseIn(*desktop.MouseEvent) {
}

// MouseMoved is called when a desktop pointer hovers over the widget
func (s *Scatter) MouseMoved(e *desktop.MouseEvent) {

}

// MouseOut is called when a desktop pointer exits the widget
func (s *Scatter) MouseOut() {
}

// CreateRenderer returns a new renderer for the separator.
//
// Implements: fyne.Widget
func (s *Scatter) CreateRenderer() fyne.WidgetRenderer {
	s.ExtendBaseWidget(s)
	ctn := container.NewMax(s.xgrid, s.ygrid, s.X, s.Y, s.T1,
		s.ChartCtn, s.txtCtn)
	return &scatterRenderer{
		WidgetRenderer: widget.NewSimpleRenderer(ctn),
		d:              s,
	}
}

// MinSize returns the minimal size of the separator.
//
// Implements: fyne.Widget
func (s *Scatter) MinSize() fyne.Size {
	s.ExtendBaseWidget(s)
	return fyne.NewSize(600, 600)
}

var _ fyne.WidgetRenderer = (*scatterRenderer)(nil)

type scatterRenderer struct {
	fyne.WidgetRenderer
	d *Scatter
}

func (r *scatterRenderer) Layout(s fyne.Size) {
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
		line := obj.(*canvas.Circle)
		line.StrokeWidth = r.d.lineWidth
		line.FillColor = r.d.lineColor
		line.Position1 = fyne.NewPos(step*float32(i)+10,
			s.Height/2-r.d.data[i])
		line.Position2 = line.Position1.Add(fyne.NewDelta(20, 20))
	}
	for i, obj := range r.d.txtCtn.Objects {
		txt := obj.(*canvas.Text)
		txt.Move(fyne.NewPos(step*float32(i)+10, s.Height/2-r.d.data[i]))
	}
	r.d.X.Position1 = fyne.NewPos(0, s.Height/2)
	r.d.X.Position2 = fyne.NewPos(s.Width, s.Height/2)
	r.d.Y.Position1 = fyne.NewPos(s.Width/2, 0)
	r.d.Y.Position2 = fyne.NewPos(s.Width/2, s.Height)
}
func (r *scatterRenderer) MinSize() fyne.Size {
	return r.d.MinSize()
}

func (r *scatterRenderer) Refresh() {
	canvas.Refresh(r.d)
}
