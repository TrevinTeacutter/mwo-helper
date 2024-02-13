package series

import (
	"image"
	"strconv"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"

	matchPage "github.com/trevinteacutter/mwo-helper/pkg/helper/pages/matches"
	"github.com/trevinteacutter/mwo-helper/pkg/mwo/api"
)

type Tab interface {
	Layout(gtx layout.Context, theme *material.Theme) layout.Dimensions
}

type TabComponent struct {
	list     layout.List
	tabs     []TabWrapper
	selected int
	series   <-chan SeriesDetails
	matches  <-chan MatchDetails

	current SeriesDetails
}

type TabWrapper struct {
	button *widget.Clickable
	title  string
	widget Tab
}

func NewTabComponent(series <-chan SeriesDetails, matches <-chan MatchDetails) *TabComponent {
	return &TabComponent{
		list:    layout.List{Axis: layout.Horizontal},
		tabs:    []TabWrapper{{title: "Overview", widget: NewOverview(nil), button: new(widget.Clickable)}},
		series:  series,
		matches: matches,
	}
}

func (t *TabComponent) Layout(gtx layout.Context, theme *material.Theme) layout.Dimensions {
	select {
	case series, ok := <-t.series:
		if ok {
			t.Reset(series)
		}
	case match, ok := <-t.matches:
		if ok {
			t.AddMatch(match)
		}
	default:
	}

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return t.list.Layout(
				gtx,
				len(t.tabs),
				func(gtx layout.Context, index int) layout.Dimensions {
					tab := t.tabs[index]

					if tab.button.Clicked(gtx) {
						t.selected = index
					}

					var tabWidth int

					return layout.Stack{Alignment: layout.S}.Layout(gtx,
						layout.Stacked(func(gtx layout.Context) layout.Dimensions {
							dims := material.Clickable(
								gtx,
								tab.button,
								func(gtx layout.Context) layout.Dimensions {
									header := material.H6(theme, tab.title)

									return layout.UniformInset(unit.Dp(15)).Layout(gtx,
										header.Layout,
									)
								})

							tabWidth = dims.Size.X

							return dims
						}),
						layout.Stacked(func(gtx layout.Context) layout.Dimensions {
							if t.selected != index {
								return layout.Dimensions{}
							}
							tabHeight := gtx.Dp(unit.Dp(4))
							tabRect := image.Rect(0, 0, tabWidth, tabHeight)

							paint.FillShape(
								gtx.Ops,
								theme.Palette.ContrastBg,
								clip.Rect(tabRect).Op(),
							)

							return layout.Dimensions{
								Size: image.Point{X: tabWidth, Y: tabHeight},
							}
						}),
					)
				})
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return layout.Center.Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return t.tabs[t.selected].widget.Layout(gtx, theme)
			})
		}),
	)
}

func (t *TabComponent) Reset(series SeriesDetails) {
	t.selected = 0
	t.current = series
	t.tabs = []TabWrapper{{title: "Overview", widget: NewOverview(&t.current), button: new(widget.Clickable)}}
}

func (t *TabComponent) AddMatch(match MatchDetails) {
	if len(t.current.Matches) <= 0 {
		t.tabs = append(t.tabs, TabWrapper{title: "Pilot Summary", widget: NewPilotSummary(&t.current), button: new(widget.Clickable)})
		t.tabs = append(t.tabs, TabWrapper{title: "Mech Summary", widget: NewMechSummary(&t.current), button: new(widget.Clickable)})
	}

	t.current.Matches = append(t.current.Matches, match)

	temp := make(chan api.MatchResponse, 1)
	temp <- match.Details
	close(temp)

	index := len(t.tabs) - 2

	t.tabs = append(t.tabs[:index+1], t.tabs[index:]...)
	t.tabs[index] = TabWrapper{title: "Match " + strconv.Itoa(len(t.current.Matches)), widget: matchPage.NewScoreboard(temp), button: new(widget.Clickable)}
}
