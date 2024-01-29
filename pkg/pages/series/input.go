package series

import (
	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/widget"
	"gioui.org/widget/material"

	"github.com/trevinteacutter/mwo-helper/pkg/mwo/api"
)

type Input struct {
	selector *widget.Clickable
	create   *widget.Clickable
	matchID  *widget.Editor
	add      *widget.Clickable

	current []api.MatchResponse
}

func NewInput() *Input {
	return &Input{
		selector: new(widget.Clickable),
		create:   new(widget.Clickable),
		matchID: &widget.Editor{
			Alignment:  text.Start,
			SingleLine: true,
			Filter:     "0123456789",
		},
		add: new(widget.Clickable),
	}
}

func (i *Input) Layout(gtx layout.Context, theme *material.Theme) layout.Dimensions {
	return layout.Flex{
		Alignment: layout.Middle,
		Axis:      layout.Horizontal,
	}.Layout(gtx,
		layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
			return layout.UniformInset(2).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				if i.selector.Clicked(gtx) {
					// launch model for selecting series

					for i.selector.Clicked(gtx) {
					}
				}
				return material.Button(theme, i.selector, "Select").Layout(gtx)
			})
		}),
		layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
			return layout.UniformInset(2).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				if i.add.Clicked(gtx) {
					// create new series

					for i.add.Clicked(gtx) {
					}
				}
				return material.Button(theme, i.add, "Add").Layout(gtx)
			})
		}),
		layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
			return layout.UniformInset(2).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return material.Editor(theme, i.matchID, "Match ID").Layout(gtx)
			})
		}),
		layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
			return layout.UniformInset(2).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				if i.add.Clicked(gtx) {
					// associate match with series

					for i.add.Clicked(gtx) {
					}
				}
				return material.Button(theme, i.add, "Add").Layout(gtx)
			})
		}),
	)
}

func (i *Input) Series() []api.MatchResponse {
	return i.current
}
