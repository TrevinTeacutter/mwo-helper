package helper

import (
	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget/material"

	"github.com/trevinteacutter/mwo-helper/pkg/pages"
	"github.com/trevinteacutter/mwo-helper/pkg/pages/about"
	"github.com/trevinteacutter/mwo-helper/pkg/pages/home"
	"github.com/trevinteacutter/mwo-helper/pkg/pages/matches"
	"github.com/trevinteacutter/mwo-helper/pkg/pages/series"
	"github.com/trevinteacutter/mwo-helper/pkg/pages/settings"
)

func Loop(window *app.Window) error {
	// database := Load()
	//
	// defer database.Close()

	theme := material.NewTheme()
	theme.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))

	var ops op.Ops

	router := pages.NewRouter()
	router.Register(0, home.New())
	// router.Register(1, teams.New())
	// router.Register(2, pilots.New())
	router.Register(1, series.New())
	router.Register(2, matches.New())
	// router.Register(5, isc.New())
	// router.Register(6, wc.New())
	router.Register(3, settings.New())
	router.Register(4, about.New())

	for {
		switch event := window.NextEvent().(type) {
		case system.DestroyEvent:
			return event.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, event)

			router.Layout(gtx, theme)
			event.Frame(gtx.Ops)
		}
	}
}
