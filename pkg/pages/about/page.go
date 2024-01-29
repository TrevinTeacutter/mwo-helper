package about

import (
	"fmt"

	"gioui.org/layout"
	"gioui.org/widget/material"
	"gioui.org/x/component"

	"github.com/trevinteacutter/mwo-helper/pkg/applayout"
	"github.com/trevinteacutter/mwo-helper/pkg/build"
	"github.com/trevinteacutter/mwo-helper/pkg/icon"
	"github.com/trevinteacutter/mwo-helper/pkg/pages"
)

// Page holds the state for a page demonstrating the features of
// the AppBar component.
type Page struct {
}

// New constructs a Page with the provided router.
func New() *Page {
	return &Page{}
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
		Name: "About",
		Icon: icon.OtherIcon,
	}
}

func (p *Page) Layout(gtx layout.Context, th *material.Theme) layout.Dimensions {
	return layout.Flex{
		Alignment: layout.Middle,
		Axis:      layout.Vertical,
	}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return applayout.DefaultInset.Layout(gtx, material.Body1(th, `This library implements material design components from https://material.io using https://gioui.org.`).Layout)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return applayout.DefaultInset.Layout(gtx, material.Body1(th, fmt.Sprintf("Version: %s", build.Build)).Layout)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return applayout.DefaultInset.Layout(gtx, material.Body1(th, fmt.Sprintf("Commit: %s", build.Commit)).Layout)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return applayout.DefaultInset.Layout(gtx, material.Body1(th, fmt.Sprintf("Time: %s", build.Date)).Layout)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return applayout.DefaultInset.Layout(gtx, material.Body1(th, fmt.Sprintf("Runtime: %s", build.Runtime)).Layout)
		}),
	)
}
