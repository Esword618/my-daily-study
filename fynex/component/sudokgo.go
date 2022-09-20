package component

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/andydotxyz/sudokgo"
)

type gui struct {
	sudoku *sudokgo.Sudoku
	score  *widget.Label
	win    fyne.Window
	cells  [sudokgo.RowSize][sudokgo.RowSize]*widget.Entry
}

func (g *gui) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	s := fyne.Min(size.Width, size.Height)
	offsetX := float32(0)
	if s < size.Width {
		offsetX = (size.Width - s) / 2
	}

	childSize := fyne.NewSize(s, s)
	childPos := fyne.NewPos(offsetX, 0)
	for _, o := range objects {
		o.Resize(childSize)
		o.Move(childPos)
	}
}

func (g *gui) MinSize(objects []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(350, 350)
}

// Show starts a new sudoku UI
// func main() {
// 	a := app.New()
// 	a.SetIcon(sudokuIcon)
// 	game := sudokgo.NewSudoku()
// 	gui := newGUI(game)

// 	win := a.NewWindow("SudokGo")
// 	content := gui.LoadUI(win)
// 	go gui.generate()

// 	win.SetIcon(sudokuIcon)
// 	win.SetContent(content)
// 	win.ShowAndRun()
// }
func SudokGo(win fyne.Window) fyne.CanvasObject {
	game := sudokgo.NewSudoku()
	gui := newGUI(game)

	content := gui.LoadUI(win)
	go gui.generate()

	return content
}
func (g *gui) newCell(x, y int) *widget.Entry {
	entry := widget.NewEntry()
	g.cells[x][y] = entry
	return entry
}

func (g *gui) newSubGrid(c *fyne.Container, x, y int) {
	grid := container.NewGridWithColumns(sudokgo.GridSize)
	// grid := fyne.NewContainerWithLayout(layout.NewGridLayout(sudokgo.GridSize))
	c.AddObject(container.NewMax(
		canvas.NewRectangle(theme.BackgroundColor()), grid))

	for j := 0; j < sudokgo.GridSize; j++ {
		for i := 0; i < sudokgo.GridSize; i++ {
			grid.Add(g.newCell(x+i, y+j))
		}
	}

}

func (g *gui) makeGrid(cells *fyne.Container) {
	// make the sub-grids
	x, y := 0, 0
	for i := 0; i < sudokgo.RowSize; i++ {
		g.newSubGrid(cells, x, y)

		x += sudokgo.GridSize
		if x == sudokgo.RowSize {
			y += sudokgo.GridSize
			x = 0
		}
	}
}

func (g *gui) refresh(updateDisabled bool) {
	for x := 0; x < sudokgo.RowSize; x++ {
		for y := 0; y < sudokgo.RowSize; y++ {
			entry := g.cells[x][y]
			value := g.sudoku.Grid[x][y]
			if value == -1 {
				entry.SetText("")
				if updateDisabled {
					entry.Enable()
				}
			} else {
				entry.SetText(fmt.Sprintf("%d", value))
				if updateDisabled {
					entry.Disable()
				}
			}
		}
	}
}

func (g *gui) loadToolbar() *widget.Toolbar {
	score := newToolbarText("")
	g.score = score.(*toolbarText).text
	return widget.NewToolbar(
		widget.NewToolbarAction(theme.FileIcon(), g.generate),
		widget.NewToolbarAction(theme.ConfirmIcon(), g.submit),

		widget.NewToolbarSpacer(),
		score,
		widget.NewToolbarSpacer(),

		widget.NewToolbarAction(theme.MediaSkipNextIcon(), g.solve),
		widget.NewToolbarAction(theme.ViewRefreshIcon(), g.reset))
}

func (g *gui) LoadUI(win fyne.Window) fyne.CanvasObject {
	g.win = win
	// cells := fyne.NewContainerWithLayout(layout.NewGridLayout(sudokgo.GridSize))
	cells := container.NewGridWithColumns(sudokgo.GridSize)
	g.makeGrid(cells)

	toolbar := g.loadToolbar()
	content := container.New(g, canvas.NewRectangle(color.Black), cells)
	return container.NewBorder(toolbar, nil, nil, nil, content)
	// content := fyne.NewContainerWithLayout(g, canvas.NewRectangle(color.Black), cells)
	// return fyne.NewContainerWithLayout(layout.NewBorderLayout(toolbar, nil, nil, nil),
	// 	toolbar, content)
}

func (g *gui) generate() {
	difficulty := sudokgo.ScoreEasy
	score := g.sudoku.Generate(difficulty)
	g.score.SetText(fmt.Sprintf("%s (%d)", sudokgo.Difficulty(difficulty), score))
	g.refresh(true)
}

func (g *gui) solve() {
	g.sudoku.Solve()
	g.refresh(false)
}

func (g *gui) submit() {
	for x := 0; x < sudokgo.RowSize; x++ {
		for y := 0; y < sudokgo.RowSize; y++ {
			entry := g.cells[x][y]

			if entry.Text == "" {
				dialog.ShowInformation("Verify puzzle",
					"Not all cells completed", g.win)
				return
			}
		}
	}

	g.sudoku.Solve()
	for x := 0; x < sudokgo.RowSize; x++ {
		for y := 0; y < sudokgo.RowSize; y++ {
			entry := g.cells[x][y]
			value := g.sudoku.Grid[x][y]

			if entry.Text != fmt.Sprintf("%d", value) {
				dialog.ShowInformation("Verify puzzle",
					"Error in puzzle values", g.win)
				return
			}
		}
	}

	dialog.ShowInformation("Verify puzzle",
		"Congratulations on completing the puzzle", g.win)
}

func (g *gui) reset() {
	for x := 0; x < sudokgo.RowSize; x++ {
		for y := 0; y < sudokgo.RowSize; y++ {
			entry := g.cells[x][y]

			if !entry.Disabled() {
				// entry.SetText("")
				entry.Text = ""
			}
		}
	}
	g.win.Content().Refresh()
}

type toolbarText struct {
	text *widget.Label
}

func (t *toolbarText) ToolbarObject() fyne.CanvasObject {
	return t.text
}

func newToolbarText(text string) widget.ToolbarItem {
	return &toolbarText{text: widget.NewLabel(text)}
}

func newGUI(s *sudokgo.Sudoku) *gui {
	ret := &gui{sudoku: s}

	return ret
}
