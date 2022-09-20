package component

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// bdFilePath := binding.NewString()
// filters:=[]string{".jpg",".jpeg",".png"}
// filters:=[]string{".zip",".rar",".tar"}
// openFile := NewOpenFile(bdFilePath,filters w)
func NewOpenFile(data binding.String, lbl string, filters []string,
	w fyne.Window) *fyne.Container {
	return container.NewBorder(nil, nil, widget.NewLabel(lbl),
		widget.NewButtonWithIcon("", theme.FileIcon(), func() {
			d := dialog.NewFileOpen(func(uc fyne.URIReadCloser, e error) {
				if uc == nil {
					return
				}
				data.Set(uc.URI().String()[7:])
			}, w)
			luri, _ := storage.ListerForURI(storage.NewFileURI("."))

			d.SetFilter(storage.NewExtensionFileFilter(filters))
			d.SetLocation(luri)
			d.Show()
		}),
		widget.NewEntryWithData(data))
}

// bdFolderPath := binding.NewString()
// openFolder := NewOpenFolder(outData, w)
func NewOpenFolder(data binding.String, lbl string, w fyne.Window) *fyne.Container {
	return container.NewBorder(nil, nil, widget.NewLabel(lbl),
		widget.NewButtonWithIcon("",
			theme.FolderOpenIcon(), func() {
				d := dialog.NewFolderOpen(func(lu fyne.ListableURI, e error) {
					if lu == nil {
						return
					}
					data.Set(lu.Path())
				}, w)
				luri, _ := storage.ListerForURI(storage.NewFileURI("."))
				d.SetLocation(luri)
				d.Show()
			}), widget.NewEntryWithData(data))
}
