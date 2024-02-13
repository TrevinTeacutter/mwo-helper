package matches

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
	"gioui.org/x/component"

	"github.com/trevinteacutter/mwo-helper/pkg/helper/pages"
	"github.com/trevinteacutter/mwo-helper/pkg/icon"
	"github.com/trevinteacutter/mwo-helper/pkg/mwo/api"
)

var _ pages.Page = (*Page)(nil)

// Page holds the state for a page demonstrating the features of
// the AppBar component.
type Page struct {
	input      *Input
	scoreboard *Scoreboard
	matches    chan api.MatchResponse
}

// New constructs a Page with the provided router.
func New() *Page {
	matches := make(chan api.MatchResponse, 1)

	return &Page{
		input:      NewInput(matches),
		scoreboard: NewScoreboard(matches),
		matches:    matches,
	}
}

func (p *Page) Actions() []component.AppBarAction {
	return []component.AppBarAction{}
}

func (p *Page) Overflow() []component.OverflowAction {
	return []component.OverflowAction{}
}

func (p *Page) NavItem() component.NavItem {
	return component.NavItem{
		Name: "Matches",
		Icon: icon.MatchIcon,
	}
}

func (p *Page) Layout(gtx layout.Context, theme *material.Theme) layout.Dimensions {
	return layout.Flex{
		Spacing:   layout.SpaceBetween,
		Alignment: layout.Middle,
		Axis:      layout.Vertical,
	}.Layout(gtx,
		layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
			return p.input.Layout(gtx, theme)
		}),
		layout.Flexed(5, func(gtx layout.Context) layout.Dimensions {
			return p.scoreboard.Layout(gtx, theme)
		}),
	)
}
