package matches

import (
	"image/color"
	"strconv"
	"strings"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"

	"github.com/trevinteacutter/mwo-helper/pkg/mwo/api"
)

var scoreboardHeaders = []string{"Lance", "Player", "Mech", "Health", "Match Score", "Damage", "Kills", "Assists", "KMDD", "Components", "Team Damage"}
var overviewHeaders = []string{"Tonnage", "Lights", "Mediums", "Heavies", "Assaults", "Heroes", "Legends", "Clan", "IS", "Dupes", "Validity"}

type TeamScoreboard struct {
	cellBorder widget.Border
	cellInset  layout.Inset
	grid       component.GridState
}

func NewTeamScoreboard() *TeamScoreboard {
	return &TeamScoreboard{
		cellBorder: widget.Border{
			Color: color.NRGBA{A: 255},
			Width: unit.Dp(1),
		},
		cellInset: layout.UniformInset(unit.Dp(2)),
	}
}

func (s *TeamScoreboard) Layout(gtx layout.Context, theme *material.Theme, score int, winner bool, valid bool, users []api.UserDetails) layout.Dimensions {
	return layout.Flex{
		Alignment: layout.Middle,
		Axis:      layout.Horizontal,
	}.Layout(
		gtx,
		s.ScoreColumn(theme, score, len(users), winner),
		layout.Flexed(float32(len(scoreboardHeaders)), func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{
				Alignment: layout.Middle,
				Axis:      layout.Vertical,
			}.Layout(
				gtx,
				s.Table(theme, users),
				s.OverviewTable(theme, valid, users),
			)
		}),
	)
}

func (s *TeamScoreboard) ScoreColumn(theme *material.Theme, score int, playerCount int, winner bool) layout.FlexChild {
	// Configure a label styled to be a heading.
	headingLabel := material.Body2(theme, "")
	headingLabel.Font.Weight = font.Bold
	headingLabel.Alignment = text.Middle
	headingLabel.MaxLines = 1

	// Configure a label styled to be a data element.
	dataLabel := material.Body2(theme, "")
	dataLabel.Font.Typeface = "Go Mono"
	dataLabel.MaxLines = 1
	dataLabel.Alignment = text.Middle

	return layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
		return component.Table(theme, &s.grid).Layout(gtx, 1, 1,
			func(axis layout.Axis, index, constraint int) int {
				switch axis {
				case layout.Vertical:
					return constraint / (playerCount + 1)
				default:
					return constraint
				}
			},
			func(gtx layout.Context, col int) layout.Dimensions {
				headingLabel.Text = "Score"

				return s.StyleCell(gtx, headingLabel.Layout)
			},
			func(gtx layout.Context, row, col int) layout.Dimensions {
				switch col {
				case 0:
					dataLabel.Text = strconv.Itoa(score)
					dataLabel.MaxLines = 1

					if winner {
						dataLabel.MaxLines = 2
						dataLabel.Text += "🏆"
					}
				}

				return s.StyleCell(gtx, dataLabel.Layout)
			},
		)
	})
}

func (s *TeamScoreboard) Table(theme *material.Theme, users []api.UserDetails) layout.FlexChild {
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

	return layout.Flexed(float32(len(users)+1), func(gtx layout.Context) layout.Dimensions {
		return component.Table(theme, &s.grid).Layout(gtx, len(users), len(scoreboardHeaders),
			func(axis layout.Axis, index, constraint int) int {
				switch axis {
				case layout.Horizontal:
					return constraint / len(scoreboardHeaders)
				case layout.Vertical:
					return constraint / (len(users) + 1)
				default:
					return constraint
				}
			},
			func(gtx layout.Context, col int) layout.Dimensions {
				headingLabel.Text = scoreboardHeaders[col]

				return s.StyleCell(gtx, headingLabel.Layout)
			},
			func(gtx layout.Context, row, col int) layout.Dimensions {
				player := users[row]
				switch col {
				case 0:
					dataLabel.Text = player.Lance
				case 1:
					dataLabel.Text = player.Username
				case 2:
					dataLabel.Text = strings.ToUpper(player.MechName)
				case 3:
					dataLabel.Text = s.HealthText(player.HealthPercentage)
				case 4:
					dataLabel.Text = strconv.Itoa(player.MatchScore)
				case 5:
					dataLabel.Text = strconv.Itoa(player.Damage)
				case 6:
					dataLabel.Text = strconv.Itoa(player.Kills)
				case 7:
					dataLabel.Text = strconv.Itoa(player.Assists)
				case 8:
					dataLabel.Text = strconv.Itoa(player.KillsMostDamage)
				case 9:
					dataLabel.Text = strconv.Itoa(player.ComponentsDestroyed)
				case 10:
					dataLabel.Text = strconv.Itoa(player.TeamDamage)
				}

				return s.StyleCell(gtx, dataLabel.Layout)
			},
		)
	})
}

func (s *TeamScoreboard) OverviewTable(theme *material.Theme, valid bool, users []api.UserDetails) layout.FlexChild {
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

	tonnage := 0
	lights := 0
	mediums := 0
	heavies := 0
	assaults := 0
	heroes := 0
	legends := 0
	clan := 0
	is := 0
	dupes := 0
	variants := make(map[string]int)

	for _, user := range users {
		mech := api.VariantFromCode(user.MechItemID)
		variant := mech.Name

		if mech.Alias != "" {
			variant = mech.Alias
		}

		variants[variant]++
		tonnage += mech.Weight

		switch strings.ToUpper(mech.WeightClass) {
		case "LIGHT":
			lights++
		case "MEDIUM":
			mediums++
		case "HEAVY":
			heavies++
		case "ASSAULT":
			assaults++
		}

		switch strings.ToUpper(mech.TechBase) {
		case "CLAN":
			clan++
		case "INNERSPHERE":
			is++
		}

		for _, tag := range mech.Tags {
			switch strings.ToUpper(tag) {
			case "HERO":
				heroes++
			case "LEGEND":
				legends++
			}
		}
	}

	for _, count := range variants {
		if count > 1 {
			dupes++
		}
	}

	return layout.Flexed(2.0, func(gtx layout.Context) layout.Dimensions {
		return component.Table(theme, &s.grid).Layout(gtx, 1, len(overviewHeaders),
			func(axis layout.Axis, index, constraint int) int {
				switch axis {
				case layout.Horizontal:
					return constraint / len(overviewHeaders)
				case layout.Vertical:
					return constraint / 2
				default:
					return constraint
				}
			},
			func(gtx layout.Context, col int) layout.Dimensions {
				headingLabel.Text = overviewHeaders[col]

				return s.StyleCell(gtx, headingLabel.Layout)
			},
			func(gtx layout.Context, _, col int) layout.Dimensions {
				switch col {
				case 0:
					dataLabel.Text = strconv.Itoa(tonnage)
				case 1:
					dataLabel.Text = strconv.Itoa(lights)
				case 2:
					dataLabel.Text = strconv.Itoa(mediums)
				case 3:
					dataLabel.Text = strconv.Itoa(heavies)
				case 4:
					dataLabel.Text = strconv.Itoa(assaults)
				case 5:
					dataLabel.Text = strconv.Itoa(heroes)
				case 6:
					dataLabel.Text = strconv.Itoa(legends)
				case 7:
					dataLabel.Text = strconv.Itoa(clan)
				case 8:
					dataLabel.Text = strconv.Itoa(is)
				case 9:
					dataLabel.Text = strconv.Itoa(dupes)
				case 10:
					if valid {
						dataLabel.Text = "VALID"
					} else {
						dataLabel.Text = "INVALID"
					}
				}

				return s.StyleCell(gtx, dataLabel.Layout)
			},
		)
	})
}

func (s *TeamScoreboard) StyleCell(gtx layout.Context, widget layout.Widget) layout.Dimensions {
	return s.cellBorder.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return s.cellInset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return widget(gtx)
		})
	})
}

func (s *TeamScoreboard) HealthText(value int) string {
	if value == 0 {
		return "DEAD"
	}

	return strconv.Itoa(value) + "%"
}
