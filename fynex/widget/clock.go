package widget

import (
	"image/color"
	"math"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/image/colornames"
)

var _ fyne.Widget = (*Clock)(nil)

// Clock is a widget for displaying a separator with themeable color.
//
// Since: 1.4
type Clock struct {
	widget.BaseWidget
	TxtTime    *canvas.Text
	Txt1       *canvas.Text
	Txt2       *canvas.Text
	Txt3       *canvas.Text
	Txt4       *canvas.Text
	Txt5       *canvas.Text
	Txt6       *canvas.Text
	Txt7       *canvas.Text
	Txt8       *canvas.Text
	Txt9       *canvas.Text
	Txt10      *canvas.Text
	Txt11      *canvas.Text
	Txt12      *canvas.Text
	Background *canvas.Circle
	Hour       *canvas.Line
	Minute     *canvas.Line
	Second     *canvas.Line
	OnTapped   func(*Clock) `json:"-"`
	OnHoverIn  func()       `json:"-"`
	OnHoverOut func()       `json:"-"`
	Data       interface{}
}

// NewClock creates a new separator.
// color reference https://www.58pic.com/peisebiao/
// Since: 1.4
func NewClock() *Clock {
	me := &Clock{
		TxtTime:    canvas.NewText("txt1", theme.ForegroundColor()),
		Txt1:       canvas.NewText("1", HexToRGBA("#FF9900")),
		Txt2:       canvas.NewText("2", HexToRGBA("#FF0033")),
		Txt3:       canvas.NewText("3", HexToRGBA("#0099CC")),
		Txt4:       canvas.NewText("4", HexToRGBA("#99CC33")),
		Txt5:       canvas.NewText("5", HexToRGBA("#FF9900")),
		Txt6:       canvas.NewText("6", HexToRGBA("#FFCC00")),
		Txt7:       canvas.NewText("7", HexToRGBA("#FF9933")),
		Txt8:       canvas.NewText("8", HexToRGBA("#99CC33")),
		Txt9:       canvas.NewText("9", HexToRGBA("#CC6699")),
		Txt10:      canvas.NewText("10", HexToRGBA("#FF9933")),
		Txt11:      canvas.NewText("11", HexToRGBA("#CC0033")),
		Txt12:      canvas.NewText("12", HexToRGBA("#3366CC")),
		Background: canvas.NewCircle(HexToRGBA("#FFFFCC")),
		Hour:       canvas.NewLine(colornames.Tan),
		Minute:     canvas.NewLine(colornames.Teal),
		Second:     canvas.NewLine(colornames.Red),
	}
	me.ExtendBaseWidget(me)
	me.Background.StrokeColor = colornames.Seagreen
	me.Background.StrokeWidth = 2
	me.Hour.StrokeWidth = 3
	me.Minute.StrokeWidth = 2
	me.Go()
	return me
}

func (t *Clock) Go() {
	go t.tickSecond()
}

func (t *Clock) tickSecond() {
	for {
		t.setPos()
		// t.refreshPos()
		t.Refresh()

		time.Sleep(time.Second)
	}
}

// func (t *Clock) tickSecond1() {
// 	for {
// 		for theta := -math.Pi / 2; theta <= math.Pi*1.5; theta += math.Pi / 180 {
// 			p1 := t.Background.Position1
// 			p2 := t.Background.Position2
// 			radius := (p2.X - p1.X) / 2
// 			t.Txt1.Text = time.Now().Format("2006-01-02 15:04:05")
// 			x := radius*float32(math.Cos(theta)) + radius
// 			y := radius*float32(math.Sin(theta)) + radius
// 			t.Second.Position2 = fyne.NewPos(x, y)
// 			t.Refresh()

// 			time.Sleep(time.Second / 1000)
// 		}
// 	}
// }

// func (d *Clock) refreshPos() {
// 	p1 := d.Background.Position1
// 	p2 := d.Background.Position2
// 	radius := (p2.X - p1.X) / 2
// 	now := time.Now()
// 	d.Txt1.Text = now.Format("2006-01-02 15:04:05")
// 	s := now.Second()
// 	thetaSecond := math.Pi*float64(s)/30 - math.Pi/2
// 	thetaMinute := thetaSecond / 60
// 	thetaHour := thetaSecond / 60
// 	sx := radius*0.9*float32(math.Cos(thetaSecond)) + radius
// 	sy := radius*0.9*float32(math.Sin(thetaSecond)) + radius
// 	d.Second.Position2 = fyne.NewPos(sx, sy)

// 	mx := radius*0.8*float32(math.Cos(thetaMinute)) + radius
// 	my := radius*0.8*float32(math.Sin(thetaMinute)) + radius
// 	d.Minute.Position2 = fyne.NewPos(mx, my)

// 	hx := radius*0.7*float32(math.Cos(thetaHour)) + radius
// 	hy := radius*0.7*float32(math.Sin(thetaHour)) + radius
// 	d.Hour.Position2 = fyne.NewPos(hx, hy)
// }

func (d *Clock) setPos() {
	p1 := d.Background.Position1
	p2 := d.Background.Position2
	radius := (p2.X - p1.X) / 2
	now := time.Now()
	d.TxtTime.Text = now.Format("2006-01-02 15:04:05")
	h := now.Hour()
	m := now.Minute()
	s := now.Second()
	thetaMinute := math.Pi*float64(m)/30 - math.Pi/2
	thetaHour := math.Pi*float64(h)/6 - 0.5*math.Pi+math.Pi*float64(m)/360
	thetaSecond := math.Pi*float64(s)/30 - math.Pi/2
	sx := radius*0.9*float32(math.Cos(thetaSecond)) + radius
	sy := radius*0.9*float32(math.Sin(thetaSecond)) + radius
	d.Second.Position2 = fyne.NewPos(sx, sy)

	mx := radius*0.8*float32(math.Cos(thetaMinute)) + radius
	my := radius*0.8*float32(math.Sin(thetaMinute)) + radius
	d.Minute.Position2 = fyne.NewPos(mx, my)

	hx := radius*0.7*float32(math.Cos(thetaHour)) + radius
	hy := radius*0.7*float32(math.Sin(thetaHour)) + radius
	d.Hour.Position2 = fyne.NewPos(hx, hy)
}
func (me *Clock) Do(fn func(t *Clock)) *Clock {
	fn(me)
	me.Refresh()
	return me
}
func (me *Clock) SText1(t string) *Clock {
	me.TxtTime.Text = t
	return me
}
func (me *Clock) SData(d interface{}) *Clock {
	me.Data = d
	return me
}
func (me *Clock) TextColor1(c color.Color) *Clock {
	me.TxtTime.Color = c
	return me
}

func (me *Clock) TextSize(size float32) *Clock {
	me.TxtTime.TextSize = size
	return me
}

// TextAlign fyne.TextAlignLeading , fyne.TextAlignCenter , fyne.TextAlignTrailing
func (me *Clock) TextAlign(align fyne.TextAlign) *Clock {
	me.TxtTime.Alignment = align
	return me
}

// Style
// fyne.TextStyle{Bold: true}
// fyne.TextStyle{Italic: true}
// fyne.TextStyle{Monospace: true}
// fyne.TextStyle{TabWidth: 10}
func (me *Clock) Style(style fyne.TextStyle) *Clock {
	me.TxtTime.TextStyle = style
	return me
}
func (me *Clock) BgStrokeColor(c color.Color) *Clock {
	me.Background.StrokeColor = c
	return me
}
func (me *Clock) BgFillColor(c color.Color) *Clock {
	me.Background.FillColor = c
	return me
}

func (me *Clock) Click(fn func(*Clock)) *Clock {
	me.OnTapped = fn
	return me
}

func (b *Clock) Tapped(*fyne.PointEvent) {
	if b.OnTapped != nil {
		b.OnTapped(b)
	}
}

// MouseIn is called when a desktop pointer enters the widget
func (b *Clock) MouseIn(*desktop.MouseEvent) {
	if b.OnHoverIn != nil {
		b.OnHoverIn()
	}
	b.Refresh()
}

// MouseMoved is called when a desktop pointer hovers over the widget
func (b *Clock) MouseMoved(*desktop.MouseEvent) {
}

// MouseOut is called when a desktop pointer exits the widget
func (b *Clock) MouseOut() {
	if b.OnHoverOut != nil {
		b.OnHoverOut()
	}
	b.Refresh()
}

// CreateRenderer returns a new renderer for the separator.
//
// Implements: fyne.Widget
func (me *Clock) CreateRenderer() fyne.WidgetRenderer {
	me.ExtendBaseWidget(me)

	return &clockRenderer{
		ctn: container.NewMax(
			me.Background, me.TxtTime, me.Hour, me.Minute, me.Second,
			me.Txt1, me.Txt2, me.Txt4, me.Txt5, me.Txt7,
			me.Txt8, me.Txt10, me.Txt11,
			me.Txt12,  me.Txt3, me.Txt6, me.Txt9),
		d: me,
	}
}

// MinSize returns the minimal size of the separator.
//
// Implements: fyne.Widget
func (me *Clock) MinSize() fyne.Size {
	me.ExtendBaseWidget(me)
	return fyne.NewSize(300, 300)
}

var _ fyne.WidgetRenderer = (*clockRenderer)(nil)

type clockRenderer struct {
	ctn *fyne.Container
	d   *Clock
}

func (r *clockRenderer) Objects() []fyne.CanvasObject {
	return r.ctn.Objects
}
func (r *clockRenderer) Layout(size fyne.Size) {
	r.d.Background.Position1 = fyne.NewPos(theme.Padding(), theme.Padding())
	min := fyne.Min(size.Width, size.Height)
	r.d.Background.Position2 = fyne.NewPos(min, min)

	p1 := r.d.Background.Position1
	p2 := r.d.Background.Position2
	radius := (p2.X - p1.X) / 2
	r.d.TxtTime.Move(fyne.NewPos(radius/2, radius/2))
	centerPos := fyne.NewPos(radius+theme.Padding(), radius+theme.Padding())

	r.d.Second.Position1 = centerPos
	r.d.Minute.Position1 = centerPos
	r.d.Hour.Position1 = centerPos

	r.d.Txt12.Move(fyne.NewPos(centerPos.X, p1.Y))
	r.d.Txt6.Move(fyne.NewPos(centerPos.X, p2.Y-r.d.Txt6.Size().Height))
	r.d.Txt3.Move(fyne.NewPos(centerPos.X+radius-r.d.Txt3.MinSize().Width, centerPos.Y))
	r.d.Txt9.Move(fyne.NewPos(p1.X, centerPos.Y))

	theta := math.Pi / 6
	r.d.Txt1.Move(fyne.NewPos(radius*float32(math.Cos(theta-math.Pi*0.5)),
		radius*float32(math.Sin(theta-math.Pi*0.5))).Add(fyne.NewDelta(radius, radius)))
	r.d.Txt2.Move(fyne.NewPos(radius*float32(math.Cos(theta*2-math.Pi*0.5)),
		radius*float32(math.Sin(theta*2-math.Pi*0.5))).Add(fyne.NewDelta(radius, radius)))
	r.d.Txt4.Move(fyne.NewPos(radius*float32(math.Cos(theta*4-math.Pi*0.5)),
		radius*float32(math.Sin(theta*4-math.Pi*0.5))).Add(fyne.NewDelta(radius*.95, radius)))
	r.d.Txt5.Move(fyne.NewPos(radius*float32(math.Cos(theta*5-math.Pi*0.5)),
		radius*float32(math.Sin(theta*5-math.Pi*0.5))).Add(fyne.NewDelta(radius*.95, radius*.95)))
	r.d.Txt7.Move(fyne.NewPos(radius*float32(math.Cos(theta*7-math.Pi*0.5)),
		radius*float32(math.Sin(theta*7-math.Pi*0.5))).Add(fyne.NewDelta(radius*1.05, radius*.95)))
	r.d.Txt8.Move(fyne.NewPos(radius*float32(math.Cos(theta*8-math.Pi*0.5)),
		radius*float32(math.Sin(theta*8-math.Pi*0.5))).Add(fyne.NewDelta(radius*1.05, radius)))
	r.d.Txt10.Move(fyne.NewPos(radius*float32(math.Cos(theta*10-math.Pi*0.5)),
		radius*float32(math.Sin(theta*10-math.Pi*0.5))).Add(fyne.NewDelta(radius, radius)))
	r.d.Txt11.Move(fyne.NewPos(radius*float32(math.Cos(theta*11-math.Pi*0.5)),
		radius*float32(math.Sin(theta*11-math.Pi*0.5))).Add(fyne.NewDelta(radius, radius)))

}
func (r *clockRenderer) Destroy() {

}
func (r *clockRenderer) MinSize() fyne.Size {
	return r.d.MinSize()
}

func (r *clockRenderer) Refresh() {

	canvas.Refresh(r.d)
}
