package matches

import (
	"time"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"

	"github.com/trevinteacutter/mwo-helper/pkg/mwo/api"
)

type Overview struct {
}

func NewOverview() *Overview {
	return &Overview{}
}

func (o *Overview) Layout(gtx layout.Context, theme *material.Theme, details api.MatchDetails) layout.Dimensions {
	return layout.Flex{
		Alignment: layout.Middle,
		Axis:      layout.Vertical,
	}.Layout(gtx,
		layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{
				Alignment: layout.Middle,
				Axis:      layout.Horizontal,
			}.Layout(
				gtx,
				layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
					return material.Label(theme, unit.Sp(15), "Completed At: "+details.CompleteTime.Format(time.RFC3339)).Layout(gtx)
				}),
				layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
					return material.Label(theme, unit.Sp(15), "Region: "+details.Region).Layout(gtx)
				}),
				layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
					return material.Label(theme, unit.Sp(15), "Duration: "+o.Seconds(details.MatchDuration)).Layout(gtx)
				}),
				layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
					return layout.Flex{
						Alignment: layout.Middle,
						Axis:      layout.Horizontal,
					}.Layout(
						gtx,
						layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
							switch details.ViewMode {
							case "Both", "FirstPersonOnly":
								return material.Label(theme, unit.Sp(10), "First Person Enabled").Layout(gtx)
							default:
								return material.Label(theme, unit.Sp(10), "First Person Disabled").Layout(gtx)
							}
						}),
						layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
							switch details.ViewMode {
							case "Both", "ThirdPersonOnly":
								return material.Label(theme, unit.Sp(10), "Third Person Enabled").Layout(gtx)
							default:
								return material.Label(theme, unit.Sp(10), "Third Person Disabled").Layout(gtx)
							}
						}),
						layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
							switch details.NoMechQuirks {
							case true:
								return material.Label(theme, unit.Sp(10), "Quirks Disabled").Layout(gtx)
							default:
								return material.Label(theme, unit.Sp(10), "Quirks Enabled").Layout(gtx)
							}
						}),
						layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
							switch details.NoMechEfficiencies {
							case true:
								return material.Label(theme, unit.Sp(10), "Skills Disabled").Layout(gtx)
							default:
								return material.Label(theme, unit.Sp(10), "Skills Enabled").Layout(gtx)
							}
						}),
						layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
							switch details.UseStockLoadout {
							case true:
								return material.Label(theme, unit.Sp(10), "Stock Loadouts").Layout(gtx)
							default:
								return material.Label(theme, unit.Sp(10), "Custom Loadouts").Layout(gtx)
							}
						}),
					)
				}),
			)
		}),
		layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{
				Alignment: layout.Middle,
				Axis:      layout.Horizontal,
			}.Layout(
				gtx,
				layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
					return material.Label(theme, unit.Sp(15), "Game Mode: "+details.GameMode).Layout(gtx)
				}),
				layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
					return material.Label(theme, unit.Sp(15), "Time Limit: "+o.Minutes(details.MatchTimeMinutes)).Layout(gtx)
				}),
				layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
					return material.Label(theme, unit.Sp(15), "Map: "+details.Map).Layout(gtx)
				}),
				layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
					return material.Label(theme, unit.Sp(15), "Map Time of Day: "+details.TimeOfDay).Layout(gtx)
				}),
			)
		}),
	)
}

func (o *Overview) Seconds(value int) string {
	return (time.Second * time.Duration(value)).String()
}

func (o *Overview) Minutes(value int) string {
	return (time.Minute * time.Duration(value)).String()
}
