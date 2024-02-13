package helper

import (
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget/material"

	"github.com/trevinteacutter/mwo-helper/pkg/prototype"
)

func Loop(window *app.Window) error {
	// database := Load()
	//
	// defer database.Close()

	theme := material.NewTheme()
	theme.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))

	var ops op.Ops

	ui := prototype.New()

	for {
		switch event := window.NextEvent().(type) {
		case system.DestroyEvent:
			return event.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, event)

			ui.Layout(gtx, theme)
			event.Frame(gtx.Ops)
		}
	}
}
