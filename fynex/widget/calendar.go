package widget

import (
	"image/color"
	"strconv"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"gitee.com/y2h/goutil"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/liujiawm/gocalendar"
	"golang.org/x/image/colornames"
)

const defaultDate string = "(today)"

// Calendar widget has a list of options, with the current one shown, and triggers an event func when clicked
type Calendar struct {
	widget.DisableableWidget

	// Alignment sets the text alignment of the select and its list of options.
	//
	// Since: 2.1
	Alignment                 fyne.TextAlign
	curYear, curMonth, curDay int
	Selected                  string
	Options                   []string
	PlaceHolder               string

	OnChanged func(string) `json:"-"`

	propertyLock    sync.RWMutex
	disabled        bool
	focused         bool
	hovered         bool
	items           [42]*Y2hText
	DefaultCalendar *gocalendar.Calendar
	grid            *fyne.Container
	isChange        binding.Bool
	popUp           *widget.PopUp
	popUpContent    fyne.CanvasObject
	tapAnim         *fyne.Animation
}

var _ fyne.Widget = (*Calendar)(nil)
var _ desktop.Hoverable = (*Calendar)(nil)
var _ fyne.Tappable = (*Calendar)(nil)
var _ fyne.Focusable = (*Calendar)(nil)
var _ fyne.Disableable = (*Calendar)(nil)

// NewCalendar creates a new select widget with the set list of options and changes handler
func NewCalendar() *Calendar {
	s := &Calendar{
		DefaultCalendar: gocalendar.DefaultCalendar(),
		PlaceHolder:     defaultDate,
	}

	s.grid = container.NewGridWithColumns(7)
	for i := 0; i < 42; i++ {
		s.items[i] = NewY2hText()
		s.grid.Add(s.items[i])
	}
	s.today()
	s.ExtendBaseWidget(s)
	s.isChange = binding.NewBool()
	s.isChange.AddListener(binding.NewDataListener(func() {
		s.RefreshItem()
	}))
	s.SetItems()
	return s
}

func (s *Calendar) today() *Calendar {
	gt := gtime.Now()
	s.curDay = gt.Day()
	s.curYear = gt.Year()
	s.curMonth = gt.Month()
	return s
}
func (s *Calendar) RefreshItem() *Calendar {
	itemscal := s.DefaultCalendar.
		GenerateWithDate(s.curYear, s.curMonth, s.curDay)
	for i, item := range itemscal {
		item := item
		s.items[i].SData(item).
			Click(func(yt *Y2hText) {
				s.updateSelected(item.String())
				s.popUp.Hide()
				// s.popUp = nil
			}).Do(func(t *Y2hText) {
			it := t.Data.(*gocalendar.CalendarItem)
			t.OnHoverIn = func() {
				t.Background.StrokeWidth = 2
				t.Background.StrokeColor = colornames.Blue
				s.updateSelected(it.String())
			}
			t.OnHoverOut = func() {
				if it.IsToday == 1 {
					t.Background.StrokeColor = colornames.Blue
					t.Background.StrokeWidth = 3
				} else {
					t.Background.StrokeWidth = 1
					t.Background.StrokeColor = colornames.Yellow
				}
			}
			t.Background.StrokeWidth = 1
			t.Background.StrokeColor = colornames.Yellow
			t.Txt1.Alignment = fyne.TextAlignCenter
			t.Txt2.Alignment = fyne.TextAlignCenter
			t.Txt1.Text = strconv.Itoa(it.Time.Day())
			if it.SolarTerm != nil {
				t.Txt2.Text = it.SolarTerm.Name
				t.Background.FillColor = goutil.HexToRGBA("#FDE3E4")
			} else if len(it.Festival.Show) > 0 {
				t.Txt2.Text = it.Festival.Show[0]
				// t.Txt2.Text = item.Festival.Secondary[0]
				t.Background.FillColor = goutil.HexToRGBA("#FDE3E4")
			} else {
				t.Txt2.Text = it.LunarDate.DayName
				if it.IsToday == 1 {
					t.Background.StrokeColor = colornames.Blue
					t.Background.StrokeWidth = 3
				}
				if it.IsAccidental == -1 {
					t.Background.FillColor = colornames.Aliceblue
				} else if it.IsAccidental == 1 {
					t.Background.FillColor = colornames.Antiquewhite
				} else {
					t.Background.FillColor = colornames.Aqua
				}
			}
		}).Refresh()
	}
	return s
}
func (me *Calendar) SetItems() *Calendar {
	isChangev := false

	msg := widget.NewEntry()
	header := container.NewGridWithColumns(7)
	weekNameArray := [7]string{"日", "一", "二", "三", "四", "五", "六"}
	for _, v := range weekNameArray {
		header.Add(widget.NewLabelWithStyle(v, fyne.TextAlignCenter,
			fyne.TextStyle{Bold: true}))
	}
	selMonth := NewSelect(goutil.NumOpts(1, 13), nil).Col(4).Row(3).Do(func(s *Select) {
		s.OnChanged = func(ss string) {
			me.curMonth, _ = strconv.Atoi(ss)
			isChangev = !isChangev
			me.isChange.Set(isChangev)
		}
	})
	selYear := NewSelect(goutil.NumOpts(2000, 3001), nil).Col(7).Row(6).Do(func(s *Select) {
		s.OnChanged = func(ss string) {
			me.curYear, _ = strconv.Atoi(ss)
			isChangev = !isChangev
			me.isChange.Set(isChangev)
		}
	})

	lblToday := widget.NewLabel("")
	lblToday.SetText(gtime.Now().String())
	top := container.NewVBox(
		container.NewHBox(selYear,
			widget.NewButton("<", func() {
				if me.curMonth < 2 {
					me.curMonth = 1
				} else {
					me.curMonth = me.curMonth - 1
				}
				selMonth.updateSelected(strconv.Itoa(me.curMonth))
				// isChangev = !isChangev
				// me.isChange.Set(isChangev)
			}),
			selMonth,
			widget.NewButton(">", func() {
				if me.curMonth > 11 {
					me.curMonth = 12
				} else {
					me.curMonth = me.curMonth + 1
				}
				selMonth.updateSelected(strconv.Itoa(me.curMonth))
				// isChangev = !isChangev
				// me.isChange.Set(isChangev)
			}),
			widget.NewButton("today", func() {
				me.today()
				// me.RefreshItem()
				selMonth.SetSelected(strconv.Itoa(me.curMonth))
				selYear.SetSelected(strconv.Itoa(me.curYear))
				// me.SetSelected(me.curCalitem.String())
				lblToday.SetText(gtime.Now().String())
			}),
			lblToday,
		),
		header,
	)
	me.popUpContent = (container.NewBorder(top,
		msg, nil, nil, me.grid))
	return me
}

// ClearSelected clears the current option of the select widget.  After
// clearing the current option, the Calendar widget's PlaceHolder will
// be displayed.
func (s *Calendar) ClearSelected() {
	s.updateSelected("")
}

// CreateRenderer is a private method to Fyne which links this widget to its renderer
func (s *Calendar) CreateRenderer() fyne.WidgetRenderer {
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
	r := &calendarRenderer{icon, txtProv, background, line, objects, s}
	background.FillColor, line.FillColor = r.bgLineColor()
	r.updateIcon()
	s.propertyLock.RUnlock() // updateLabel and some text handling isn't quite right, resolve in text refactor for 2.0
	r.updateLabel()
	return r
}

// FocusGained is called after this Calendar has gained focus.
//
// Implements: fyne.Focusable
func (s *Calendar) FocusGained() {
	s.focused = true
	s.Refresh()
}

// FocusLost is called after this Calendar has lost focus.
//
// Implements: fyne.Focusable
func (s *Calendar) FocusLost() {
	s.focused = false
	s.Refresh()
}

// Hide hides the select.
//
// Implements: fyne.Widget
func (s *Calendar) Hide() {
	if s.popUp != nil {
		s.popUp.Hide()
		s.popUp = nil
	}
	s.BaseWidget.Hide()
}

// MinSize returns the size that this widget should not shrink below
func (s *Calendar) MinSize() fyne.Size {
	s.ExtendBaseWidget(s)
	return s.BaseWidget.MinSize()
}

// MouseIn is called when a desktop pointer enters the widget
func (s *Calendar) MouseIn(*desktop.MouseEvent) {
	s.hovered = true
	s.Refresh()
}

// MouseMoved is called when a desktop pointer hovers over the widget
func (s *Calendar) MouseMoved(*desktop.MouseEvent) {
}

// MouseOut is called when a desktop pointer exits the widget
func (s *Calendar) MouseOut() {
	s.hovered = false

	s.Refresh()
}

// Move changes the relative position of the select.
//
// Implements: fyne.Widget
func (s *Calendar) Move(pos fyne.Position) {
	s.BaseWidget.Move(pos)

	if s.popUp != nil {
		s.popUp.Move(s.popUpPos())
	}
}

// Resize sets a new size for a widget.
// Note this should not be used if the widget is being managed by a Layout within a Container.
func (s *Calendar) Resize(size fyne.Size) {
	s.BaseWidget.Resize(size)

	if s.popUp != nil {
		s.popUp.Resize(fyne.NewSize(size.Width, s.popUp.MinSize().Height))
	}
}

// SelectedIndex returns the index value of the currently selected item in Options list.
// It will return -1 if there is no selection.
func (s *Calendar) SelectedIndex() int {
	for i, option := range s.Options {
		if s.Selected == option {
			return i
		}
	}
	return -1 // not selected/found
}

// SetSelected sets the current option of the select widget
func (s *Calendar) SetSelected(text string) {
	for _, option := range s.Options {
		if text == option {
			s.updateSelected(text)
		}
	}
}

// SetSelectedIndex will set the Selected option from the value in Options list at index position.
func (s *Calendar) SetSelectedIndex(index int) {
	if index < 0 || index >= len(s.Options) {
		return
	}

	s.updateSelected(s.Options[index])
}

// Tapped is called when a pointer tapped event is captured and triggers any tap handler
func (s *Calendar) Tapped(*fyne.PointEvent) {
	if s.Disabled() {
		return
	}

	s.tapAnimation()
	s.Refresh()

	s.showPopUp()
}

// TypedKey is called if a key event happens while this Calendar is focused.
//
// Implements: fyne.Focusable
func (s *Calendar) TypedKey(event *fyne.KeyEvent) {
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

// TypedRune is called if a text event happens while this Calendar is focused.
//
// Implements: fyne.Focusable
func (s *Calendar) TypedRune(_ rune) {
	// intentionally left blank
}

func (s *Calendar) popUpPos() fyne.Position {
	buttonPos := fyne.CurrentApp().Driver().AbsolutePositionForObject(s)
	return buttonPos.Add(fyne.NewPos(0, s.Size().Height-theme.InputBorderSize()))
}

func (s *Calendar) showPopUp() {
	if s.popUp == nil {
		c := fyne.CurrentApp().Driver().CanvasForObject(s)
		s.popUp = widget.NewPopUp(s.popUpContent, c)
	}

	s.popUp.ShowAtPosition(s.popUpPos())
	s.popUp.Resize(fyne.NewSize(s.Size().Width, s.popUp.MinSize().Height))
}

func (s *Calendar) tapAnimation() {
	if s.tapAnim == nil {
		return
	}
	s.tapAnim.Stop()
	s.tapAnim.Start()
}

func (s *Calendar) updateSelected(text string) {
	s.Selected = text

	if s.OnChanged != nil {
		s.OnChanged(s.Selected)
	}

	s.Refresh()
}

type calendarRenderer struct {
	icon             *widget.Icon
	label            *widget.RichText
	background, line *canvas.Rectangle

	objects []fyne.CanvasObject
	combo   *Calendar
}

func (s *calendarRenderer) Objects() []fyne.CanvasObject {
	return s.objects
}

func (s *calendarRenderer) Destroy() {}

// Layout the components of the button widget
func (s *calendarRenderer) Layout(size fyne.Size) {
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
func (s *calendarRenderer) MinSize() fyne.Size {
	s.combo.propertyLock.RLock()
	defer s.combo.propertyLock.RUnlock()

	minPlaceholderWidth := fyne.MeasureText(s.combo.PlaceHolder, theme.TextSize(), fyne.TextStyle{}).Width
	min := s.label.MinSize()
	min.Width = minPlaceholderWidth
	min = min.Add(fyne.NewSize(theme.Padding()*6, theme.Padding()*2))
	return min.Add(fyne.NewSize(theme.IconInlineSize()+theme.Padding()*2, 0))
}

func (s *calendarRenderer) Refresh() {
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

func (s *calendarRenderer) bgLineColor() (bg color.Color, line color.Color) {
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

func (s *calendarRenderer) updateIcon() {
	if s.combo.Disabled() {
		s.icon.Resource = theme.NewDisabledResource(theme.MenuDropDownIcon())
	} else {
		s.icon.Resource = theme.MenuDropDownIcon()
	}
	s.icon.Refresh()
}

func (s *calendarRenderer) updateLabel() {
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
