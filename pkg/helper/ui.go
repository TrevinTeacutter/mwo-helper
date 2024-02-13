package helper

import (
	"gioui.org/layout"
	"gioui.org/widget/material"

	"github.com/trevinteacutter/mwo-helper/pkg/helper/pages"
	"github.com/trevinteacutter/mwo-helper/pkg/helper/pages/about"
	"github.com/trevinteacutter/mwo-helper/pkg/helper/pages/home"
	"github.com/trevinteacutter/mwo-helper/pkg/helper/pages/matches"
	"github.com/trevinteacutter/mwo-helper/pkg/helper/pages/series"
	"github.com/trevinteacutter/mwo-helper/pkg/helper/pages/settings"
)

type UI struct {
	router pages.Router
}

func New() *UI {
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

	return &UI{
		router: router,
	}
}

func (ui *UI) Layout(gtx layout.Context, theme *material.Theme) {
	ui.router.Layout(gtx, theme)
}
