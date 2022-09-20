package component

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"

	"github.com/gogf/gf/v2/container/garray"
	// "github.com/xyproto/binary"
)

func FileTree(root string, fn func(uid widget.TreeNodeID)) *widget.Tree {
	gs := garray.NewSortedStrArray()
	data := make(map[string][]string)
	sepStr := string(os.PathSeparator)
	filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		dn := strings.Split(path, sepStr)
		switch len(dn) {
		case 1:
			gs.Add(path)
			gs.RemoveValue(".")
			data[""] = gs.Unique().Sort().Slice()
			// data[""] = append(data[""], path)
		case 2:
			data[dn[0]] = append(data[dn[0]], path)
		case 3:
			key := strings.Join(dn[:2], sepStr)
			data[key] = append(data[key], path)
			// data[dn[0]+"/"+dn[1]] = append(data[dn[0]+"/"+dn[1]], path)
		case 4:
			key := strings.Join(dn[:3], sepStr)
			data[key] = append(data[key], path)
			// data[dn[0]+"/"+dn[1]+"/"+dn[2]] =
			// 	append(data[dn[0]+"/"+dn[1]+"/"+dn[2]], path)
		}

		return nil
	})
	tree := widget.NewTreeWithStrings(data)

	tree.UpdateNode = func(uid widget.TreeNodeID, branch bool, node fyne.CanvasObject) {
		lbl := node.(*widget.Label)
		dn := strings.Split(uid, sepStr)
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
	tree.OnSelected = fn
	return tree
}
