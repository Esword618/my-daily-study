package container

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func HBox(objects ...fyne.CanvasObject) *fyne.Container {
	return container.NewHBox(objects...)
}
func VBox(objects ...fyne.CanvasObject) *fyne.Container {
	return container.NewVBox(objects...)
}
func GridWithRows(rows int, objects ...fyne.CanvasObject) *fyne.Container{
	return container.NewGridWithRows(rows,objects...)
}
func GridWithCols(rows int, objects ...fyne.CanvasObject) *fyne.Container{
	return container.NewGridWithColumns(rows,objects...)
}

func Border(top fyne.CanvasObject, bottom fyne.CanvasObject, left fyne.CanvasObject, right fyne.CanvasObject,
	 objects ...fyne.CanvasObject) *fyne.Container{
return container.NewBorder(top , bottom , left , right , objects ...) 
}