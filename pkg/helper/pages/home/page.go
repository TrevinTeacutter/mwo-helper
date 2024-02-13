package home

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
	"gioui.org/x/component"

	"github.com/trevinteacutter/mwo-helper/pkg/helper/pages"
	"github.com/trevinteacutter/mwo-helper/pkg/icon"
)

var _ pages.Page = (*Page)(nil)

// Page holds the state for a page demonstrating the features of
// the AppBar component.
type Page struct{}

// New constructs a Page with the provided router.
func New() *Page {
	return &Page{}
}

func (p *Page) Actions() []component.AppBarAction {
	return []component.AppBarAction{}
}

func (p *Page) Overflow() []component.OverflowAction {
	return []component.OverflowAction{}
}

func (p *Page) NavItem() component.NavItem {
	return component.NavItem{
		Name: "Home",
		Icon: icon.HomeIcon,
	}
}

func (p *Page) Layout(gtx layout.Context, theme *material.Theme) layout.Dimensions {
	return layout.Flex{
		Spacing:   layout.SpaceBetween,
		Alignment: layout.Middle,
		Axis:      layout.Vertical,
	}.Layout(gtx)
}
