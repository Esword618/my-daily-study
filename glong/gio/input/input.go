/*
 * @Author: Esword
 * @Description:
 * @FileName:  input
 * @Version: 1.0.0
 * @Date: 2022-06-18 18:39
 */

package main

import (
	"fmt"
	"os"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
	"gioui.org/x/notify"

	alo "MoocDownload/gui/input/applayout"
)

func main() {

	go func() {
		w := app.NewWindow(
			app.Title("input"),
		)
		var ops op.Ops
		p := Page{}
		for event := range w.Events() {
			switch event := event.(type) {
			case system.DestroyEvent:
				os.Exit(0)
			case system.FrameEvent:
				// event.Frame(p.frame(layout.NewContext(&ops, event)))
				gtx := layout.NewContext(&ops, event)
				p.frame(gtx)
				event.Frame(gtx.Ops)

			default:
				fmt.Println(event)
			}

			// fmt.Println(event)
		}
	}()
	app.Main()
}

type (
	// C quick alias for Context.
	C = layout.Context
	// D quick alias for Dimensions.
	D = layout.Dimensions
)

type Page struct {
	inputAlignment                                               layout.Alignment
	inputAlignmentEnum                                           widget.Enum
	username, password                                           component.TextField
	nameInput, addressInput, priceInput, tweetInput, numberInput component.TextField
	widget.List
}

var (
	th       = material.NewTheme(gofont.Collection())
	notifier = func() notify.Notifier {
		n, err := notify.NewNotifier()
		if err != nil {
			panic(fmt.Errorf("init notification manager: %w", err))
		}
		return n
	}()
	editor    component.TextField
	notifyBtn widget.Clickable
)

func (p *Page) frame(gtx C) D {
	if notifyBtn.Clicked() {
		msg := "This is a notification send from gio."
		if txt := editor.Text(); txt != "" {
			msg = txt
		}
		go notifier.CreateNotification("Esword", msg)
	}
	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(gtx,
		layout.Rigid(func(gtx C) D {
			return alo.DefaultInset.Layout(gtx, material.Body2(th, "input username").Layout)
		}),
		layout.Rigid(func(gtx C) D {
			p.username.Alignment = p.inputAlignment
			return p.username.Layout(gtx, th, "UserName")
		}),
		layout.Rigid(func(gtx C) D {
			return alo.DefaultInset.Layout(gtx, material.Body2(th, "input password").Layout)
		}),
		layout.Rigid(func(gtx C) D {
			p.password.Alignment = p.inputAlignment
			return p.password.Layout(gtx, th, "Password")
		}),
		layout.Rigid(func(gtx C) D {
			// info := fmt.Sprintf("um:%S,pwd:%S",p.username,p.password)
			gtx.Constraints.Max.X = gtx.Dp(unit.Dp(300))
			return material.Button(th, &notifyBtn, "notify").Layout(gtx)
		}),
	)
	// layout.Center.Layout(gtx, func(gtx C) D {
	// 	gtx.Constraints.Max.X = gtx.Dp(unit.Dp(300))
	// 	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
	// 		layout.Rigid(func(gtx C) D {
	// 			return editor.Layout(gtx, th, "enter a notification message1")
	// 		}),
	// 		layout.Rigid(func(gtx C) D {
	// 			return layout.Spacer{Height: unit.Dp(10)}.Layout(gtx)
	// 		}),
	// 		layout.Rigid(func(gtx C) D {
	// 			return material.Button(th, &notifyBtn, "notify").Layout(gtx)
	// 		}),
	// 	)
	// })

	// layout.Center.Layout(gtx, func(gtx C) D {
	// 	gtx.Constraints.Max.X = gtx.Dp(unit.Dp(300))
	// 	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
	// 		// layout.Rigid(func(gtx C) D {
	// 		// 	return editor.Layout(gtx, th, "enter a notification message2")
	// 		// }),
	// 		// layout.Rigid(func(gtx C) D {
	// 		// 	return layout.Spacer{Height: unit.Dp(10)}.Layout(gtx)
	// 		// }),
	// 		layout.Rigid(func(gtx C) D {
	// 			return material.Button(th, &notifyBtn, "notify").Layout(gtx)
	// 		}),
	// 	)
	// })
	// return gtx.Ops
}
