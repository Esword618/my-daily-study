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

var _ fyne.Widget = (*PolyLine)(nil)

// PolyLine is a widget for displaying a separator with themeable color.
//
// Since: 1.4
type PolyLine struct {
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
	txtCtn   *fyne.Container

	xgrid *fyne.Container
	ygrid *fyne.Container
}

// NewPolyLine creates a new separator.
//
// Since: 1.4
func NewPolyLine() *PolyLine {
	s := &PolyLine{}
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

// axis := chart.NewPolyLine().Data([]float32{120, 20, 30, 40, 60, 180, 240})
func (s *PolyLine) Data(data ...float32) *PolyLine {
	s.data = data
	s.ChartCtn.Objects = make([]fyne.CanvasObject, len(data)-1)
	s.txtCtn.Objects = make([]fyne.CanvasObject, len(data))
	for i, v := range data {
		s.txtCtn.Objects[i] = canvas.NewText(fmt.Sprintf("%.2f", v),
			colornames.Red)
	}
	for i := 0; i < len(data)-1; i++ {
		line := canvas.NewLine(colornames.Chocolate)
		line.StrokeWidth = 4
		s.ChartCtn.Objects[i] = line
	}
	return s
}

func (s *PolyLine) Do(fn func(*PolyLine)) *PolyLine {
	fn(s)
	return s
}
func (s *PolyLine) LineWidth(w float32) *PolyLine {
	s.lineWidth = w
	return s
}
func (s *PolyLine) LineColor(c color.Color) *PolyLine {
	s.lineColor = c
	return s
}
func (s *PolyLine) Tapped(e *fyne.PointEvent) {

}
func (s *PolyLine) MouseIn(*desktop.MouseEvent) {
}

// MouseMoved is called when a desktop pointer hovers over the widget
func (s *PolyLine) MouseMoved(e *desktop.MouseEvent) {

}

// MouseOut is called when a desktop pointer exits the widget
func (s *PolyLine) MouseOut() {
}

// CreateRenderer returns a new renderer for the separator.
//
// Implements: fyne.Widget
func (s *PolyLine) CreateRenderer() fyne.WidgetRenderer {
	s.ExtendBaseWidget(s)
	ctn := container.NewMax(s.xgrid, s.ygrid, s.X, s.Y, s.T1,
		s.ChartCtn, s.txtCtn)
	return &polyLineRenderer{
		WidgetRenderer: widget.NewSimpleRenderer(ctn),
		d:              s,
	}
}

// MinSize returns the minimal size of the separator.
//
// Implements: fyne.Widget
func (s *PolyLine) MinSize() fyne.Size {
	s.ExtendBaseWidget(s)
	return fyne.NewSize(600, 600)
}

var _ fyne.WidgetRenderer = (*lineRenderer)(nil)

type polyLineRenderer struct {
	fyne.WidgetRenderer
	d *PolyLine
}

func (r *polyLineRenderer) Layout(s fyne.Size) {
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
		line := obj.(*canvas.Line)
		line.StrokeWidth = r.d.lineWidth
		line.StrokeColor = r.d.lineColor
		line.Position1 = fyne.NewPos(step*float32(i)+10, s.Height/2-r.d.data[i])
		line.Position2 = fyne.NewPos(step*float32(i+1)+10,
			s.Height/2-r.d.data[i+1])
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
func (r *polyLineRenderer) MinSize() fyne.Size {
	return r.d.MinSize()
}

func (r *polyLineRenderer) Refresh() {
	canvas.Refresh(r.d)
}
