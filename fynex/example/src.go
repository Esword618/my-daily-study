package main

import (
	"bytes"
	"context"
	iofs "io/fs"
	"path/filepath"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"gitee.com/y2h/fynex"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/leaanthony/debme"
	// "github.com/xyproto/binary"
)

func srcui() fyne.CanvasObject {
	gs := garray.NewSortedStrArray()
	data := make(map[string][]string)
	root, _ := debme.FS(fynex.Srcs, ".")
	// println(root.ReadFile("main.go"))
	// Fully compatible FS
	iofs.WalkDir(root, ".", func(path string, d iofs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		dn := strings.Split(path, "/")
		switch len(dn) {
		case 1:
			gs.Add(path)
			gs.RemoveValue(".")
			data[""] = gs.Unique().Sort().Slice()
			// data[""] = append(data[""], path)
		case 2:
			data[dn[0]] = append(data[dn[0]], path)
		case 3:
			key := strings.Join(dn[:2], "/")
			data[key] = append(data[key], path)
			// data[dn[0]+"/"+dn[1]] = append(data[dn[0]+"/"+dn[1]], path)
		case 4:
			key := strings.Join(dn[:3], "/")
			data[key] = append(data[key], path)
			// data[dn[0]+"/"+dn[1]+"/"+dn[2]] =
			// 	append(data[dn[0]+"/"+dn[1]+"/"+dn[2]], path)
		}

		return nil
	})
	// fmt.Printf("data: %v\n", data)
	msg := widget.NewLabel("")
	entryFileContent := widget.NewMultiLineEntry()
	content := container.NewMax(entryFileContent)
	tree := widget.NewTreeWithStrings(data)

	tree.UpdateNode = func(uid widget.TreeNodeID, branch bool, node fyne.CanvasObject) {
		lbl := node.(*widget.Label)
		dn := strings.Split(uid, "/")
		switch len(dn) {
		case 1:
			lbl.SetText(uid)
		case 2:
			lbl.SetText(dn[1])
		case 3:
			lbl.SetText(dn[2])
		case 4:
			lbl.SetText(dn[3])
		}
	}
	tree.OnSelected = func(uid widget.TreeNodeID) {
		msg.SetText(uid)
		var str []byte
		ctx := context.Background()
		done := make(chan struct{}, 1)
		go func(ctx context.Context) {
			// 超时任务，如 发送HTTP请求
			switch filepath.Ext(uid) {
			case ".jpg", ".png", ".jpeg":
				bts, _ := root.ReadFile(uid)
				rd := bytes.NewBuffer(nil)
				rd.Write(bts)
				content.Objects = []fyne.CanvasObject{
					canvas.NewImageFromReader(rd, uid),
				}
			default:
				str, _ = root.ReadFile(uid)
				entryFileContent.SetText(string(str)) //碰到大文件会阻塞界面
				// entryFileContent.SetText("Can not display the file")
				content.Objects = []fyne.CanvasObject{
					entryFileContent,
				}
			}
			// isBinary, err := binary.File(uid)
			// if err != nil {
			// 	entryFileContent.SetText("Read file error:" + err.Error())
			// }
			// if !isBinary {
			// 	str, _ = root.ReadFile(uid)
			// 	entryFileContent.SetText(string(str)) //碰到大文件会阻塞界面
			// }
			// else {
			//
			// 	}
			// }
			done <- struct{}{}
		}(ctx)
		select {
		case <-done:
			// fmt.Println("call successfully!!!")
			return
		case <-time.After(time.Duration(time.Millisecond * 100)):
			// fmt.Println("timeout!!!")
			entryFileContent.SetText("the file is too big")
			return
		}
	}
	return container.NewBorder(nil, msg, tree, nil, content)
}
