package widget

import (
	"image"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	osdlg "github.com/sqweek/dialog"
)

var _ fyne.Widget = (*ClickImage)(nil)

// ClickImage is a widget for displaying a separator with themeable color.
//
// Since: 1.4
type ClickImage struct {
	widget.BaseWidget
	Img         *canvas.Image
	PlaceHolder *canvas.Text
	OnTapped    func(*ClickImage) `json:"-"`
	OnTapped2   func(me *ClickImage,e *fyne.PointEvent) `json:"-"`
}

// NewClickImage creates a new separator.
//
// Since: 1.4
func NewClickImage() *ClickImage {
	s := &ClickImage{}
	s.ExtendBaseWidget(s)
	s.Img = canvas.NewImageFromImage(nil)
	s.PlaceHolder = canvas.NewText("Click here to choose an image file",
		theme.PrimaryColor())
	s.PlaceHolder.Alignment = fyne.TextAlignCenter
	return s
}
func (s *ClickImage) SetImage(img image.Image) *ClickImage {
	s.Img.Image = img
	s.Refresh()
	return s
}
func (s *ClickImage) Do(fn func(t *ClickImage)) *ClickImage {
	fn(s)
	return s
}

func (b *ClickImage) TappedSecondary(e *fyne.PointEvent) {
	if b.OnTapped2 != nil {
		b.OnTapped2(b,e)
	}
}

// 打开文件选择对话框，选择需要显示的图片
// 使用系统原生对话框 https://hub.fastgit.xyz/ncruces/zenity
func (b *ClickImage) Tapped(*fyne.PointEvent) {
	if b.OnTapped != nil {
		b.OnTapped(b)
	}
}

// 打开文件选择对话框，选择需要显示的图片
// 使用系统原生对话框 "github.com/sqweek/dialog"
func (b *ClickImage) Tapped1(*fyne.PointEvent) {
	imgpath, _ := osdlg.File().Filter("image", ".pg", "jpeg", "png").
		Title("Select an image").
		SetStartDir(".").Load()
	b.Img.File = imgpath
	b.Img.Refresh()
	b.PlaceHolder.Hide()
	if b.OnTapped != nil {
		b.OnTapped(b)
	}
}

// 打开文件选择对话框，选择需要显示的图片
// fyne框架自带对话框
// func (b *ClickImage) Tapped(*fyne.PointEvent) {
// 	d := dialog.NewFileOpen(func(uc fyne.URIReadCloser, err error) {
// 		if uc == nil {
// 			return
// 		}
// 		imgpath := (uc.URI().String()[7:])
// 		b.Img.File=imgpath
// 		input, err := os.Open(imgpath)
// 		if err != nil {
// 			// log.Fatalf("error opening the source file: %v", err)
// 		}
// 		img, _ := decodeImage(input)
// 		b.SetImage(img)

// 	}, fyne.CurrentApp().Driver().AllWindows()[0])
// 	d.SetFilter(storage.NewExtensionFileFilter(
// 		[]string{".jpg", ".jpeg", ".png"}))
// 	luri, _ := storage.ListerForURI(storage.NewFileURI("."))
// 	d.SetLocation(luri)
// 	d.Show()

// }

// Cursor returns the cursor type of this widget
func (b *ClickImage) Cursor() desktop.Cursor {
	return desktop.PointerCursor
}

// CreateRenderer returns a new renderer for the separator.
//
// Implements: fyne.Widget
func (s *ClickImage) CreateRenderer() fyne.WidgetRenderer {
	s.ExtendBaseWidget(s)
	return &clickImageRenderer{
		WidgetRenderer: widget.NewSimpleRenderer(container.NewMax(s.Img,
			s.PlaceHolder)),
		d: s,
	}
}

// MinSize returns the minimal size of the separator.
//
// Implements: fyne.Widget
func (s *ClickImage) MinSize() fyne.Size {
	s.ExtendBaseWidget(s)
	// t := theme.ClickImageThicknessSize()
	// return fyne.NewSize(t, t)
	return s.Img.MinSize()
}

var _ fyne.WidgetRenderer = (*clickImageRenderer)(nil)

type clickImageRenderer struct {
	fyne.WidgetRenderer
	d *ClickImage
}

func (r *clickImageRenderer) MinSize() fyne.Size {
	return r.d.MinSize()
}

func (r *clickImageRenderer) Refresh() {
	// r.bar.FillColor = theme.DisabledColor()
	if r.d.Img.Image == nil {
		r.d.PlaceHolder.Show()
		// r.d.PlaceHolder.Text = "Click here to choose an image file"
	} else {
		r.d.PlaceHolder.Hide()
		// r.d.PlaceHolder.Text = r.d.Img.File
	}
	canvas.Refresh(r.d)
}
