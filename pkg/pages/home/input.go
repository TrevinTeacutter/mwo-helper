package home

import (
	"context"
	"fmt"
	"os"
	"time"

	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/go-resty/resty/v2"

	"github.com/trevinteacutter/mwo-helper/pkg/mwo/api"
)

type Input struct {
	apiKey  *widget.Editor
	matchID *widget.Editor
	submit  *widget.Clickable
	client  *resty.Client

	results api.MatchResponse
}

func NewInput() *Input {
	return &Input{
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
		submit: new(widget.Clickable),
		client: resty.New(),
	}
}

func (i *Input) Layout(gtx layout.Context, theme *material.Theme) layout.Dimensions {
	return layout.Flex{
		Alignment: layout.Middle,
		Axis:      layout.Horizontal,
	}.Layout(gtx,
		layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
			return layout.UniformInset(2).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return material.Editor(theme, i.apiKey, "API Key").Layout(gtx)
			})
		}),
		layout.Flexed(1.0, func(gtx layout.Context) layout.Dimensions {
			return layout.UniformInset(2).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return material.Editor(theme, i.matchID, "Match ID").Layout(gtx)
			})
		}),
		layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
			return layout.UniformInset(2).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				if i.submit.Clicked(gtx) {
					if err := i.updateDetails(); err != nil {
						_, _ = fmt.Fprintln(os.Stderr, err)
					}

					for i.submit.Clicked(gtx) {
					}
				}
				return material.Button(theme, i.submit, "Submit").Layout(gtx)
			})
		}),
	)
}

func (i *Input) Results() api.MatchResponse {
	return i.results
}

func (i *Input) updateDetails() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	var err error

	i.results, err = api.Match(ctx, i.client, i.apiKey.Text(), i.matchID.Text())
	if err != nil {
		return err
	}

	return nil
}
