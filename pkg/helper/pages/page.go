package pages

import (
	"log"
	"time"

	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/widget/material"
	"gioui.org/x/component"

	"github.com/trevinteacutter/mwo-helper/pkg/icon"
)

type Page interface {
	Actions() []component.AppBarAction
	Overflow() []component.OverflowAction
	Layout(gtx layout.Context, th *material.Theme) layout.Dimensions
	NavItem() component.NavItem
}

type Router struct {
	pages     map[any]Page
	current   any
	navDrawer *component.NavDrawer
	navAnim   component.VisibilityAnimation
	appBar    *component.AppBar
	modals    *component.ModalLayer
}

func NewRouter() Router {
	modal := component.NewModal()
	nav := component.NewNav("Navigation Drawer", "This is an example.")
	bar := component.NewAppBar(modal)
	na := component.VisibilityAnimation{
		State:    component.Invisible,
		Duration: time.Millisecond * 250,
	}

	bar.NavigationIcon = icon.MenuIcon

	return Router{
		pages:     make(map[interface{}]Page),
		modals:    modal,
		navDrawer: &nav,
		appBar:    bar,
		navAnim:   na,
	}
}

func (r *Router) Register(tag interface{}, p Page) {
	r.pages[tag] = p

	navItem := p.NavItem()

	navItem.Tag = tag

	if r.current == interface{}(nil) {
		r.current = tag
		r.appBar.Title = navItem.Name
		r.appBar.SetActions(p.Actions(), p.Overflow())
	}

	r.navDrawer.AddNavItem(navItem)
}

func (r *Router) SwitchTo(tag interface{}) {
	p, ok := r.pages[tag]
	if !ok {
		return
	}

	navItem := p.NavItem()

	r.current = tag
	r.appBar.Title = navItem.Name

	r.appBar.SetActions(p.Actions(), p.Overflow())
}

func (r *Router) Layout(gtx layout.Context, th *material.Theme) layout.Dimensions {
	for _, event := range r.appBar.Events(gtx) {
		switch event := event.(type) {
		case component.AppBarNavigationClicked:
			r.navAnim.ToggleVisibility(gtx.Now)
		case component.AppBarContextMenuDismissed:
			log.Printf("Context menu dismissed: %v", event)
		case component.AppBarOverflowActionClicked:
			log.Printf("Overflow action selected: %v", event)
		}
	}

	if r.navDrawer.NavDestinationChanged() {
		r.SwitchTo(r.navDrawer.CurrentNavDestination())
	}

	paint.Fill(gtx.Ops, th.Palette.Bg)

	content := layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				gtx.Constraints.Max.X /= 5
				return r.navDrawer.Layout(gtx, th, &r.navAnim)
			}),
			layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
				return r.pages[r.current].Layout(gtx, th)
			}),
		)
	})
	bar := layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		return r.appBar.Layout(gtx, th, "Menu", "Actions")
	})
	flex := layout.Flex{Axis: layout.Vertical}

	flex.Layout(gtx, bar, content)
	r.modals.Layout(gtx, th)

	return layout.Dimensions{Size: gtx.Constraints.Max}
}
