package settings

import (
	"fmt"
	"os"

	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"

	"github.com/trevinteacutter/mwo-helper/pkg/icon"
	"github.com/trevinteacutter/mwo-helper/pkg/pages"
	"github.com/trevinteacutter/mwo-helper/pkg/settings"
)

// Page holds the state for a page demonstrating the features of
// the AppBar component.
type Page struct {
	apiKey *widget.Editor
	save   *widget.Clickable
}

// New constructs a Page with the provided router.
func New() *Page {
	p := &Page{
		apiKey: &widget.Editor{
			Alignment:  text.Start,
			SingleLine: true,
			Mask:       '*',
			Filter:     "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
		},
		save: new(widget.Clickable),
	}

	p.apiKey.SetText(settings.Get().APIKey)

	return p
}

var _ pages.Page = &Page{}

func (p *Page) Actions() []component.AppBarAction {
	return []component.AppBarAction{}
}

func (p *Page) Overflow() []component.OverflowAction {
	return []component.OverflowAction{}
}

func (p *Page) NavItem() component.NavItem {
	return component.NavItem{
		Name: "Settings",
		Icon: icon.SettingsIcon,
	}
}

func (p *Page) Layout(gtx layout.Context, theme *material.Theme) layout.Dimensions {
	return layout.Flex{
		Alignment: layout.Middle,
		Axis:      layout.Vertical,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.UniformInset(2).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return material.Editor(theme, p.apiKey, "API Key").Layout(gtx)
			})
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.UniformInset(2).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				if p.save.Clicked(gtx) {
					setting := settings.Get()
					setting.APIKey = p.apiKey.Text()

					if err := settings.Save(); err != nil {
						_, _ = fmt.Fprintf(os.Stderr, "failed to save settings: %v\n", err)
					}

					for p.save.Clicked(gtx) {
					}
				}

				return material.Button(theme, p.save, "Save").Layout(gtx)
			})
		}),
	)
}
