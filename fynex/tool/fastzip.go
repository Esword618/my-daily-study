package tool

import (
	"context"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/saracen/fastzip"
	"github.com/saracen/walker"
)

func fastZip() {
	win := fyne.CurrentApp().NewWindow("Html to txt")

	lblMsg := widget.NewLabel("")
	bdzip := binding.NewString()
	bdzip.Set("archive.zip")
	openFile := NewOpenFile(bdzip, "Zip File:", []string{".zip"}, win)
	bdFolderPath := binding.NewString()
	bdFolderPath.Set("./")
	openFolder := NewOpenFolder(bdFolderPath, "Ziped Folder:", win)
	btnZip := widget.NewButton("Zip", func() {

		zippath, err := bdzip.Get()
		if err != nil {
			lblMsg.SetText(err.Error())
			return
		}
		zipFolderPath, err := bdFolderPath.Get()
		if err != nil {
			lblMsg.SetText(err.Error())
			return
		}
		// Create archive file
		w, err := os.Create(zippath)
		if err != nil {
			lblMsg.SetText(err.Error())
			return
		}
		defer w.Close()

		// Create new Archiver
		a, err := fastzip.NewArchiver(w, zipFolderPath)
		if err != nil {
			lblMsg.SetText(err.Error())
			return
		}
		defer a.Close()

		// Register a non-default level compressor if required
		// a.RegisterCompressor(zip.Deflate, fastzip.FlateCompressor(1))

		// Walk directory, adding the files we want to add
		files := make(map[string]os.FileInfo)
		// err = filepath.Walk(zipFolderPath,
		// 	func(pathname string, info os.FileInfo, err error) error {
		// 		files[pathname] = info
		// 		return nil
		// 	})

		// walk function called for every path found
		walkFn := func(pathname string, fi os.FileInfo) error {
			files[pathname] = fi
			// fmt.Printf("%s: %d bytes\n", pathname, fi.Size())
			return nil
		}

		// error function called for every error encountered
		errorCallbackOption := walker.WithErrorCallback(
			func(pathname string, err error) error {
				// ignore permissione errors
				if os.IsPermission(err) {
					return nil
				}
				// halt traversal on any other error
				return err
			})

		walker.Walk(zipFolderPath, walkFn, errorCallbackOption)

		// Archive
		if err = a.Archive(context.Background(), files); err != nil {
			lblMsg.SetText(err.Error())
			return
		}
		folder, _ := filepath.Abs(zipFolderPath)
		zipfile, _ := filepath.Abs(zippath)
		lblMsg.SetText("Zip {" + folder + "} to {" + zipfile + "}")
	})
	bdunzip := binding.NewString()
	bdunzip.Set("archive.zip")
	unzipFile := NewOpenFile(bdunzip, "Zip File:", []string{".zip"}, win)
	bdunzipFolderPath := binding.NewString()
	bdunzipFolderPath.Set("./")
	unzipFolder := NewOpenFolder(bdunzipFolderPath, "Unziped Folder:", win)
	btnUnzip := widget.NewButton("Unzip", func() {
		zippath, err := bdunzip.Get()
		if err != nil {
			lblMsg.SetText(err.Error())
			return
		}
		unzippath, err := bdunzipFolderPath.Get()
		if err != nil {
			lblMsg.SetText(err.Error())
			return
		}
		// Create new extractor
		e, err := fastzip.NewExtractor(zippath, unzippath)
		if err != nil {
			lblMsg.SetText(err.Error())
			return
		}
		defer e.Close()
		e.Files()
		// Extract archive files
		if err = e.Extract(context.Background()); err != nil {
			lblMsg.SetText(err.Error())
			return
		}
		folder, _ := filepath.Abs(unzippath)
		zipfile, _ := filepath.Abs(zippath)
		lblMsg.SetText("UnZip {" + zipfile + "} to {" + folder + "}")
	})
	btnBrowse := widget.NewButton("Browse Zip file", func() {
		zippath, err := bdunzip.Get()
		if err != nil {
			lblMsg.SetText(err.Error())
			return
		}
		unzippath, err := bdunzipFolderPath.Get()
		if err != nil {
			lblMsg.SetText(err.Error())
			return
		}
		// Create new extractor
		e, err := fastzip.NewExtractor(zippath, unzippath)
		if err != nil {
			lblMsg.SetText(err.Error())
			return
		}
		defer e.Close()
		ent := widget.NewMultiLineEntry()
		for _, f := range e.Files() {
			ent.Text = ent.Text + f.Name + "\n"
		}

		pop := widget.NewPopUp(ent, win.Canvas())
		pop.Resize(fyne.NewSize(600, 600))
		pop.Show()
	})
	cc := container.NewVBox(
		container.NewGridWithColumns(2,
			container.NewVBox(openFile, openFolder, btnZip),
			container.NewVBox(unzipFile, unzipFolder, btnUnzip),
		),
		btnBrowse,
		lblMsg,
	)

	win.SetContent(cc)
	win.CenterOnScreen()
	win.Resize(fyne.NewSize(700, 700))
	win.Show()
}

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
