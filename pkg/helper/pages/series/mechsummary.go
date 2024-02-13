package series

import (
	"image/color"
	"sort"
	"strconv"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"

	"github.com/trevinteacutter/mwo-helper/pkg/mwo/api"
)

var (
	chassisMechSummaryHeadingText = []string{"Chassis", "Count"}
	variantMechSummaryHeadingText = []string{"Variant", "Count"}
)

type MechSummary struct {
	TeamA  *TeamMechSummary
	TeamB  *TeamMechSummary
	series *SeriesDetails
}

func NewMechSummary(series *SeriesDetails) *MechSummary {
	return &MechSummary{
		TeamA:  NewTeamMechSummary(),
		TeamB:  NewTeamMechSummary(),
		series: series,
	}
}

func (m *MechSummary) Layout(gtx layout.Context, theme *material.Theme) layout.Dimensions {
	return layout.Flex{
		Alignment: layout.Middle,
		Axis:      layout.Horizontal,
	}.Layout(gtx,
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return m.TeamA.Layout(gtx, theme, m.series.TeamA, m.series)
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return m.TeamB.Layout(gtx, theme, m.series.TeamB, m.series)
		}),
	)
}

type TeamMechSummary struct {
	cellBorder widget.Border
	cellInset  layout.Inset
	grid       component.GridState
}

func NewTeamMechSummary() *TeamMechSummary {
	return &TeamMechSummary{
		cellBorder: widget.Border{
			Color: color.NRGBA{A: 255},
			Width: unit.Dp(1),
		},
		cellInset: layout.UniformInset(unit.Dp(2)),
	}
}

func (t *TeamMechSummary) Layout(gtx layout.Context, theme *material.Theme, team string, series *SeriesDetails) layout.Dimensions {
	return layout.Flex{
		Alignment: layout.Middle,
		Axis:      layout.Vertical,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return material.Label(theme, unit.Sp(15), team).Layout(gtx)
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{
				Alignment: layout.Middle,
				Axis:      layout.Horizontal,
			}.Layout(
				gtx,
				t.ChassisTable(theme, team, series),
				t.VariantTable(theme, team, series),
			)
		}),
	)
}

func (t *TeamMechSummary) ChassisTable(theme *material.Theme, team string, series *SeriesDetails) layout.FlexChild {
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

	chassiss := make(map[string]int)
	keys := make([]string, 0)

	for _, match := range series.Matches {
		if match.Mapping == nil {
			continue
		}

		for _, player := range match.Details.UserDetails {
			if match.Mapping[team] == player.Team {
				chassiss[api.VariantFromCode(player.MechItemID).Chassis]++
			}
		}
	}

	for key := range chassiss {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	return layout.Flexed(float32(1*len(chassisMechSummaryHeadingText)), func(gtx layout.Context) layout.Dimensions {
		return component.Table(theme, &t.grid).Layout(gtx, len(keys), len(chassisMechSummaryHeadingText),
			func(axis layout.Axis, index, constraint int) int {
				switch axis {
				case layout.Horizontal:
					return constraint / len(chassisMechSummaryHeadingText)
				case layout.Vertical:
					return constraint / (len(keys) + 1)
				default:
					return constraint
				}
			},
			func(gtx layout.Context, col int) layout.Dimensions {
				headingLabel.Text = chassisMechSummaryHeadingText[col]

				return t.StyleCell(gtx, headingLabel.Layout)
			},
			func(gtx layout.Context, row, col int) layout.Dimensions {
				chassis := keys[row]
				switch col {
				case 0:
					dataLabel.Text = chassis
				case 1:
					dataLabel.Text = strconv.Itoa(chassiss[chassis])
				}

				return t.StyleCell(gtx, dataLabel.Layout)
			},
		)
	})
}

func (t *TeamMechSummary) VariantTable(theme *material.Theme, team string, series *SeriesDetails) layout.FlexChild {
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

	variants := make(map[string]int)
	keys := make([]string, 0)

	for _, match := range series.Matches {
		if match.Mapping == nil {
			continue
		}

		for _, player := range match.Details.UserDetails {
			if match.Mapping[team] == player.Team {
				variant := api.VariantFromCode(player.MechItemID)
				name := variant.Name

				if variant.Alias != "" {
					name = variant.Alias
				}

				variants[name]++
			}
		}
	}

	for key := range variants {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	return layout.Flexed(float32(1*len(variantMechSummaryHeadingText)), func(gtx layout.Context) layout.Dimensions {
		return component.Table(theme, &t.grid).Layout(gtx, len(keys), len(variantMechSummaryHeadingText),
			func(axis layout.Axis, index, constraint int) int {
				switch axis {
				case layout.Horizontal:
					return constraint / len(variantMechSummaryHeadingText)
				case layout.Vertical:
					return constraint / (len(keys) + 1)
				default:
					return constraint
				}
			},
			func(gtx layout.Context, col int) layout.Dimensions {
				headingLabel.Text = variantMechSummaryHeadingText[col]

				return t.StyleCell(gtx, headingLabel.Layout)
			},
			func(gtx layout.Context, row, col int) layout.Dimensions {
				variant := keys[row]
				switch col {
				case 0:
					dataLabel.Text = variant
				case 1:
					dataLabel.Text = strconv.Itoa(variants[variant])
				}

				return t.StyleCell(gtx, dataLabel.Layout)
			},
		)
	})
}

func (t *TeamMechSummary) StyleCell(gtx layout.Context, widget layout.Widget) layout.Dimensions {
	return t.cellBorder.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return t.cellInset.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			return widget(gtx)
		})
	})
}
