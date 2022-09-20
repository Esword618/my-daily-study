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

var _ fyne.Widget = (*CircleButton)(nil)

// CircleButton is a widget for displaying a separator with themeable color.
//
// Since: 1.4
type CircleButton struct {
	widget.BaseWidget
	Txt1       *canvas.Text
	Background *canvas.Circle
	OnTapped   func(*CircleButton) `json:"-"`
	OnHoverIn  func()              `json:"-"`
	OnHoverOut func()              `json:"-"`
	Data       interface{}
}

// NewCircleButton creates a new separator.
//
// Since: 1.4
func NewCircleButton() *CircleButton {
	me := &CircleButton{
		Txt1:       canvas.NewText("txt1", theme.ForegroundColor()),
		Background: canvas.NewCircle(theme.DisabledColor()),
	}
	me.ExtendBaseWidget(me)
	me.Txt1.Alignment = fyne.TextAlignCenter
	me.Background.StrokeColor = colornames.Blue
	me.Background.StrokeWidth = 1
	return me
}

func (me *CircleButton) Do(fn func(t *CircleButton)) *CircleButton {
	fn(me)
	me.Refresh()
	return me
}
func (me *CircleButton) SText1(t string) *CircleButton {
	me.Txt1.Text = t
	return me
}
func (me *CircleButton) SData(d interface{}) *CircleButton {
	me.Data = d
	return me
}
func (me *CircleButton) TextColor1(c color.Color) *CircleButton {
	me.Txt1.Color = c
	return me
}

func (me *CircleButton) TextSize(size float32) *CircleButton {
	me.Txt1.TextSize = size
	return me
}

// TextAlign fyne.TextAlignLeading , fyne.TextAlignCenter , fyne.TextAlignTrailing
func (me *CircleButton) TextAlign(align fyne.TextAlign) *CircleButton {
	me.Txt1.Alignment = align
	return me
}

// Style
// fyne.TextStyle{Bold: true}
// fyne.TextStyle{Italic: true}
// fyne.TextStyle{Monospace: true}
// fyne.TextStyle{TabWidth: 10}
func (me *CircleButton) Style(style fyne.TextStyle) *CircleButton {
	me.Txt1.TextStyle = style
	return me
}
func (me *CircleButton) BgStrokeColor(c color.Color) *CircleButton {
	me.Background.StrokeColor = c
	return me
}
func (me *CircleButton) BgFillColor(c color.Color) *CircleButton {
	me.Background.FillColor = c
	return me
}

func (me *CircleButton) Click(fn func(*CircleButton)) *CircleButton {
	me.OnTapped = fn
	return me
}

// Cursor returns the cursor type of this widget
func (b *CircleButton) Cursor() desktop.Cursor {
	return desktop.PointerCursor
}

func (b *CircleButton) Tapped(*fyne.PointEvent) {
	if b.OnTapped != nil {
		b.OnTapped(b)
	}
}

// MouseIn is called when a desktop pointer enters the widget
func (b *CircleButton) MouseIn(*desktop.MouseEvent) {
	b.Background.StrokeColor = colornames.Red
	if b.OnHoverIn != nil {
		b.OnHoverIn()
	}
	b.Refresh()
}

// MouseMoved is called when a desktop pointer hovers over the widget
func (b *CircleButton) MouseMoved(*desktop.MouseEvent) {
}

// MouseOut is called when a desktop pointer exits the widget
func (b *CircleButton) MouseOut() {
	b.Background.StrokeColor = colornames.Blue
	if b.OnHoverOut != nil {
		b.OnHoverOut()
	}
	b.Refresh()
}

// CreateRenderer returns a new renderer for the separator.
//
// Implements: fyne.Widget
func (me *CircleButton) CreateRenderer() fyne.WidgetRenderer {
	me.ExtendBaseWidget(me)
	return &circleButtonRenderer{
		WidgetRenderer: widget.NewSimpleRenderer(container.NewMax(
			me.Background, me.Txt1)),
		d: me,
	}
}

// MinSize returns the minimal size of the separator.
//
// Implements: fyne.Widget
func (me *CircleButton) MinSize() fyne.Size {
	me.ExtendBaseWidget(me)
	ts := fyne.MeasureText(me.Txt1.Text, theme.TextSize(), fyne.TextStyle{})
	pd := theme.Padding() * 2
	max := fyne.Max(ts.Width, ts.Height)
	return fyne.NewSize(max, max).Add(fyne.NewDelta(pd, pd))
}

var _ fyne.WidgetRenderer = (*circleButtonRenderer)(nil)

type circleButtonRenderer struct {
	fyne.WidgetRenderer
	d *CircleButton
}

func (r *circleButtonRenderer) MinSize() fyne.Size {
	return r.d.MinSize()
}

func (r *circleButtonRenderer) Refresh() {
	canvas.Refresh(r.d)
}
