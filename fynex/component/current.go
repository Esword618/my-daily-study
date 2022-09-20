package component

import (
	"path"
	"runtime"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func Current() fyne.CanvasObject {
	return container.NewVBox(
		widget.NewLabel("current file name: "+CurrentFile()),
		widget.NewLabel("current directory: "+CurrentDir()),
		widget.NewLabel("current line: "+strconv.Itoa(CurrentLine())),
		widget.NewLabel("current function name: "+CurrentFunc()),
		widget.NewLabel("current file basename: "+Basename()),
	)
}

// https://hub.fastgit.org/otiai10/curr
const (
	depthOfFunctionCaller = 1
)

// File is current file name provider,
// like `__FILE__` of PHP.
func CurrentFile() string {
	_, fi, _, _ := runtime.Caller(depthOfFunctionCaller)
	return fi
}

// Basename is current file basename provider,
// like `basename(__FILE__)` of PHP.
func Basename() string {
	_, fi, _, _ := runtime.Caller(depthOfFunctionCaller)
	return path.Base(fi)
}

// Dir is current directory provider,
// like `__DIR__` of PHP.
func CurrentDir() string {
	_, fi, _, _ := runtime.Caller(depthOfFunctionCaller)
	return path.Dir(fi)
}

// Func is current function name provider,
// like `__FUNCTION__` of PHP.
func CurrentFunc() string {
	pc, _, _, _ := runtime.Caller(depthOfFunctionCaller)
	fn := runtime.FuncForPC(pc)
	elems := strings.Split(fn.Name(), ".")
	return elems[len(elems)-1]
}

// Line is current line provider,
// like `__LINE__` of PHP.
func CurrentLine() int {
	_, _, li, _ := runtime.Caller(depthOfFunctionCaller)
	return li
}
