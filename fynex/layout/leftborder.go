package layout

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// Declare conformity with Layout interface
var _ fyne.Layout = (*borderLayout)(nil)

type borderLayout struct {
	left   fyne.CanvasObject
	isLeft bool
}

func NewLeftBorder(left fyne.CanvasObject, objects ...fyne.CanvasObject) *fyne.Container {
	all := objects

	if left != nil {
		all = append(all, left)
	}

	return fyne.NewContainerWithLayout(NewBorderLayout(left, true), all...)
}

func NewRightBorder(left fyne.CanvasObject, objects ...fyne.CanvasObject) *fyne.Container {
	all := objects

	if left != nil {
		all = append(all, left)
	}

	return fyne.NewContainerWithLayout(NewBorderLayout(left, false), all...)
}

// NewBorderLayout creates a new BorderLayout instance with top, bottom, left
// and right objects set. All other items in the container will fill the centre
// space
func NewBorderLayout(left fyne.CanvasObject, isLeft bool) fyne.Layout {
	return &borderLayout{left, isLeft}
}

// Layout is called to pack all child objects into a specified size.
// For BorderLayout this arranges the top, bottom, left and right widgets at
// the sides and any remaining widgets are maximised in the middle space.
func (b *borderLayout) Layout(objects []fyne.CanvasObject, size fyne.Size) {
	var leftSize fyne.Size
	if b.left != nil && b.left.Visible() {
		b.left.Resize(fyne.NewSize(b.left.MinSize().Width, size.Height))
		if b.isLeft {
			b.left.Move(fyne.NewPos(0, 0))
		} else {
			b.left.Move(fyne.NewPos(size.Width-b.left.MinSize().Width, 0))
		}
		leftSize = fyne.NewSize(b.left.MinSize().Width+theme.Padding(), size.Height)
	}

	middleSize := fyne.NewSize(size.Width-leftSize.Width, size.Height)
	middlePos := fyne.NewPos(0, 0)
	if b.isLeft {
		middlePos.X = leftSize.Width
		// mpos := fyne.NewPos(leftSize.Width, 0)
	} else {

	}
	for _, child := range objects {
		if !child.Visible() {
			continue
		}

		if child != b.left {
			child.Resize(middleSize)
			child.Move(middlePos)
		}
	}
}

// MinSize finds the smallest size that satisfies all the child objects.
// For BorderLayout this is determined by the MinSize height of the top and
// plus the MinSize width of the left and right, plus any padding needed.
// This is then added to the union of the MinSize for any remaining content.
func (b *borderLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	minSize := fyne.NewSize(0, 0)
	for _, child := range objects {
		if !child.Visible() {
			continue
		}

		if child != b.left {
			minSize = minSize.Max(child.MinSize())
		}
	}

	if b.left != nil && b.left.Visible() {
		minHeight := fyne.Max(minSize.Height, b.left.MinSize().Height)
		minSize = fyne.NewSize(minSize.Width+b.left.MinSize().Width+theme.Padding(), minHeight)
	}

	return minSize
}
