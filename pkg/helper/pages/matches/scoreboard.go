package matches

import (
	"slices"

	"gioui.org/layout"
	"gioui.org/widget/material"

	"github.com/trevinteacutter/mwo-helper/pkg/mwo/api"
)

type Scoreboard struct {
	overview *Overview
	team1    *TeamScoreboard
	team2    *TeamScoreboard

	matches <-chan api.MatchResponse

	results api.MatchResponse
}

func NewScoreboard(matches <-chan api.MatchResponse) *Scoreboard {
	return &Scoreboard{
		overview: NewOverview(),
		team1:    NewTeamScoreboard(),
		team2:    NewTeamScoreboard(),
		matches:  matches,
	}
}

func (s *Scoreboard) Layout(gtx layout.Context, theme *material.Theme) layout.Dimensions {
	select {
	case results, ok := <-s.matches:
		if ok {
			s.results = results
		}
	default:
	}

	team1 := make([]api.UserDetails, 0, 12)
	team2 := make([]api.UserDetails, 0, 12)

	for _, user := range s.results.UserDetails {
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
		Spacing:   layout.SpaceBetween,
		Alignment: layout.Middle,
		Axis:      layout.Vertical,
	}.Layout(gtx,
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return s.overview.Layout(gtx, theme, s.results.MatchDetails)
		}),
		layout.Flexed(2, func(gtx layout.Context) layout.Dimensions {
			return s.team1.Layout(gtx, theme, s.results.MatchDetails.Team1Score, s.results.MatchDetails.WinningTeam == "1", team1)
		}),
		layout.Flexed(2, func(gtx layout.Context) layout.Dimensions {
			return s.team2.Layout(gtx, theme, s.results.MatchDetails.Team2Score, s.results.MatchDetails.WinningTeam == "2", team2)
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
