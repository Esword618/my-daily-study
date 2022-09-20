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

var _ fyne.Widget = (*Bar)(nil)

// Bar is a widget for displaying a separator with themeable color.
//
// Since: 1.4
type Bar struct {
	widget.BaseWidget
	Txt1       *canvas.Text
	Txt2       *canvas.Text
	Bg         *canvas.Rectangle
	OnTapped   func(*Bar) `json:"-"`
	OnHoverIn  func()     `json:"-"`
	OnHoverOut func()     `json:"-"`
	Data       interface{}
}

// NewBar creates a new separator.
//
// Since: 1.4
func NewBar() *Bar {
	me := &Bar{
		Txt1: canvas.NewText("txt1", theme.ForegroundColor()),
		Txt2: canvas.NewText("txt2", theme.ForegroundColor()),
		Bg:   canvas.NewRectangle(theme.DisabledColor()),
	}
	me.ExtendBaseWidget(me)
	return me
}

func (me *Bar) Do(fn func(t *Bar)) *Bar {
	fn(me)
	me.Refresh()
	return me
}
func (me *Bar) SText1(t string) *Bar {
	me.Txt1.Text = t
	me.Txt1.Refresh()
	return me
}
func (me *Bar) SData(d interface{}) *Bar {
	me.Data = d
	return me
}
func (me *Bar) TextColor1(c color.Color) *Bar {
	me.Txt1.Color = c
	return me
}
func (me *Bar) TextColor2(c color.Color) *Bar {
	me.Txt2.Color = c
	return me
}
func (me *Bar) TextSize(size float32) *Bar {
	me.Txt1.TextSize = size
	return me
}

// TextAlign fyne.TextAlignLeading , fyne.TextAlignCenter , fyne.TextAlignTrailing
func (me *Bar) TextAlign(align fyne.TextAlign) *Bar {
	me.Txt1.Alignment = align
	return me
}

// Style
// fyne.TextStyle{Bold: true}
// fyne.TextStyle{Italic: true}
// fyne.TextStyle{Monospace: true}
// fyne.TextStyle{TabWidth: 10}
func (me *Bar) Style(style fyne.TextStyle) *Bar {
	me.Txt1.TextStyle = style
	return me
}
func (me *Bar) StrokeColor(c color.Color) *Bar {
	me.Bg.StrokeColor = c
	me.Bg.Refresh()
	return me
}
func (me *Bar) StrokeWidth(w float32) *Bar {
	me.Bg.StrokeWidth = w
	me.Bg.Refresh()
	return me
}

func (me *Bar) FillColor(c color.Color) *Bar {
	me.Bg.FillColor = c
	me.Bg.Refresh()
	return me
}

func (me *Bar) Click(fn func(*Bar)) *Bar {
	me.OnTapped = fn
	return me
}

func (b *Bar) Tapped(*fyne.PointEvent) {
	if b.OnTapped != nil {
		b.OnTapped(b)
	}
}

// MouseIn is called when a desktop pointer enters the widget
func (b *Bar) MouseIn(*desktop.MouseEvent) {
	b.StrokeColor(colornames.Red)
	if b.OnHoverIn != nil {
		b.OnHoverIn()
	}
	b.Refresh()
}

// MouseMoved is called when a desktop pointer hovers over the widget
func (b *Bar) MouseMoved(*desktop.MouseEvent) {
}

// MouseOut is called when a desktop pointer exits the widget
func (b *Bar) MouseOut() {
	b.StrokeColor(colornames.Blue)
	if b.OnHoverOut != nil {
		b.OnHoverOut()
	}
	b.Refresh()
}

// CreateRenderer returns a new renderer for the separator.
//
// Implements: fyne.Widget
func (me *Bar) CreateRenderer() fyne.WidgetRenderer {
	me.ExtendBaseWidget(me)
	return &barRenderer{
		WidgetRenderer: widget.NewSimpleRenderer(container.NewMax(
			me.Bg, me.Txt1, me.Txt2)),
		d: me,
	}
}

// MinSize returns the minimal size of the separator.
//
// Implements: fyne.Widget
func (me *Bar) MinSize() fyne.Size {
	me.ExtendBaseWidget(me)
	w := fyne.Max(me.Txt1.MinSize().Width, me.Txt2.MinSize().Width)
	h := me.Txt1.MinSize().Height + me.Txt2.MinSize().Height
	pd := theme.Padding() * 2
	return fyne.NewSize(w, h).Add(fyne.NewDelta(pd, pd))
}

var _ fyne.WidgetRenderer = (*barRenderer)(nil)

type barRenderer struct {
	fyne.WidgetRenderer
	d *Bar
}

func (r *barRenderer) Layout(s fyne.Size) {
	r.d.Txt1.Move(fyne.NewPos(0, 0))
	r.d.Txt2.Move(fyne.NewPos(0, s.Height))
	r.d.Bg.Resize(s)
}
func (r *barRenderer) MinSize() fyne.Size {
	return r.d.MinSize()
}

func (r *barRenderer) Refresh() {

	canvas.Refresh(r.d)
}
