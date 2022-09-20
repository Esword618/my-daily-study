package tool

import (
	"errors"
	"fmt"
	"image/color"
	"math/rand"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	wdg "fyne.io/fyne/v2/widget"
	. "gitee.com/y2h/fynex/widget"
	"github.com/jacalz/eval"
)

var (
	red1   = color.RGBA{254, 67, 101, 255}
	red2   = color.RGBA{252, 157, 154, 255}
	green1 = color.RGBA{131, 175, 155, 255}
)

func jacalzEval() {
	w := fyne.CurrentApp().NewWindow("https://hub.fastgit.xyz/Jacalz/eval")
	rand.Seed(time.Now().UTC().UnixNano())
	lblRight := canvas.NewText("right:0", red1)
	lblWrong := canvas.NewText("wrong:0", green1)
	lblRight.TextSize = 16
	lblWrong.TextSize = 16

	numRight := 0
	numWrong := 0
	sel1 := wdg.NewRadioGroup(NumSlice(5), nil)
	sel1.Horizontal = true
	map1 := map[string]int{"1": 10, "2": 100, "3": 1000, "4": 10000}
	sel2 := wdg.NewRadioGroup([]string{"+", "-", "*", "/"}, nil)
	sel2.Horizontal = true

	sel3 := wdg.NewRadioGroup([]string{"2", "3", "4", "5"}, nil)
	sel3.Horizontal = true
	sel3.SetSelected("2")

	sel1.SetSelected("2")
	const n = 10
	data := []int{}
	right := [10]int{}
	exp := [10]string{}
	resultEntries := [10]*wdg.Entry{}

	refreshData := func() {
		if _, ok := map1[sel1.Selected]; !ok {
			return
		}
		if sel3.Selected == "" {
			return
		}
		data = []int{}
		num, _ := strconv.Atoi(sel3.Selected)
		for i := 0; i < num*n; i++ {
			data = append(data, rand.Intn(map1[sel1.Selected]))
			// data[i] = rand.Intn(map1[sel1.Selected])
		}
		switch sel3.Selected {
		case "3":
			for i := 0; i < 10; i++ {
				exp[i] = fmt.Sprint(data[i], " ", sel2.Selected, " ",
					data[i+10], " ", sel2.Selected, " ", data[i+20])
			}

		case "4":
			for i := 0; i < 10; i++ {
				exp[i] = fmt.Sprint(data[i],
					" ", sel2.Selected, " ", data[i+10],
					" ", sel2.Selected, " ", data[i+20],
					" ", sel2.Selected, " ", data[i+30])
			}
		case "5":
			for i := 0; i < 10; i++ {
				exp[i] = fmt.Sprint(data[i],
					" ", sel2.Selected, " ", data[i+10],
					" ", sel2.Selected, " ", data[i+20],
					" ", sel2.Selected, " ", data[i+30],
					" ", sel2.Selected, " ", data[i+40])
			}
		default:
			for i := 0; i < 10; i++ {
				exp[i] = fmt.Sprint(data[i], " ", sel2.Selected, " ",
					data[i+10])
			}
		}
		for i := 0; i < 10; i++ {
			result, _ := eval.Evaluate(exp[i])
			right[i] = int(result)

			j := i
			if resultEntries[j] != nil {
				resultEntries[j].SetText("")
			}
		}
		numRight = 0
		numWrong = 0
		lblWrong.Text = ("wrong:")
		lblRight.Text = ("right:")
		// lblRight.Refresh()
		// lblWrong.Refresh()
	}
	refreshData()
	list1 := wdg.NewList(func() int {
		return 10
	}, func() fyne.CanvasObject {
		return item()
	}, func(lii wdg.ListItemID, co fyne.CanvasObject) {
		grid := co.(*fyne.Container)
		left := grid.Objects[0].(*canvas.Text)
		left.Alignment = fyne.TextAlignTrailing
		left.TextSize = 16
		righto := grid.Objects[1].(*wdg.Entry)
		righto.Validator = func(s string) error {
			si, err := strconv.Atoi(s)
			if err != nil {
				return err
			}
			if si != right[lii] {
				return errors.New("Wrong")
			}
			return nil
		}
		left.Text = exp[lii] + " ="
		// n3.SetText("")
		resultEntries[lii] = righto
	})
	refreshFn := func(s string) {
		if s == "" {
			return
		}
		refreshData()
		list1.Refresh()
	}

	sel1.OnChanged = func(s string) {
		refreshFn(s)
	}
	sel1.SetSelected("2")
	sel2.OnChanged = func(s string) {
		refreshFn(s)
	}
	sel2.SetSelected("+")

	sel3.OnChanged = func(s string) {
		refreshFn(s)
	}
	btnDone := wdg.NewButton("Done and Submit", func() {
		if sel2.Selected == "" {
			return
		}
		numWrong = 0
		numRight = 0
		for i := 0; i < 10; i++ {
			nn3, _ := strconv.Atoi(resultEntries[i].Text)
			if right[i] != nn3 {
				numWrong += 1
			} else {
				numRight += 1
			}
		}

		lblWrong.Text = (fmt.Sprint("wrong:", numWrong))
		lblRight.Text = (fmt.Sprint("right:", numRight))
		lblRight.Refresh()
		lblWrong.Refresh()
	})
	c1 := container.NewVBox(wdg.NewLabel("Compute Happy"),
		wdg.NewSeparator(),
		container.NewHBox(lblRight, layout.NewSpacer(), sel1),
		wdg.NewSeparator(),
		container.NewHBox(lblWrong, layout.NewSpacer(), sel2),
		container.NewGridWithColumns(2, container.NewHBox(
			canvas.NewText("Num: ", red2), sel3), btnDone),
		wdg.NewSeparator(),
	)
	cc := container.NewBorder(c1, nil, nil, nil, list1)

	w.SetContent(cc)
	w.Resize(fyne.NewSize(900, 600))
	w.CenterOnScreen()
	w.FixedSize()
	w.Show()
}
func item() *fyne.Container {
	return container.NewGridWithColumns(2,
		canvas.NewText("", red1), NumberEntry())
}

func NumSlice(n int) (result []string) {
	for i := 1; i < n; i++ {
		result = append(result, strconv.Itoa(i))
	}
	return
}
