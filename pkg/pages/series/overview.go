package series

import (
	"image/color"
	"strconv"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
)

var headingText = []string{"Team", "Wins", "Losses", "Kills", "Deaths", "Match Score", "Damage", "Result"}

type Overview struct {
	TeamA      *TeamOverview
	TeamB      *TeamOverview
	cellBorder widget.Border
	cellInset  layout.Inset
	grid       component.GridState
	series     *SeriesDetails
}

func NewOverview(series *SeriesDetails) *Overview {
	return &Overview{
		TeamA: NewTeamOverview(),
		TeamB: NewTeamOverview(),
		cellBorder: widget.Border{
			Color: color.NRGBA{A: 255},
			Width: unit.Dp(1),
		},
		cellInset: layout.UniformInset(unit.Dp(2)),
		series:    series,
	}
}

func (o *Overview) Layout(gtx layout.Context, theme *material.Theme) layout.Dimensions {
	return layout.Flex{
		Alignment: layout.Middle,
		Axis:      layout.Horizontal,
	}.Layout(
		gtx,
		o.Table(theme),
	)
	// var children []layout.FlexChild
	//
	// if o.series != nil {
	// 	children = append(
	// 		children,
	// 		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
	// 			return o.TeamA.Layout(gtx, theme, o.series.TeamA)
	// 		}),
	// 		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
	// 			return o.TeamB.Layout(gtx, theme, o.series.TeamB)
	// 		}),
	// 	)
	// }
	//
	// return layout.Flex{
	// 	Alignment: layout.Middle,
	// 	Axis:      layout.Horizontal,
	// }.Layout(gtx, children...)
}

func (o *Overview) Table(theme *material.Theme) layout.FlexChild {
	// Configure a label styled to be a heading.
	headingLabel := material.Body2(theme, "")
	headingLabel.Font.Weight = font.Bold
	headingLabel.Alignment = text.Middle
	headingLabel.MaxLines = 1

	// Configure a label styled to be a data element.
	dataLabel := material.Body2(theme, "")
	dataLabel.Font.Typeface = "Go Mono"
	dataLabel.MaxLines = 1
	dataLabel.Alignment = text.Start

	return layout.Flexed(float32(1*len(headingText)), func(gtx layout.Context) layout.Dimensions {
		if o.series == nil {
			return layout.Dimensions{}
		}

		rows := [][]string{
			o.BuildRow(o.series.TeamA),
			o.BuildRow(o.series.TeamB),
		}

		return component.Table(theme, &o.grid).Layout(gtx, 2, len(headingText),
			func(axis layout.Axis, index, constraint int) int {
				switch axis {
				case layout.Horizontal:
					return constraint / len(headingText)
				case layout.Vertical:
					return constraint / 3
				default:
					return constraint
				}
			},
			func(gtx layout.Context, col int) layout.Dimensions {
				headingLabel.Text = headingText[col]

				return o.StyleCell(gtx, headingLabel)
			},
			func(gtx layout.Context, row, col int) layout.Dimensions {
				dataLabel.Text = rows[row][col]

				return o.StyleCell(gtx, dataLabel)
			},
		)
	})
}

func (o *Overview) BuildRow(team string) []string {
	var wins, losses, kills, deaths, matchScore, damage int

	for _, match := range o.series.Matches {
		actualTeam := match.Mapping[team]

		switch match.Details.MatchDetails.WinningTeam {
		case actualTeam:
			wins++
		default:
			losses++
		}

		for _, user := range match.Details.UserDetails {
			if user.Team == actualTeam {
				kills += user.Kills
				matchScore += user.MatchScore
				damage += user.Damage

				if user.HealthPercentage <= 0 {
					deaths++
				}
			}
		}
	}

	var result string

	if wins > losses {
		result = "🏆"
	}

	return []string{
		team,
		strconv.Itoa(wins),
		strconv.Itoa(losses),
		strconv.Itoa(kills),
		strconv.Itoa(deaths),
		strconv.Itoa(matchScore),
		strconv.Itoa(damage),
		result,
	}
}

func (o *Overview) StyleCell(gtx layout.Context, l material.LabelStyle) layout.Dimensions {
	return o.cellBorder.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return o.cellInset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return l.Layout(gtx)
		})
	})
}

type TeamOverview struct{}

func NewTeamOverview() *TeamOverview {
	return &TeamOverview{}
}

func (t *TeamOverview) Layout(gtx layout.Context, theme *material.Theme, team string) layout.Dimensions {
	return layout.Flex{
		Alignment: layout.Middle,
		Axis:      layout.Vertical,
	}.Layout(gtx,
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return material.Label(theme, unit.Sp(15), team).Layout(gtx)
		}),
	)
}
