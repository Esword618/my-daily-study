package tool

import (
	"bytes"
	"fmt"
	"image"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/ncruces/zenity"
	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
)

// https://hub.fastgit.xyz/stuartmscott/svge
func svge() {
	var d []byte
	l := &widget.Label{
		Wrapping: fyne.TextTruncate,
	}
	v := &canvas.Raster{
		Generator: func(width, height int) image.Image {
			l.SetText("")
			output := image.NewNRGBA(image.Rect(0, 0, width, height))
			if len(d) == 0 {
				l.SetText("No Data")
				return output
			}
			icon, err := oksvg.ReadIconStream(bytes.NewReader(d))
			if err != nil {
				l.SetText(err.Error())
				return output
			}
			inputW, inputH := icon.ViewBox.W, icon.ViewBox.H
			iconAspect := inputW / inputH
			viewAspect := float64(width) / float64(height)
			outputW, outputH := width, height
			if viewAspect > iconAspect {
				outputW = int(float64(height) * iconAspect)
			} else if viewAspect < iconAspect {
				outputH = int(float64(width) / iconAspect)
			}
			scanner := rasterx.NewScannerGV(int(inputW), int(inputH), output, output.Bounds())
			raster := rasterx.NewDasher(width, height, scanner)
			icon.SetTarget(0, 0, float64(outputW), float64(outputH))
			icon.Draw(raster, 1)
			defer func() {
				if r := recover(); r != nil {
					l.SetText(fmt.Sprintf("Crash when rendering SVG: %v", r))
				}
			}()
			return output
		},
	}
	e := &widget.Entry{
		MultiLine: true,
		OnChanged: func(data string) {
			d = []byte(data)
			v.Refresh()
		},
		Wrapping: fyne.TextWrapBreak,
	}

	w := fyne.CurrentApp().NewWindow("Svg Editor")
	s := container.NewHSplit(container.NewVScroll(e), container.NewScroll(v))
	t := widget.NewToolbar(
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			e.SetText("")
			w.SetTitle("SVG Editor")
		}),
		widget.NewToolbarAction(theme.FileIcon(), func() {
			spath, _ := zenity.SelectFile(
				zenity.FileFilters{
					{"SVG file", []string{"*.svg"}},
				},
				zenity.Title("Please choose a svg file"),
				zenity.Filename("."),
				// zenity.Directory(),
			)
			data, err := os.ReadFile(spath)
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			e.SetText(string(data))
			// fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			// 	if err != nil {
			// 		dialog.ShowError(err, w)
			// 		return
			// 	}
			// 	if reader == nil {
			// 		return
			// 	}
			// 	data, err := ioutil.ReadAll(reader)
			// 	if err != nil {
			// 		dialog.ShowError(err, w)
			// 		return
			// 	}
			// 	e.SetText(string(data))
			// 	w.SetTitle("SVG Editor - " + reader.URI().Name())
			// }, w)

			// setLocation(fd,".")
			// fd.SetFilter(storage.NewExtensionFileFilter([]string{".svg"}))
			// fd.Show()
		}),
		widget.NewToolbarAction(theme.DocumentSaveIcon(), func() {
			fd := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
				if err != nil {
					dialog.ShowError(err, w)
					return
				}
				if writer == nil {
					return
				}
				if _, err := writer.Write(d); err != nil {
					dialog.ShowError(err, w)
					return
				}
				l.SetText("Save to " + writer.URI().Path())
			}, w)
			setLocation(fd, ".")
			fd.Show()
		}),
	)
	w.SetContent(container.NewBorder(t, l, nil, nil, s))
	w.CenterOnScreen()
	w.Resize(fyne.NewSize(800, 600))
	w.Show()
}

func setLocation(fd *dialog.FileDialog, path string) {
	luri, _ := storage.ListerForURI(storage.NewFileURI(path))
	fd.SetLocation(luri)
}
