package home

import (
	"slices"

	"gioui.org/layout"
	"gioui.org/widget/material"
	"gioui.org/x/component"

	"github.com/trevinteacutter/mwo-helper/pkg/icon"
	"github.com/trevinteacutter/mwo-helper/pkg/mwo/api"
	"github.com/trevinteacutter/mwo-helper/pkg/pages"
)

// Page holds the state for a page demonstrating the features of
// the AppBar component.
type Page struct {
	input    *Input
	overview *Overview
	team1    *Scoreboard
	team2    *Scoreboard
}

// New constructs a Page with the provided router.
func New() *Page {
	p := &Page{
		input:    NewInput(),
		overview: NewOverview(),
		team1:    NewScoreboard(),
		team2:    NewScoreboard(),
	}

	return p
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
		Name: "Match Lookup",
		Icon: icon.HomeIcon,
	}
}

func (p *Page) Layout(gtx layout.Context, theme *material.Theme) layout.Dimensions {
	team1 := make([]api.UserDetails, 0, 12)
	team2 := make([]api.UserDetails, 0, 12)
	results := p.input.Results()

	for _, user := range results.UserDetails {
		if user.IsSpectator {
			continue
		}

		switch user.Team {
		case "1":
			team1 = append(team1, user)
		case "2":
			team2 = append(team2, user)
		default:
		}
	}

	slices.SortFunc(team1, sortByLance)
	slices.SortFunc(team2, sortByLance)

	return layout.Flex{
		Alignment: layout.Middle,
		Axis:      layout.Vertical,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return p.input.Layout(gtx, theme)
		}),
		layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
			return p.overview.Layout(gtx, theme, results.MatchDetails)
		}),
		layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
			return p.team1.Layout(gtx, theme, results.MatchDetails.Team1Score, results.MatchDetails.WinningTeam == "1", team1)
		}),
		layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
			return p.team2.Layout(gtx, theme, results.MatchDetails.Team2Score, results.MatchDetails.WinningTeam == "2", team2)
		}),
	)
}

func sortByLance(a, b api.UserDetails) int {
	switch {
	case a.Lance < b.Lance:
		return -1
	case a.Lance > b.Lance:
		return 1
	default:
		return 0
	}
}
