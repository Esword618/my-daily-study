package component

import (
	_ "embed"
	"image/color"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	ctn "fyne.io/fyne/v2/container"

	// "fyne.io/fyne/v2/driver/software"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	wdg "fyne.io/fyne/v2/widget"
	"gitee.com/y2h/fynex/common"
	xwdg "gitee.com/y2h/fynex/widget"
	"gitee.com/y2h/goutil"
	"github.com/tidwall/gjson"
	"golang.org/x/image/colornames"
)

//go:embed shuerte.json
var titles string

func Shuerte() fyne.CanvasObject {
	txtCol := wdg.NewEntry()
	col := 4
	content := ctn.NewMax()
	txtCol.SetText("3")
	txtCol.OnChanged = func(s string) {
		c, _ := strconv.Atoi(s)
		col = c
	}

	// isDone == 1 表示游戏结束；isDone == 2 表示游戏开始；isDone == 0 表示游戏未开始；
	isDone := 0

	txtTimeUsed := wdg.NewLabel("")
	timer := func() { /* 计时 */
		startTime := time.Now()
		for range time.Tick(time.Millisecond) {
			if isDone == 1 {
				return
			}
			if isDone == 2 {
				txtTimeUsed.SetText(time.Since(startTime).Truncate(time.Millisecond).String())
			}
			time.Sleep(time.Millisecond * 100)
		}
	}

	var width float32 = 600

	// original data
	data := []string{}
	// shuffled data
	sdata := []string{}
	dataResult := []string{}
	cellWidth := float32(0)
	txtOrigin := wdg.NewMultiLineEntry()
	txtOrigin.Wrapping = fyne.TextWrapWord

	txtResult := wdg.NewEntry()
	txtSuccess := wdg.NewLabel("")
	// txtOrigin.SetText("氢,氦,锂,铍,硼,碳,氮,氧,氟")

	// 使用自定义按钮，方便更换颜色、大小
	var cellBtns [64]*xwdg.TextField

	for i := 0; i < 64; i++ {
		btn := xwdg.NewTextField().Do(func(t *xwdg.TextField) {
			t.Background.FillColor = colornames.Aliceblue
			t.Background.StrokeColor = colornames.Antiquewhite
			t.Background.StrokeWidth = 2
			t.Txt1.Alignment = fyne.TextAlignCenter
			t.OnTapped = func(tf *xwdg.TextField) {
				len_data := len(data)
				rdl := len(dataResult)
				if rdl < len_data {
					if tf.Txt1.Text == data[rdl] {
						dataResult = append(dataResult, tf.Txt1.Text)
						txtResult.SetText(strings.Join(dataResult, ","))
						tf.Txt1.Color = color.RGBA{0, 0, 255, 255}
						tf.Txt1.Refresh()
					}
				}

				if rdl == 1 {
					isDone = 2
				}
				if rdl+1 == len_data {
					txtSuccess.SetText("done")
					isDone = 1
				}

			}
		})
		cellBtns[i] = btn
	}

	size := fyne.NewSize(0, 0) // 每个按钮的大小
	i := 0                     // 内容数组索引
	grid := container.NewGridWithColumns(col)
	refresh := func() {
		// 重置相关变量和控件
		txtSuccess.SetText("")
		txtResult.SetText("")
		txtTimeUsed.SetText("")
		dataResult = []string{}
		isDone = 0
		data = setStrData(txtOrigin.Text)
		if col < 2 {
			col = 2
		}
		if col*col > len(data) {
			col = int(math.Floor(math.Sqrt(float64(len(data)))))
		}
		txtCol.SetText(strconv.Itoa(col))
		// data = setStrData(txtOrigin.Text)
		data = data[:col*col]
		grid = container.NewGridWithColumns(col)
		go timer()

		if width == 0 {
			width = 300
		}
		cellWidth = width / float32(col) * 0.6
		size.Height = cellWidth
		size.Width = cellWidth
		sdata = shuffleStrArr(data)
		txtSize := cellWidth * 0.9 / float32(col)
		if txtSize < 18 {
			txtSize = 18
		}
		i = 0
		for c := 0; c < col; c++ {
			for r := 0; r < col; r++ {
				cellBtns[i].Do(func(t *xwdg.TextField) {
					t.Txt1.Text = sdata[i]
					t.Txt1.TextSize = txtSize
					t.Txt1.Color = theme.TextColor()
					t.Txt1.SetMinSize(fyne.NewSize(cellWidth, cellWidth))
					// t.Txt1.Refresh()
					// t.Background.Resize(fyne.NewSize(cellWidth, cellWidth))
					// t.Background.SetMinSize(fyne.NewSize(cellWidth, cellWidth))
				})
				grid.Add(cellBtns[i])
				i++
			}
		}
		content.Objects = []fyne.CanvasObject{grid}
		content.Refresh()
	}

	// refresh()
	btnRefresh := wdg.NewButtonWithIcon("", theme.ViewRefreshIcon(), func() {
		refresh()
	})

	numEnt := xwdg.NumberEntry()
	numEnt.SetText("3")
	btnSave := wdg.NewButtonWithIcon("", theme.DocumentSaveIcon(), func() {
		num, _ := strconv.Atoi(numEnt.Text)
		size := fyne.NewSize(600, 600)
		for i := 0; i < num; i++ {
			// img := software.Render(grid, fyne.CurrentApp().Settings().Theme())
			img := common.Capture(grid, size, 1)
			goutil.SavePNG("./shuerte"+strconv.Itoa(i)+".png", img)
			refresh()
			time.Sleep(time.Second)
		}

		txtSuccess.SetText("saved done")
	})
	btnRefresh.Importance = wdg.HighImportance
	btnColPlus := wdg.NewButton("+", func() {
		col++
		refresh()
	})
	btnColMinus := wdg.NewButton("-", func() {
		col--
		refresh()
	})

	rdg := wdg.NewRadioGroup(goutil.NumOpts(2, 9),
		func(s string) {
			col, _ = strconv.Atoi(s)
			refresh()
		})
	rdg.Horizontal = true

	rdg.SetSelected("2")
	frm := wdg.NewForm(
		wdg.NewFormItem("Status:", ctn.NewHBox(txtSuccess, layout.NewSpacer(), txtTimeUsed)),
		wdg.NewFormItem("Columns:",
			ctn.NewHBox(txtCol, btnColPlus, btnColMinus,
				btnRefresh, numEnt, btnSave)),
		wdg.NewFormItem("Result \n data:", txtResult),
		wdg.NewFormItem("Original \n data:", txtOrigin),
	)
	var titleList []string
	gjson.Get(titles, "..#.title").ForEach(func(key, value gjson.Result) bool {
		titleList = append(titleList, value.String())
		return true
	})
	var contentList []string
	gjson.Get(titles, "..#.content").ForEach(func(key, value gjson.Result) bool {
		contentList = append(contentList, value.String())
		return true
	})

	tbs := ctn.NewAppTabs()
	tbs.SetTabLocation(ctn.TabLocationBottom)
	lstTitle := CyyList(titleList)
	lstTitle.OnSelected = func(id wdg.ListItemID) {
		t := contentList[id]
		txtOrigin.SetText(t)

		refresh()
		if fyne.CurrentDevice().IsMobile() {
			tbs.SelectTabIndex(1)
		}
	}
	lstTitle.Select(3)

	tutorial := ctn.NewVScroll(ctn.NewBorder(nil,
		container.NewVBox(rdg, frm), nil, nil, content))
	txtJson := wdg.NewMultiLineEntry()
	txtJson.SetText(titles)
	txtJson.Wrapping = fyne.TextWrapWord

	splt := ctn.NewHSplit(lstTitle, tutorial)
	splt.SetOffset(0.3)
	tabs := ctn.NewAppTabs(
		ctn.NewTabItem("main", splt),
		ctn.NewTabItem("json", txtJson),
	)
	tabs.SetTabLocation(ctn.TabLocationBottom)
	tabs.OnChanged = func(tab *ctn.TabItem) {
		if tab.Text == "main" {
			refresh()
		}
	}

	return tabs
}
func shuffleStrArr(nums []string) (snums []string) {
	rand.Seed(time.Now().UTC().UnixNano())
	snums = append(snums, nums...)
	aliceRng := rand.New(rand.NewSource(time.Now().Unix()))
	aliceRng.Shuffle(
		len(snums),
		func(i, j int) {
			snums[i], snums[j] = snums[j], snums[i]
		})
	return
}

func setStrData(str string) (arr []string) {
	if strings.Contains(str, ",") {
		return strings.Split(str, ",")
	} else {
		for _, s := range str {
			arr = append(arr, string(s))
		}
		return
	}
}
