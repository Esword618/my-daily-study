package main
import (
	"bytes"
	"context"
	"os"
	"path/filepath"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"gitee.com/y2h/fynex/component"
)

func FileTree() fyne.CanvasObject {
	msg := widget.NewLabel("")
	entryFileContent := widget.NewMultiLineEntry()
	content := container.NewMax(entryFileContent)
	tree := component.FileTree(".", nil)
	tree.OnSelected = func(uid widget.TreeNodeID) {
		msg.SetText(uid)
		var str []byte
		ctx := context.Background()
		done := make(chan struct{}, 1)
		go func(ctx context.Context) {
			// 超时任务，如 发送HTTP请求
			switch filepath.Ext(uid) {
			case ".jpg", ".png", ".jpeg":
				bts, _ := os.ReadFile(uid)
				rd := bytes.NewBuffer(nil)
				rd.Write(bts)
				content.Objects = []fyne.CanvasObject{
					canvas.NewImageFromReader(rd, uid),
				}
			default:
				str, _ = os.ReadFile(uid)
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
