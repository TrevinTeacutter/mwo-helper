package series

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
	"gioui.org/x/component"

	"github.com/trevinteacutter/mwo-helper/pkg/icon"
	"github.com/trevinteacutter/mwo-helper/pkg/pages"
)

// Page holds the state for a page demonstrating the features of
// the AppBar component.
type Page struct {
	input *Input
	tabs  *TabComponent

	series  chan SeriesDetails
	matches chan MatchDetails
}

// New constructs a Page with the provided router.
func New() *Page {
	series := make(chan SeriesDetails, 1)
	matches := make(chan MatchDetails, 1)

	return &Page{
		input:   NewInput(series, matches),
		tabs:    NewTabComponent(series, matches),
		series:  series,
		matches: matches,
	}
}

var _ pages.Page = &Page{}

func (p *Page) Actions() []component.AppBarAction {
	return []component.AppBarAction{}
}

func (p *Page) Overflow() []component.OverflowAction {
	return []component.OverflowAction{}
}

func (p *Page) NavItem() component.NavItem {
	return component.NavItem{
		Name: "Series",
		Icon: icon.SeriesIcon,
	}
}

func (p *Page) Layout(gtx layout.Context, theme *material.Theme) layout.Dimensions {
	return layout.Flex{
		Alignment: layout.Middle,
		Axis:      layout.Vertical,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return p.input.Layout(gtx, theme)
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return p.tabs.Layout(gtx, theme)
		}),
	)
}
