package series

import (
	"context"
	"fmt"
	"os"
	"time"

	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/go-resty/resty/v2"

	"github.com/trevinteacutter/mwo-helper/pkg/helper/settings"
	"github.com/trevinteacutter/mwo-helper/pkg/mwo/api"
)

type SeriesDetails struct {
	TeamA   string
	TeamB   string
	Matches []MatchDetails
}

type MatchDetails struct {
	Details api.MatchResponse
	Mapping map[string]string
}

type Input struct {
	create  *widget.Clickable
	teamA   *widget.Editor
	teamB   *widget.Editor
	team1   *widget.Enum
	apiKey  *widget.Editor
	matchID *widget.Editor
	add     *widget.Clickable
	client  *resty.Client

	series  chan<- SeriesDetails
	matches chan<- MatchDetails
}

func NewInput(series chan<- SeriesDetails, matches chan<- MatchDetails) *Input {
	return &Input{
		create: new(widget.Clickable),
		teamA: &widget.Editor{
			Alignment:  text.Start,
			SingleLine: true,
		},
		teamB: &widget.Editor{
			Alignment:  text.Start,
			SingleLine: true,
		},
		team1: &widget.Enum{},
		apiKey: &widget.Editor{
			Alignment:  text.Start,
			SingleLine: true,
			Mask:       '*',
			Filter:     "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		},
		matchID: &widget.Editor{
			Alignment:  text.Start,
			SingleLine: true,
			Filter:     "0123456789",
		},
		add:     new(widget.Clickable),
		client:  resty.New(),
		series:  series,
		matches: matches,
	}
}

func (i *Input) Layout(gtx layout.Context, theme *material.Theme) layout.Dimensions {
	return layout.Flex{
		Alignment: layout.Middle,
		Axis:      layout.Vertical,
	}.Layout(gtx,
		i.createRow(theme),
		i.modifyRow(theme),
	)
}

func (i *Input) createRow(theme *material.Theme) layout.FlexChild {
	return layout.Rigid(
		func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{
				Alignment: layout.Middle,
				Axis:      layout.Horizontal,
			}.Layout(
				gtx,
				layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
					return layout.UniformInset(2).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return material.Editor(theme, i.teamA, "Team A").Layout(gtx)
					})
				}),
				layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
					return layout.UniformInset(2).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return material.Editor(theme, i.teamB, "Team B").Layout(gtx)
					})
				}),
				layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
					return layout.UniformInset(2).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						if i.create.Clicked(gtx) {
							series := SeriesDetails{
								TeamA:   i.teamA.Text(),
								TeamB:   i.teamB.Text(),
								Matches: make([]MatchDetails, 0),
							}

							select {
							case i.series <- series:
							default:
							}

							for i.create.Clicked(gtx) {
							}
						}

						return material.Button(theme, i.create, "Create").Layout(gtx)
					})
				}),
			)
		},
	)
}

func (i *Input) modifyRow(theme *material.Theme) layout.FlexChild {
	return layout.Rigid(
		func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{
				Alignment: layout.Middle,
				Axis:      layout.Horizontal,
			}.Layout(
				gtx,
				layout.Rigid(material.Label(theme, unit.Sp(15), "Team 1").Layout),
				layout.Rigid(material.RadioButton(theme, i.team1, i.teamA.Text(), i.teamA.Text()).Layout),
				layout.Rigid(material.RadioButton(theme, i.team1, i.teamB.Text(), i.teamB.Text()).Layout),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.UniformInset(2).Layout(gtx, material.Editor(theme, i.matchID, "Match ID").Layout)
				}),
				layout.Rigid(func(gtx layout.Context) layout.Dimensions {
					return layout.UniformInset(2).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						if i.add.Clicked(gtx) {
							results, err := i.updateDetails()
							if err != nil {
								_, _ = fmt.Fprintln(os.Stderr, err)
							}

							match := MatchDetails{
								Mapping: map[string]string{
									i.team1.Value: "1",
								},
								Details: results,
							}

							switch i.team1.Value {
							case i.teamA.Text():
								match.Mapping[i.teamB.Text()] = "2"
							default:
								match.Mapping[i.teamA.Text()] = "2"
							}

							select {
							case i.matches <- match:
							default:
							}

							for i.add.Clicked(gtx) {
							}
						}

						return material.Button(theme, i.add, "Add").Layout(gtx)
					})
				}),
			)
		},
	)
}

func (i *Input) updateDetails() (api.MatchResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	return api.Match(ctx, i.client, settings.Get().APIKey, i.matchID.Text())
}
