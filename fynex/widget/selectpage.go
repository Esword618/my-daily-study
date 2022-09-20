package widget

import (
	"image/color"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/image/colornames"
)

const defaultPlaceHolder string = "(Select one)"

// Select widget has a list of options, with the current one shown, and triggers an event func when clicked
type Select struct {
	widget.DisableableWidget

	// Alignment sets the text alignment of the select and its list of options.
	//
	// Since: 2.1
	Alignment   fyne.TextAlign
	Selected    string
	Options     []string
	PlaceHolder string
	numCol      int
	numRow      int
	curPage     int

	OnChanged func(string) `json:"-"`

	propertyLock sync.RWMutex
	disabled     bool
	focused      bool
	hovered      bool
	popUpContent fyne.CanvasObject
	items        []fyne.CanvasObject
	popUp        *widget.PopUp
	tapAnim      *fyne.Animation
}

var _ fyne.Widget = (*Select)(nil)
var _ desktop.Hoverable = (*Select)(nil)
var _ fyne.Tappable = (*Select)(nil)
var _ fyne.Focusable = (*Select)(nil)
var _ fyne.Disableable = (*Select)(nil)

// NewSelect creates a new select widget with the set list of options and changes handler
func NewSelect(options []string,
	changed func(string)) *Select {
	s := &Select{
		OnChanged:   changed,
		Options:     options,
		numCol:      1,
		numRow:      1000,
		curPage:     1,
		PlaceHolder: defaultPlaceHolder,
	}
	s.ExtendBaseWidget(s)

	s.items = make([]fyne.CanvasObject, len(s.Options))
	s.RefreshItem()
	return s
}

func (s *Select) Do(fn func(*Select)) *Select {
	fn(s)
	s.RefreshItem()
	return s
}
func (s *Select) Col(num int) *Select {
	s.numCol = num
	s.RefreshItem()
	return s
}

func (s *Select) Row(num int) *Select {
	s.numRow = num
	s.RefreshItem()
	return s
}
func (s *Select) CurPage(curPage int) *Select {
	s.curPage = curPage
	s.RefreshItem()
	return s
}

// ClearSelected clears the current option of the select widget.  After
// clearing the current option, the Select widget's PlaceHolder will
// be displayed.
func (s *Select) ClearSelected() {
	s.updateSelected("")
}

// CreateRenderer is a private method to Fyne which links this widget to its renderer
func (s *Select) CreateRenderer() fyne.WidgetRenderer {
	s.ExtendBaseWidget(s)
	s.propertyLock.RLock()
	icon := widget.NewIcon(theme.MenuDropDownIcon())
	if s.PlaceHolder == "" {
		s.PlaceHolder = defaultPlaceHolder
	}
	txtProv := widget.NewRichTextWithText(s.Selected)
	// txtProv.inset = fyne.NewSize(theme.Padding(), theme.Padding())
	txtProv.ExtendBaseWidget(txtProv)
	txtProv.Wrapping = fyne.TextTruncate
	if s.disabled {
		txtProv.Segments[0].(*widget.TextSegment).Style.ColorName = theme.ColorNameDisabled
	}

	background := &canvas.Rectangle{}
	line := canvas.NewRectangle(theme.ShadowColor())
	tapBG := canvas.NewRectangle(color.Transparent)
	s.tapAnim = newButtonTapAnimation(tapBG, s)
	s.tapAnim.Curve = fyne.AnimationEaseOut
	objects := []fyne.CanvasObject{background, line, tapBG, txtProv, icon}
	r := &selectRenderer{icon, txtProv, background, line, objects, s}
	background.FillColor, line.FillColor = r.bgLineColor()
	r.updateIcon()
	s.propertyLock.RUnlock() // updateLabel and some text handling isn't quite right, resolve in text refactor for 2.0
	r.updateLabel()
	return r
}

// FocusGained is called after this Select has gained focus.
//
// Implements: fyne.Focusable
func (s *Select) FocusGained() {
	s.focused = true
	s.Refresh()
}

// FocusLost is called after this Select has lost focus.
//
// Implements: fyne.Focusable
func (s *Select) FocusLost() {
	s.focused = false
	s.Refresh()
}

// Hide hides the select.
//
// Implements: fyne.Widget
func (s *Select) Hide() {
	if s.popUp != nil {
		s.popUp.Hide()
		s.popUp = nil
	}
	s.BaseWidget.Hide()
}

// MinSize returns the size that this widget should not shrink below
func (s *Select) MinSize() fyne.Size {
	s.ExtendBaseWidget(s)
	return s.BaseWidget.MinSize()
}

// MouseIn is called when a desktop pointer enters the widget
func (s *Select) MouseIn(*desktop.MouseEvent) {
	s.hovered = true
	s.Refresh()
}

// MouseMoved is called when a desktop pointer hovers over the widget
func (s *Select) MouseMoved(*desktop.MouseEvent) {
}

// MouseOut is called when a desktop pointer exits the widget
func (s *Select) MouseOut() {
	s.hovered = false
	s.Refresh()
}

// Move changes the relative position of the select.
//
// Implements: fyne.Widget
func (s *Select) Move(pos fyne.Position) {
	s.BaseWidget.Move(pos)

	if s.popUp != nil {
		s.popUp.Move(s.popUpPos())
	}
}

// Resize sets a new size for a widget.
// Note this should not be used if the widget is being managed by a Layout within a Container.
func (s *Select) Resize(size fyne.Size) {
	s.BaseWidget.Resize(size)

	if s.popUp != nil {
		s.popUp.Resize(fyne.NewSize(size.Width, s.popUp.MinSize().Height))
	}
}

// SelectedIndex returns the index value of the currently selected item in Options list.
// It will return -1 if there is no selection.
func (s *Select) SelectedIndex() int {
	for i, option := range s.Options {
		if s.Selected == option {
			return i
		}
	}
	return -1 // not selected/found
}

// SetSelected sets the current option of the select widget
func (s *Select) SetSelected(text string) {
	for _, option := range s.Options {
		if text == option {
			s.updateSelected(text)
		}
	}
}

// SetSelectedIndex will set the Selected option from the value in Options list at index position.
func (s *Select) SetSelectedIndex(index int) {
	if index < 0 || index >= len(s.Options) {
		return
	}

	s.updateSelected(s.Options[index])
}

// Tapped is called when a pointer tapped event is captured and triggers any tap handler
func (s *Select) Tapped(*fyne.PointEvent) {
	if s.Disabled() {
		return
	}

	s.tapAnimation()
	s.Refresh()

	s.showPopUp()
}

// TypedKey is called if a key event happens while this Select is focused.
//
// Implements: fyne.Focusable
func (s *Select) TypedKey(event *fyne.KeyEvent) {
	switch event.Name {
	case fyne.KeySpace, fyne.KeyUp, fyne.KeyDown:
		s.showPopUp()
	case fyne.KeyRight:
		i := s.SelectedIndex() + 1
		if i >= len(s.Options) {
			i = 0
		}
		s.SetSelectedIndex(i)
	case fyne.KeyLeft:
		i := s.SelectedIndex() - 1
		if i < 0 {
			i = len(s.Options) - 1
		}
		s.SetSelectedIndex(i)
	}
}

// TypedRune is called if a text event happens while this Select is focused.
//
// Implements: fyne.Focusable
func (s *Select) TypedRune(_ rune) {
	// intentionally left blank
}

func (s *Select) RefreshItem() {
	// items := make([]fyne.CanvasObject, len(s.Options))
	for i := range s.Options {
		text := s.Options[i] // capture
		s.items[i] = NewY2hText().Click(func(yt *Y2hText) {
			s.updateSelected(text)
			s.popUp.Hide()
			// s.popUp = nil
		}).Do(func(t *Y2hText) {
			t.Txt1.Text = text
			t.Txt2.Text = ""
			t.Txt1.Alignment = fyne.TextAlignCenter
			t.Txt2.Alignment = fyne.TextAlignCenter
			t.Background.StrokeColor = colornames.Yellow
			t.Background.StrokeWidth = 1
			t.OnHoverIn = func() {
				t.Background.StrokeWidth = 2
				t.Background.StrokeColor = colornames.Blue
				s.updateSelected(text)
			}
			t.OnHoverOut = func() {
				t.Background.StrokeColor = colornames.Yellow
				t.Background.StrokeWidth = 1
			}
		})
	}
	if s.numCol < 1 {
		s.numCol = 1
	}
	// 每次显示多少项
	numDisp := s.numCol * s.numRow
	// 选项总数
	lenOpt := len(s.Options)
	if s.curPage*numDisp > lenOpt || s.curPage < 1 {
		s.curPage = 1
	}

	if numDisp < lenOpt {
		// 总页数为n或者n+1
		n := lenOpt / numDisp
		// 默认第一页
		i := s.curPage
		// i := 1
		grid := container.NewGridWithColumns(s.numCol,
			s.items[(i-1)*numDisp:i*numDisp]...)
		btnPrev := widget.NewButton("Prev", func() {
			i = i - 1
			// fmt.Print("\ni: ", i, " n: ", n)
			if i < 1 {
				i = 1
			}
			grid.Objects = s.items[(i-1)*numDisp : i*numDisp]
		})
		btnNext := widget.NewButton("Next", func() {
			if i != n+1 {
				i = i + 1
				if lenOpt == (i-1)*numDisp {
					i = i - 1
				}
			}
			// fmt.Print("\ni: ", i, " n: ", n)
			if i > n {
				if lenOpt > (i-1)*numDisp {
					grid.Objects = s.items[(i-1)*numDisp:]
				}
				i = n + 1
				return
			}

			grid.Objects = s.items[(i-1)*numDisp : i*numDisp]
		})
		s.popUpContent = container.NewBorder(nil,
			container.NewGridWithColumns(2, btnPrev, btnNext), nil, nil, grid)

	} else {
		grid := container.NewGridWithColumns(s.numCol, s.items...)
		s.popUpContent = grid
	}
}
func (s *Select) popUpPos() fyne.Position {
	buttonPos := fyne.CurrentApp().Driver().AbsolutePositionForObject(s)
	return buttonPos.Add(fyne.NewPos(0, s.Size().Height-theme.InputBorderSize()))
}

func (s *Select) showPopUp() {
	if s.popUp == nil {
		c := fyne.CurrentApp().Driver().CanvasForObject(s)
		s.popUp = widget.NewPopUp(s.popUpContent, c)
	}
	s.popUp.ShowAtPosition(s.popUpPos())
	s.popUp.Resize(fyne.NewSize(s.Size().Width, s.popUp.MinSize().Height))
}

func (s *Select) tapAnimation() {
	if s.tapAnim == nil {
		return
	}
	s.tapAnim.Stop()
	s.tapAnim.Start()
}

func (s *Select) updateSelected(text string) {
	s.Selected = text

	if s.OnChanged != nil {
		s.OnChanged(s.Selected)
	}

	s.Refresh()
}

type selectRenderer struct {
	icon             *widget.Icon
	label            *widget.RichText
	background, line *canvas.Rectangle

	objects []fyne.CanvasObject
	combo   *Select
}

func (s *selectRenderer) Objects() []fyne.CanvasObject {
	return s.objects
}

func (s *selectRenderer) Destroy() {}

// Layout the components of the button widget
func (s *selectRenderer) Layout(size fyne.Size) {
	s.line.Resize(fyne.NewSize(size.Width, theme.InputBorderSize()))
	s.line.Move(fyne.NewPos(0, size.Height-theme.InputBorderSize()))
	s.background.Resize(fyne.NewSize(size.Width, size.Height-theme.InputBorderSize()*2))
	s.background.Move(fyne.NewPos(0, theme.InputBorderSize()))
	// s.label.inset = fyne.NewSize(theme.Padding(), theme.Padding())

	iconPos := fyne.NewPos(size.Width-theme.IconInlineSize()-theme.Padding()*2, (size.Height-theme.IconInlineSize())/2)
	labelSize := fyne.NewSize(iconPos.X-theme.Padding(), s.label.MinSize().Height)

	s.label.Resize(labelSize)
	s.label.Move(fyne.NewPos(theme.Padding(), (size.Height-labelSize.Height)/2))

	s.icon.Resize(fyne.NewSize(theme.IconInlineSize(), theme.IconInlineSize()))
	s.icon.Move(iconPos)
}

// MinSize calculates the minimum size of a select button.
// This is based on the selected text, the drop icon and a standard amount of padding added.
func (s *selectRenderer) MinSize() fyne.Size {
	s.combo.propertyLock.RLock()
	defer s.combo.propertyLock.RUnlock()

	minPlaceholderWidth := fyne.MeasureText(s.combo.PlaceHolder, theme.TextSize(), fyne.TextStyle{}).Width
	min := s.label.MinSize()
	min.Width = minPlaceholderWidth
	min = min.Add(fyne.NewSize(theme.Padding()*6, theme.Padding()*2))
	return min.Add(fyne.NewSize(theme.IconInlineSize()+theme.Padding()*2, 0))
}

func (s *selectRenderer) Refresh() {
	s.combo.propertyLock.RLock()
	s.updateLabel()
	s.updateIcon()
	s.background.FillColor, s.line.FillColor = s.bgLineColor()
	s.combo.propertyLock.RUnlock()

	s.Layout(s.combo.Size())
	if s.combo.popUp != nil {
		// s.combo.popUp.alignment = s.combo.Alignment
		s.combo.popUp.Move(s.combo.popUpPos())
		s.combo.popUp.Resize(fyne.NewSize(s.combo.MinSize().Width, s.combo.popUp.MinSize().Height))
		// s.combo.popUp.Resize(fyne.NewSize(s.combo.size.Width, s.combo.popUp.MinSize().Height))
		s.combo.popUp.Refresh()
	}
	s.background.Refresh()
	canvas.Refresh(s.combo)
}

func (s *selectRenderer) bgLineColor() (bg color.Color, line color.Color) {
	if s.combo.Disabled() {
		return theme.InputBackgroundColor(), theme.DisabledColor()
	}
	if s.combo.focused {
		return theme.FocusColor(), theme.PrimaryColor()
	}
	if s.combo.hovered {
		return theme.HoverColor(), theme.ShadowColor()
	}
	return theme.InputBackgroundColor(), theme.ShadowColor()
}

func (s *selectRenderer) updateIcon() {
	if s.combo.Disabled() {
		s.icon.Resource = theme.NewDisabledResource(theme.MenuDropDownIcon())
	} else {
		s.icon.Resource = theme.MenuDropDownIcon()
	}
	s.icon.Refresh()
}

func (s *selectRenderer) updateLabel() {
	if s.combo.PlaceHolder == "" {
		s.combo.PlaceHolder = defaultPlaceHolder
	}

	s.label.Segments[0].(*widget.TextSegment).Style.Alignment = s.combo.Alignment
	if s.combo.disabled {
		s.label.Segments[0].(*widget.TextSegment).Style.ColorName = theme.ColorNameDisabled
	} else {
		s.label.Segments[0].(*widget.TextSegment).Style.ColorName = theme.ColorNameForeground
	}
	if s.combo.Selected == "" {
		s.label.Segments[0].(*widget.TextSegment).Text = s.combo.PlaceHolder
	} else {
		s.label.Segments[0].(*widget.TextSegment).Text = s.combo.Selected
	}
	s.label.Refresh()
}
