package prototype

import (
	"gioui.org/layout"
	"gioui.org/widget/material"

	"github.com/trevinteacutter/mwo-helper/pkg/helper/pages/pilots"
	"github.com/trevinteacutter/mwo-helper/pkg/prototype/pages"
	"github.com/trevinteacutter/mwo-helper/pkg/prototype/pages/about"
	"github.com/trevinteacutter/mwo-helper/pkg/prototype/pages/encyclopedia"
	"github.com/trevinteacutter/mwo-helper/pkg/prototype/pages/home"
	"github.com/trevinteacutter/mwo-helper/pkg/prototype/pages/mechlab"
	"github.com/trevinteacutter/mwo-helper/pkg/prototype/pages/settings"
)

type UI struct {
	router pages.Router
}

func New() *UI {
	router := pages.NewRouter()

	router.Register(0, home.New())
	router.Register(1, mechlab.New())
	router.Register(2, pilots.New())
	router.Register(3, encyclopedia.New())
	router.Register(4, settings.New())
	router.Register(5, about.New())

	return &UI{
		router: router,
	}
}

func (ui *UI) Layout(gtx layout.Context, theme *material.Theme) {
	ui.router.Layout(gtx, theme)
}
