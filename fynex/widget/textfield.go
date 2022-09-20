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

var _ fyne.Widget = (*TextField)(nil)

// TextField is a widget for displaying a separator with themeable color.
//
// Since: 1.4
type TextField struct {
	widget.BaseWidget
	Txt1       *canvas.Text
	Background *canvas.Rectangle
	OnTapped   func(*TextField) `json:"-"`
	OnHoverIn  func()           `json:"-"`
	OnHoverOut func()           `json:"-"`
	Data       interface{}
}

// NewTextField creates a new separator.
//
// Since: 1.4
func NewTextField() *TextField {
	me := &TextField{
		Txt1:       canvas.NewText("txt1", theme.ForegroundColor()),
		Background: canvas.NewRectangle(theme.DisabledColor()),
	}
	me.ExtendBaseWidget(me)
	return me
}

func (me *TextField) Do(fn func(t *TextField)) *TextField {
	fn(me)
	me.Refresh()
	return me
}
func (me *TextField) SText1(t string) *TextField {
	me.Txt1.Text = t
	return me
}
func (me *TextField) SData(d interface{}) *TextField {
	me.Data = d
	return me
}
func (me *TextField) TextColor1(c color.Color) *TextField {
	me.Txt1.Color = c
	return me
}

func (me *TextField) TextSize(size float32) *TextField {
	me.Txt1.TextSize = size
	return me
}

// TextAlign fyne.TextAlignLeading , fyne.TextAlignCenter , fyne.TextAlignTrailing
func (me *TextField) TextAlign(align fyne.TextAlign) *TextField {
	me.Txt1.Alignment = align
	return me
}

// Style
// fyne.TextStyle{Bold: true}
// fyne.TextStyle{Italic: true}
// fyne.TextStyle{Monospace: true}
// fyne.TextStyle{TabWidth: 10}
func (me *TextField) Style(style fyne.TextStyle) *TextField {
	me.Txt1.TextStyle = style
	return me
}
func (me *TextField) BgStrokeColor(c color.Color) *TextField {
	me.Background.StrokeColor = c
	return me
}
func (me *TextField) BgFillColor(c color.Color) *TextField {
	me.Background.FillColor = c
	return me
}

func (me *TextField) Click(fn func(*TextField)) *TextField {
	me.OnTapped = fn
	return me
}

func (b *TextField) Tapped(*fyne.PointEvent) {
	if b.OnTapped != nil {
		b.OnTapped(b)
	}
}

// MouseIn is called when a desktop pointer enters the widget
func (b *TextField) MouseIn(*desktop.MouseEvent) {
	if b.OnHoverIn != nil {
		b.OnHoverIn()
	}
	b.Refresh()
}

// MouseMoved is called when a desktop pointer hovers over the widget
func (b *TextField) MouseMoved(*desktop.MouseEvent) {
}

// MouseOut is called when a desktop pointer exits the widget
func (b *TextField) MouseOut() {
	if b.OnHoverOut != nil {
		b.OnHoverOut()
	}
	b.Refresh()
}

// CreateRenderer returns a new renderer for the separator.
//
// Implements: fyne.Widget
func (me *TextField) CreateRenderer() fyne.WidgetRenderer {
	me.ExtendBaseWidget(me)
	return &textFieldRenderer{
		WidgetRenderer: widget.NewSimpleRenderer(container.NewMax(
			me.Background, me.Txt1)),
		d: me,
	}
}

// MinSize returns the minimal size of the separator.
//
// Implements: fyne.Widget
func (me *TextField) MinSize() fyne.Size {
	me.ExtendBaseWidget(me)

	pd := theme.Padding() * 2
	return me.Txt1.MinSize().Add(fyne.NewDelta(pd, pd))
}

var _ fyne.WidgetRenderer = (*textFieldRenderer)(nil)

type textFieldRenderer struct {
	fyne.WidgetRenderer
	d *TextField
}

func (r *textFieldRenderer) MinSize() fyne.Size {
	return r.d.MinSize()
}

func (r *textFieldRenderer) Refresh() {
	canvas.Refresh(r.d)
}
