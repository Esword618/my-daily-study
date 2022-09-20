package main

import (
	"image"
	"io"
	"log"
	"net/url"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/cmd/fyne_settings/settings"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"gitee.com/y2h/fynex/common"
	"gitee.com/y2h/fynex/component"
	"gitee.com/y2h/fynex/datastruct"
	"gitee.com/y2h/fynex/tool"
	"gitee.com/y2h/fynex/view"
	xwdg "gitee.com/y2h/fynex/widget"
	"gitee.com/y2h/fynex/widget/golocation"
	"github.com/ncruces/zenity"
)

func main() {
	a := app.NewWithID("io.fyne.demo")
	a.Settings().SetTheme(&myTheme{})
	w := a.NewWindow("gitee.com/y2h/fynex")

	opts := []string{"Tools", "SuperText", "Calendar", "CnRegion", "GoLocation",
		"DataStruct", "SudokGo", "StarLight", "Clock", "MandelbrotUI",
		"ClickImage", "Ease", "Slider", "Shuerte", "Current", "Clock1", "FileTree",
		"LineView", "CircleButton", "Axis", "Zenity", "Chart"}
	mapp := make(map[string]fyne.CanvasObject, len(opts))
	mapp[opts[0]] = tool.Tool()
	mapp[opts[1]] = supertext()
	mapp[opts[2]] = calendar()
	mapp[opts[3]] = xwdg.Cnregion()
	mapp[opts[4]] = golocation.UI()
	mapp[opts[5]] = datastruct.DataStruct()
	mapp[opts[6]] = component.SudokGo(w)
	mapp[opts[7]] = component.StarLight()
	mapp[opts[8]] = xwdg.NewClock()
	mapp[opts[9]] = mandelbrotUI()
	mapp[opts[10]] = xwdg.NewClickImage().Do(func(b *xwdg.ClickImage) {
		b.OnTapped = func(ci *xwdg.ClickImage) {
			imgpath, _ := zenity.SelectFile(
				zenity.FileFilters{
					{"Image files", []string{"*.png", "*.gif", "*.ico", "*.jpg", "*.webp"}},
				},
				zenity.Title("Please choose an image"),
			)
			if imgpath == "" {
				return
			}
			b.Img.File = imgpath
			img, _ := os.Open(imgpath)
			b.Img.Image, _ = decodeImage(img)
			b.Img.Refresh()
			b.PlaceHolder.Hide()
		}
		b.OnTapped2 = func(ci *xwdg.ClickImage,e *fyne.PointEvent) {
			d := fyne.CurrentApp().Driver()

			pop := widget.NewPopUp(xwdg.NewTextField().Do(func(t *xwdg.TextField) {
				t.Txt1.Text = "Open width Default app"
				t.OnTapped = func(tf *xwdg.TextField) {
					if common.PathExists(b.Img.File) {
						common.OpenWithDefaultApp(b.Img.File)
					} else {
						b.PlaceHolder.Text = "Please choose an image first"
						b.PlaceHolder.Refresh()
					}
				}
			}), d.CanvasForObject(b))
			pop.ShowAtPosition(e.AbsolutePosition)
			// buttonPos := d.AbsolutePositionForObject(b)
			// buttonSize := b.Size()
			// pop.ShowAtPosition(buttonPos.Add(fyne.NewDelta(0, buttonSize.Height/2)))
		}
	})
	mapp[opts[11]] = Easy()
	mapp[opts[12]] = container.NewGridWithColumns(2, slider(), selectPage())
	mapp[opts[13]] = component.Shuerte()
	mapp[opts[14]] = component.Current()
	mapp[opts[15]] = Clock()
	mapp[opts[16]] = FileTree()
	mapp[opts[17]] = view.LineView()
	mapp[opts[18]] = container.NewVSplit(
		container.NewVBox(xwdg.NewCircleButton(),
			xwdg.NewCircleButton(),
			xwdg.NewCircleButton(),
		), container.NewHBox(
			xwdg.NewCircleButton(),
			xwdg.NewCircleButton(),
		),
	)

	mapp[opts[19]] = xwdg.NewAxis()

	mapp[opts[20]] = Zenity()
	mapp[opts[21]] = Chart()
	lst := widget.NewList(func() int {
		return len(opts)
	}, func() fyne.CanvasObject {
		return widget.NewLabel("aaaaaaaaaaaaaaaaaaa")
	}, func(lii widget.ListItemID, co fyne.CanvasObject) {
		lbl := co.(*widget.Label)
		lbl.SetText(opts[lii])
	})

	makeTray(lst, a)
	center := container.NewMax()
	lst.OnSelected = func(id widget.ListItemID) {
		center.Objects = []fyne.CanvasObject{mapp[opts[id]]}
		center.Refresh()
	}

	// lblMsg := widget.NewLabel("")
	tba := widget.NewToolbarAction(theme.FyneLogo(), func() {})
	// tba.SetIcon(theme.QuestionIcon())
	tba.OnActivated = func() {
		tba.SetIcon(theme.QuestionIcon())
	}
	tb := widget.NewToolbar(tba)
	// obj1 := widget.NewButton("Do", func() {
	// 	tba.SetIcon(theme.QuestionIcon())
	// })
	bd, _ := url.Parse("www.baidu.com")
	obj1 := widget.NewHyperlink("baidu", bd)

	bdb := binding.NewBool()
	bds := binding.BoolToStringWithFormat(bdb, "True %v")
	bdf := binding.NewFloat()
	bds = binding.FloatToStringWithFormat(bdf, "float %.2f")

	bdi := binding.NewInt()
	bds = binding.IntToStringWithFormat(bdi, "int %d")
	lbl1 := widget.NewLabelWithData(bds)

	entry1 := widget.NewMultiLineEntry()
	entry1.Validator = validation.NewTime("2006-01-02")
	entry1.SetText("2006-01-02")

	entry2 := widget.NewMultiLineEntry()
	entry2.SetMinRowsVisible(2)
	entry2.Validator = validation.NewRegexp(`1\d{3}-\d{2}-\d{2}`, "Only years before 2000 are allowed")
	entry2.SetText("1006-01-02")

	entry3 := widget.NewEntry()
	entry3.Validator = validation.NewAllStrings(entry1.Validator, entry2.Validator)
	var bottom *fyne.Container
	obj1.OnTapped = func() {
		bdf.Set(4.56789)
		bdi.Set(12345)
		obj1.SetText("jd")
		bottom.RemoveAll()
		// a.OpenURL(bd)
	}

	bottom = container.NewVBox(entry1, entry2, entry3,
		container.NewHBox(obj1, lbl1),
	)
	content := container.NewBorder(tb, bottom, lst, nil,
		container.NewScroll(center))
	// widget.NewSeparator()
	tabs := container.NewAppTabs(container.NewTabItem("Main", content),
		container.NewTabItem("Source", srcui()))
	tabs.SetTabLocation(container.TabLocationBottom)

	w.SetMainMenu(makeMenu(a, w))
	w.SetContent(tabs)
	w.Resize(fyne.NewSize(1000, 700))
	w.SetMaster()
	w.SetOnClosed(func() {
	})
	w.ShowAndRun()
}

func makeMenu(a fyne.App, w fyne.Window) *fyne.MainMenu {
	openSettings := func() {
		w := a.NewWindow("Fyne Settings")
		w.SetContent(settings.NewSettings().LoadAppearanceScreen(w))
		w.Resize(fyne.NewSize(480, 480))
		w.Show()
	}
	settingsItem := fyne.NewMenuItem("Settings", openSettings)
	settingsShortcut := &desktop.CustomShortcut{KeyName: fyne.KeyComma, Modifier: fyne.KeyModifierShortcutDefault}
	settingsItem.Shortcut = settingsShortcut
	w.Canvas().AddShortcut(settingsShortcut, func(shortcut fyne.Shortcut) {
		openSettings()
	})
	main := fyne.NewMainMenu(
		fyne.NewMenu("File", settingsItem,
			fyne.NewMenuItem("Json Theme", func() {
				th, _ := theme.FromJSON(`{
				"Colors": {"background": "#c0c0c0ff"},
				"Colors-light": {"foreground": "#ffffffff"},
				"Sizes": {"iconInline": 5.0},
				"Fonts": {"monospace": "file://./testdata/NotoMono-Regular.ttf"},
				"Icons": {"cancel": "file://./testdata/cancel_Paths.svg"}
				}`)
				a.Settings().SetTheme(th)
			}), fyne.NewMenuItem("Default Theme", func() {
				a.Settings().SetTheme(&myTheme{})
			})),
	)

	return main
}
func makeTray(lst *widget.List, a fyne.App) {
	if desk, ok := a.(desktop.App); ok {
		h := fyne.NewMenuItem("Hello", func() {})
		m1 := fyne.NewMenuItem("Theme1", func() {
			rd, _ := os.Open("theme1.json")
			th, _ := theme.FromJSONReader(rd)
			a.Settings().SetTheme(th)
		})

		m2 := fyne.NewMenuItem("Theme2", func() {
			rd, _ := os.Open("theme2.json")
			th, _ := theme.FromJSONReader(rd)
			a.Settings().SetTheme(th)
		})
		menu := fyne.NewMenu("Hello World", h, m1, m2)
		h.Action = func() {
			log.Println("System tray menu tapped")
			lst.Select(0)
			h.Label = "Welcome"
			menu.Refresh()
		}
		desk.SetSystemTrayMenu(menu)
	}
}
func decodeImage(input io.Reader) (image.Image, error) {
	src, _, err := image.Decode(input)
	if err != nil {
		return nil, err
	}
	return src, nil
}
