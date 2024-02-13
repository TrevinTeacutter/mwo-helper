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
	"golang.org/x/exp/slices"
)

var pilotSummerHeadingText = []string{"Team", "Player", "WLR", "Match Score/R", "Damage/R", "Kills/R", "Assists/R", "Deaths/R", "KMDD/R", "Components/R", "Team Damage/R"}

type PilotSummary struct {
	cellBorder widget.Border
	cellInset  layout.Inset
	grid       component.GridState
	series     *SeriesDetails
}

func NewPilotSummary(series *SeriesDetails) *PilotSummary {
	return &PilotSummary{
		cellBorder: widget.Border{
			Color: color.NRGBA{A: 255},
			Width: unit.Dp(1),
		},
		cellInset: layout.UniformInset(unit.Dp(2)),
		series:    series,
	}
}

func (p *PilotSummary) Layout(gtx layout.Context, theme *material.Theme) layout.Dimensions {
	return layout.Flex{
		Alignment: layout.Middle,
		Axis:      layout.Horizontal,
	}.Layout(
		gtx,
		p.Table(theme),
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

func (p *PilotSummary) Table(theme *material.Theme) layout.FlexChild {
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

	return layout.Flexed(float32(1*len(pilotSummerHeadingText)), func(gtx layout.Context) layout.Dimensions {
		if p.series == nil {
			return layout.Dimensions{}
		}

		var (
			rows   [][]string
			pilots []string
		)

		for _, match := range p.series.Matches {
			for _, pilot := range match.Details.UserDetails {
				if pilot.IsSpectator {
					continue
				}

				pilots = append(pilots, pilot.Username)
			}
		}

		slices.Sort(pilots)

		pilots = slices.Compact(pilots)

		for _, pilot := range pilots {
			rows = append(rows, p.BuildRow(pilot))
		}

		return component.Table(theme, &p.grid).Layout(gtx, len(pilots), len(pilotSummerHeadingText),
			func(axis layout.Axis, index, constraint int) int {
				switch axis {
				case layout.Horizontal:
					return constraint / len(pilotSummerHeadingText)
				case layout.Vertical:
					return constraint/len(pilots) + 1
				default:
					return constraint
				}
			},
			func(gtx layout.Context, col int) layout.Dimensions {
				headingLabel.Text = pilotSummerHeadingText[col]

				return p.StyleCell(gtx, headingLabel)
			},
			func(gtx layout.Context, row, col int) layout.Dimensions {
				dataLabel.Text = rows[row][col]

				return p.StyleCell(gtx, dataLabel)
			},
		)
	})
}

func (p *PilotSummary) BuildRow(name string) []string {
	var wins, losses, matchScore, damage, kills, deaths, assists, kmdds, components, teamDamage int

	var team string

	for _, match := range p.series.Matches {
		for _, pilot := range match.Details.UserDetails {
			if pilot.IsSpectator {
				continue
			}

			if pilot.Username != name {
				continue
			}

			if team == "" {
				for key, value := range match.Mapping {
					if value == pilot.Team {
						team = key
					}
				}
			}

			switch match.Details.MatchDetails.WinningTeam {
			case pilot.Team:
				wins++
			default:
				losses++
			}

			kills += pilot.Kills
			assists += pilot.Assists
			matchScore += pilot.MatchScore
			damage += pilot.Damage
			kmdds += pilot.KillsMostDamage
			components += pilot.ComponentsDestroyed
			teamDamage += pilot.TeamDamage

			if pilot.HealthPercentage <= 0 {
				deaths++
			}
		}
	}

	matches := float64(wins + losses)

	return []string{
		team,
		name,
		strconv.FormatFloat(float64(wins)/float64(losses), 'G', -1, 64),
		strconv.FormatFloat(float64(matchScore)/matches, 'G', -1, 64),
		strconv.FormatFloat(float64(damage)/matches, 'G', -1, 64),
		strconv.FormatFloat(float64(kills)/matches, 'G', -1, 64),
		strconv.FormatFloat(float64(assists)/matches, 'G', -1, 64),
		strconv.FormatFloat(float64(deaths)/matches, 'G', -1, 64),
		strconv.FormatFloat(float64(kmdds)/matches, 'G', -1, 64),
		strconv.FormatFloat(float64(components)/matches, 'G', -1, 64),
		strconv.FormatFloat(float64(teamDamage)/matches, 'G', -1, 64),
	}
}

func (p *PilotSummary) StyleCell(gtx layout.Context, l material.LabelStyle) layout.Dimensions {
	return p.cellBorder.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return p.cellInset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return l.Layout(gtx)
		})
	})
}
