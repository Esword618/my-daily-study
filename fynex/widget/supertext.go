package widget

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var _ fyne.Widget = (*Y2hText)(nil)

// Y2hText is a widget for displaying a separator with themeable color.
//
// Since: 1.4
type Y2hText struct {
	widget.BaseWidget
	Txt1       *canvas.Text
	Txt2       *canvas.Text
	Background *canvas.Rectangle
	OnTapped   func(*Y2hText) `json:"-"`
	OnHoverIn  func()         `json:"-"`
	OnHoverOut func()         `json:"-"`
	Data       interface{}
}

// NewY2hText creates a new separator.
//
// Since: 1.4
func NewY2hText() *Y2hText {
	me := &Y2hText{
		Txt1:       canvas.NewText("txt1", theme.ForegroundColor()),
		Txt2:       canvas.NewText("txt2", theme.ForegroundColor()),
		Background: canvas.NewRectangle(theme.DisabledColor()),
	}
	me.ExtendBaseWidget(me)
	return me
}

func (me *Y2hText) Do(fn func(t *Y2hText)) *Y2hText {
	fn(me)
	me.Refresh()
	return me
}
func (me *Y2hText) SText1(t string) *Y2hText {
	me.Txt1.Text = t
	return me
}
func (me *Y2hText) SData(d interface{}) *Y2hText {
	me.Data = d
	return me
}
func (me *Y2hText) TextColor1(c color.Color) *Y2hText {
	me.Txt1.Color = c
	return me
}
func (me *Y2hText) TextColor2(c color.Color) *Y2hText {
	me.Txt2.Color = c
	return me
}
func (me *Y2hText) TextSize(size float32) *Y2hText {
	me.Txt1.TextSize = size
	return me
}

// TextAlign fyne.TextAlignLeading , fyne.TextAlignCenter , fyne.TextAlignTrailing
func (me *Y2hText) TextAlign(align fyne.TextAlign) *Y2hText {
	me.Txt1.Alignment = align
	return me
}

// Style
// fyne.TextStyle{Bold: true}
// fyne.TextStyle{Italic: true}
// fyne.TextStyle{Monospace: true}
// fyne.TextStyle{TabWidth: 10}
func (me *Y2hText) Style(style fyne.TextStyle) *Y2hText {
	me.Txt1.TextStyle = style
	return me
}
func (me *Y2hText) BgStrokeColor(c color.Color) *Y2hText {
	me.Background.StrokeColor = c
	return me
}
func (me *Y2hText) BgFillColor(c color.Color) *Y2hText {
	me.Background.FillColor = c
	return me
}

func (me *Y2hText) Click(fn func(*Y2hText)) *Y2hText {
	me.OnTapped = fn
	return me
}

func (b *Y2hText) Tapped(*fyne.PointEvent) {
	if b.OnTapped != nil {
		b.OnTapped(b)
	}
}

// MouseIn is called when a desktop pointer enters the widget
func (b *Y2hText) MouseIn(*desktop.MouseEvent) {
	if b.OnHoverIn != nil {
		b.OnHoverIn()
	}
	b.Refresh()
}

// MouseMoved is called when a desktop pointer hovers over the widget
func (b *Y2hText) MouseMoved(*desktop.MouseEvent) {
}

// MouseOut is called when a desktop pointer exits the widget
func (b *Y2hText) MouseOut() {
	if b.OnHoverOut != nil {
		b.OnHoverOut()
	}
	b.Refresh()
}

// CreateRenderer returns a new renderer for the separator.
//
// Implements: fyne.Widget
func (me *Y2hText) CreateRenderer() fyne.WidgetRenderer {
	me.ExtendBaseWidget(me)
	return &separatorRenderer{
		WidgetRenderer: widget.NewSimpleRenderer(container.NewMax(
			me.Background, container.NewVBox(me.Txt1, me.Txt2))),
		d: me,
	}
}

// MinSize returns the minimal size of the separator.
//
// Implements: fyne.Widget
func (me *Y2hText) MinSize() fyne.Size {
	me.ExtendBaseWidget(me)
	w := fyne.Max(me.Txt1.MinSize().Width, me.Txt2.MinSize().Width)
	h := me.Txt1.MinSize().Height + me.Txt2.MinSize().Height
	pd := theme.Padding() * 2
	return fyne.NewSize(w, h).Add(fyne.NewDelta(pd, pd))
}

var _ fyne.WidgetRenderer = (*separatorRenderer)(nil)

type separatorRenderer struct {
	fyne.WidgetRenderer
	d *Y2hText
}

func (r *separatorRenderer) MinSize() fyne.Size {
	return r.d.MinSize()
}

func (r *separatorRenderer) Refresh() {
	if r.d.Txt2.Text == "" {
		r.d.Txt1.Move(fyne.NewPos(0, 0).
			Add(fyne.NewDelta(0, r.d.Size().Height/2-r.d.Txt1.TextSize/2)))
	} else {
		r.d.Txt1.Move(fyne.NewPos(0, 0).Add(fyne.NewDelta(0, theme.Padding())))
	}

	canvas.Refresh(r.d)
}
