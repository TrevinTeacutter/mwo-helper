package series

import (
	"sort"
	"strconv"
	"strings"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
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
	chassis *ChassisList
	variant *VariantList
}

func NewTeamMechSummary() *TeamMechSummary {
	return &TeamMechSummary{
		chassis: NewChassisList(),
		variant: NewVariantList(),
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
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return t.chassis.Layout(gtx, theme, team, series)
				}),
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return t.variant.Layout(gtx, theme, team, series)
				}),
			)
		}),
	)
}

type ChassisList struct {
}

func NewChassisList() *ChassisList {
	return &ChassisList{}
}

func (c *ChassisList) Layout(gtx layout.Context, theme *material.Theme, team string, series *SeriesDetails) layout.Dimensions {
	chassiss := make(map[string]int)

	for _, match := range series.Matches {
		if match.Mapping == nil {
			continue
		}

		for _, player := range match.Details.UserDetails {
			if match.Mapping[team] == player.Team {
				chassiss[variantToChassis(player.MechName)]++
			}
		}
	}

	keys := make([]string, 0)

	for key := range chassiss {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	children := make([]layout.FlexChild, 0, len(chassiss)+1)

	children = append(
		children,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return material.Label(theme, unit.Sp(15), "Chassis Usage").Layout(gtx)
		}),
	)

	for _, key := range keys {
		chassis := key
		count := chassiss[key]
		children = append(
			children,
			layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{
					Alignment: layout.Middle,
					Axis:      layout.Horizontal,
				}.Layout(
					gtx,
					layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
						return material.Label(theme, unit.Sp(15), chassis).Layout(gtx)
					}),
					layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
						return material.Label(theme, unit.Sp(15), strconv.Itoa(count)).Layout(gtx)
					}),
				)
			}),
		)
	}

	return layout.Flex{
		Alignment: layout.Middle,
		Axis:      layout.Vertical,
	}.Layout(gtx, children...)
}

type VariantList struct {
}

func NewVariantList() *VariantList {
	return &VariantList{}
}

func (vl *VariantList) Layout(gtx layout.Context, theme *material.Theme, team string, series *SeriesDetails) layout.Dimensions {
	variants := make(map[string]int)

	for _, match := range series.Matches {
		if match.Mapping == nil {
			continue
		}

		for _, player := range match.Details.UserDetails {
			if match.Mapping[team] == player.Team {
				variants[player.MechName]++
			}
		}
	}

	keys := make([]string, 0)

	for key := range variants {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	children := make([]layout.FlexChild, 0, len(variants)+1)

	children = append(
		children,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return material.Label(theme, unit.Sp(15), "Variant Usage").Layout(gtx)
		}),
	)

	for _, key := range keys {
		variant := key
		count := variants[key]
		children = append(
			children,
			layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{
					Alignment: layout.Middle,
					Axis:      layout.Horizontal,
				}.Layout(
					gtx,
					layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
						return material.Label(theme, unit.Sp(15), strings.ToUpper(variantDedupe(variant))).Layout(gtx)
					}),
					layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
						return material.Label(theme, unit.Sp(15), strconv.Itoa(count)).Layout(gtx)
					}),
				)
			}),
		)
	}

	return layout.Flex{
		Alignment: layout.Middle,
		Axis:      layout.Vertical,
	}.Layout(gtx, children...)
}
