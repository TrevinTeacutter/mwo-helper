package matches

import (
	"slices"

	"gioui.org/layout"
	"gioui.org/widget/material"
	"gioui.org/x/component"

	"github.com/trevinteacutter/mwo-helper/pkg/icon"
	"github.com/trevinteacutter/mwo-helper/pkg/mwo/api"
	"github.com/trevinteacutter/mwo-helper/pkg/pages"
	"github.com/trevinteacutter/mwo-helper/pkg/pages/home"
)

// Page holds the state for a page demonstrating the features of
// the AppBar component.
type Page struct {
	input    *Input
	overview *home.Overview
	team1    *home.Scoreboard
	team2    *home.Scoreboard
}

// New constructs a Page with the provided router.
func New() *Page {
	return &Page{
		input:    NewInput(),
		overview: home.NewOverview(),
		team1:    home.NewScoreboard(),
		team2:    home.NewScoreboard(),
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
		Name: "Matches",
		Icon: icon.MatchIcon,
	}
}

func (p *Page) Layout(gtx layout.Context, theme *material.Theme) layout.Dimensions {
	team1 := make([]api.UserDetails, 0, 12)
	team2 := make([]api.UserDetails, 0, 12)
	results := p.input.Match()

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
