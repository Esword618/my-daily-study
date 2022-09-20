package tool

import (
	"context"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/mholt/archiver/v4"
	"github.com/saracen/fastzip"

	. "gitee.com/y2h/fynex/container"
)

func archiverr() {
	win := fyne.CurrentApp().NewWindow("https://hub.fastgit.xyz/mholt/archiver")

	lblMsg := widget.NewLabel("")
	bdzip := binding.NewString()
	bdzip.Set("example.tar.gz")
	openFile := NewOpenFile(bdzip, "Zip File:", []string{".zip"}, win)
	bdFolderPath := binding.NewString()
	bdFolderPath.Set("./")
	openFolder := NewOpenFolder(bdFolderPath, "Ziped Folder:", win)
	btnZip := widget.NewButton("Archive", func() {
		zippath, err := bdzip.Get()
		lblerr(err, lblMsg)

		zipFolderPath, err := bdFolderPath.Get()
		lblerr(err, lblMsg)
	
		// map files on disk to their paths in the archive
		files, err := archiver.FilesFromDisk(nil, map[string]string{
			"/path/on/disk/file1.txt": "file1.txt",
			"/path/on/disk/file2.txt": "subfolder/file2.txt",
			"/path/on/disk/file3.txt": "",              // put in root of archive as file3.txt
			"/path/on/disk/file4.txt": "subfolder/",    // put in subfolder as file4.txt
			zipFolderPath:    "Custom Folder", // contents added recursively
		})

		lblerr(err, lblMsg)

		// create the output file we'll write to
		out, err := os.Create(zippath)
		lblerr(err, lblMsg)
		defer out.Close()

		// we can use the CompressedArchive type to gzip a tarball
		// (compression is not required; you could use Tar directly)
		format := archiver.CompressedArchive{
			Compression: archiver.Gz{},
			Archival:    archiver.Tar{},
		}
		// create the archive
		err = format.Archive(context.Background(), out, files)
		lblerr(err, lblMsg)
		zipfile, _ := filepath.Abs(zippath)
		folder, _ := filepath.Abs(zipFolderPath)
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
	cc := VBox(GridWithCols(2,
		VBox(openFile, openFolder, btnZip),
		VBox(unzipFile, unzipFolder, btnUnzip),
	),
		btnBrowse,
		lblMsg,
	)

	win.SetContent(cc)
	win.CenterOnScreen()
	win.Resize(fyne.NewSize(700, 700))
	win.Show()
}

func lblerr(err error, lbl *widget.Label) {
	if err != nil {
		lbl.SetText(err.Error())
		return
	}
}
