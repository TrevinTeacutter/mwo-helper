package matches

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
	"github.com/trevinteacutter/mwo-helper/pkg/settings"
)

type Input struct {
	apiKey  *widget.Editor
	matchID *widget.Editor
	submit  *widget.Clickable
	client  *resty.Client

	matches chan<- api.MatchResponse
}

func NewInput(matches chan<- api.MatchResponse) *Input {
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
		submit:  new(widget.Clickable),
		client:  resty.New(),
		matches: matches,
	}
}

func (i *Input) Layout(gtx layout.Context, theme *material.Theme) layout.Dimensions {
	return layout.Flex{
		Alignment: layout.Middle,
		Axis:      layout.Horizontal,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.UniformInset(2).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return material.Editor(theme, i.matchID, "Match ID").Layout(gtx)
			})
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
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

func (i *Input) updateDetails() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	results, err := api.Match(ctx, i.client, settings.Get().APIKey, i.matchID.Text())
	if err != nil {
		return err
	}

	select {
	case i.matches <- results:
	}

	return nil
}
