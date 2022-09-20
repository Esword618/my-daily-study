package widget

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/image/colornames"
)

var _ fyne.Widget = (*Axis)(nil)

// Axis is a widget for displaying a separator with themeable color.
//
// Since: 1.4
type Axis struct {
	widget.BaseWidget
	isMoveDone bool
	W, H       float32
	Data       []float32
	X          *canvas.Line
	Y          *canvas.Line
	Line1      *canvas.Line
	P1         *canvas.Circle
	T1         *canvas.Text
	Points     []fyne.Position
	MousePos   fyne.Position
	Line1P1    fyne.Position
	Line1P2    fyne.Position
	Center     fyne.Position
	ChartCtn   *fyne.Container

	xgrid *fyne.Container
	ygrid *fyne.Container
}

// NewAxis creates a new separator.
//
// Since: 1.4
func NewAxis() *Axis {
	s := &Axis{}
	s.ExtendBaseWidget(s)
	s.X = canvas.NewLine(colornames.Blue)
	s.X.StrokeWidth = 3
	s.Y = canvas.NewLine(colornames.Red)
	s.Y.StrokeWidth = 3
	s.Line1 = canvas.NewLine(colornames.Burlywood)
	s.Line1.StrokeWidth = 2
	s.P1 = canvas.NewCircle(nil)
	s.P1.StrokeWidth = 2
	s.P1.StrokeColor = colornames.Red
	s.T1 = canvas.NewText("", colornames.Red)
	s.xgrid = container.NewMax()
	s.ygrid = container.NewMax()
	s.ChartCtn = container.NewWithoutLayout()
	for i := 0; i < 10; i++ {
		s.xgrid.Add(canvas.NewLine(color.Black))
		s.ygrid.Add(canvas.NewLine(color.Black))
	}
	return s
}

// axis := wdg.NewAxis().Lines([]float32{120, 20, 30, 40, 60, 180, 240})
// func (s *Axis) Lines(data []float32) *Axis {
// 	s.Data = data
// 	s.ChartCtn.Objects = make([]fyne.CanvasObject, len(data))
// 	for i := range data {
// 		line := canvas.NewLine(colornames.Chocolate)
// 		line.StrokeWidth = 4
// 		s.ChartCtn.Objects[i] = line
// 	}
// 	return s
// }
func (s *Axis) Do(fn func(*Axis)) *Axis {
	fn(s)
	return s
}
func (s *Axis) Tapped(e *fyne.PointEvent) {
	// s.isMoveDone = !s.isMoveDone
	// s.Line1P2 = e.Position
	// s.Line1.Position2 = e.Position
	// s.Line1.Refresh()
	s.Points = append(s.Points, e.Position)
	line := canvas.NewLine(colornames.Chocolate)
	line.StrokeWidth = 4
	s.ChartCtn.Add(line)
	// s.Refresh()
}
func (s *Axis) MouseIn(*desktop.MouseEvent) {
}

var r1 float32 = 10

// MouseMoved is called when a desktop pointer hovers over the widget
func (s *Axis) MouseMoved(e *desktop.MouseEvent) {
	// s.P1.Position1 = e.Position.Subtract(fyne.NewDelta(r1, r1))
	// s.P1.Position2 = e.Position.Add(fyne.NewDelta(r1, r1))
	// s.MousePos = e.Position.
	// 	Subtract(fyne.NewDelta(s.Size().Width/2,
	// 		s.Size().Height/2))
	// s.MousePos.Y = -s.MousePos.Y
	// s.T1.Text = fmt.Sprint(s.MousePos)
	// if !s.isMoveDone {
	// 	s.Line1.Position2 = e.Position
	// 	s.Line1.Refresh()
	// }
	// s.T1.Refresh()
	// s.P1.Refresh()
	// time.Sleep(time.Millisecond * 5)
}

// MouseOut is called when a desktop pointer exits the widget
func (s *Axis) MouseOut() {
}

// CreateRenderer returns a new renderer for the separator.
//
// Implements: fyne.Widget
func (s *Axis) CreateRenderer() fyne.WidgetRenderer {
	s.ExtendBaseWidget(s)
	bar := canvas.NewRectangle(theme.DisabledColor())
	ctn := container.NewMax(s.xgrid, s.ygrid, bar, s.X, s.Y,
		s.P1, s.T1, s.Line1, s.ChartCtn)
	return &axisRenderer{
		WidgetRenderer: widget.NewSimpleRenderer(ctn),
		bar:            bar,
		d:              s,
	}
}

// MinSize returns the minimal size of the separator.
//
// Implements: fyne.Widget
func (s *Axis) MinSize() fyne.Size {
	s.ExtendBaseWidget(s)
	return fyne.NewSize(600, 600)
}

var _ fyne.WidgetRenderer = (*axisRenderer)(nil)

type axisRenderer struct {
	fyne.WidgetRenderer
	bar *canvas.Rectangle
	d   *Axis
}

func (r *axisRenderer) Layout(s fyne.Size) {
	r.d.W = s.Width
	r.d.H = s.Height
	r.d.Center = fyne.NewPos(s.Width/2, s.Height/2)
	r.d.Line1.Position1 = fyne.NewPos(r.d.Size().Width/2,
		r.d.Size().Height/2)

	r.d.Line1.Position2 = r.d.Line1P2
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

	// step := r.d.W / float32(len(r.d.Data))
	// for i, obj := range r.d.ChartCtn.Objects {
	// 	line := obj.(*canvas.Line)
	// 	line.Position1 = fyne.NewPos(step*float32(i), s.Height/2)
	// 	line.Position2 = fyne.NewPos(step*float32(i),
	// 		s.Height/2-r.d.Data[i])
	// }

	for i, obj := range r.d.ChartCtn.Objects {
		line := obj.(*canvas.Line)
		line.Position1 = r.d.Points[i]
		if i+1 < len(r.d.Points) {
			line.Position2 = r.d.Points[i+1]
		} else {
			line.Position2 = r.d.Points[0]
		}
	}

	r.d.X.Position1 = fyne.NewPos(0, s.Height/2)
	r.d.X.Position2 = fyne.NewPos(s.Width, s.Height/2)
	r.d.Y.Position1 = fyne.NewPos(s.Width/2, 0)
	r.d.Y.Position2 = fyne.NewPos(s.Width/2, s.Height)
}
func (r *axisRenderer) MinSize() fyne.Size {
	return r.d.MinSize()
}

func (r *axisRenderer) Refresh() {
	r.bar.FillColor = theme.DisabledColor()
	r.Layout(r.d.Size())
	canvas.Refresh(r.d)
}
