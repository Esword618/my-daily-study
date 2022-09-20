package tool

import (
	"fmt"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/theriault/maths/combinatorics"
	"github.com/theriault/maths/numbertheory"
	"github.com/theriault/maths/statistics"
)

func theriault_maths() {
	w := fyne.CurrentApp().NewWindow("https://hub.fastgit.xyz/theriault/maths")
	input1 := widget.NewEntry()
	input1.SetText("10")
	input2 := widget.NewEntry()
	input2.SetText("10")
	msg := widget.NewLabel("")
	output := widget.NewMultiLineEntry()

	btn1 := widget.NewButton("Factorial", func() {
		n1, _ := strconv.Atoi(input1.Text)
		o1, _ := combinatorics.Factorial(n1)
		output.SetText(fmt.Sprint(o1))
	})
	btn2 := widget.NewButton("Mean", func() {
		ss := strings.Split(input1.Text, ",")
		nn := []int{}
		for _, s := range ss {
			n1, _ := strconv.Atoi(s)
			nn = append(nn, n1)

		}
		o1 := statistics.Mean(nn...)
		output.SetText(fmt.Sprint(o1))
	})
	btn3 := widget.NewButton("PrimeFactorization", func() {
		n1, _ := strconv.Atoi(input1.Text)
		o1 := numbertheory.PrimeFactorization(n1)
		output.SetText(fmt.Sprint(o1))
	})

	btn4 := widget.NewButton("Primorial", func() {
		n1, _ := strconv.Atoi(input1.Text)
		o1, _ := numbertheory.Primorial(n1)
		output.SetText(fmt.Sprint(o1))
	})
	top := container.NewVBox(container.NewGridWithColumns(2,
		input1, input2), container.NewGridWithColumns(6, btn1, btn2,
		btn3, btn4))
	w.SetContent(container.NewBorder(top, msg, nil, nil, output))
	w.Resize(fyne.NewSize(800, 700))
	w.Show()
}
