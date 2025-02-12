package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/ecelis/topo/internal/route"
)

type RouteList struct {
	list   *widget.List
	routes []route.Route
}

func NewRouteList(routes []route.Route) *RouteList {
	rl := &RouteList{
		routes: routes,
	}

	rl.list = widget.NewList(
		func() int {
			return len(rl.routes)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Route Info") // Template item
		},
		func(id widget.ListItemID, o fyne.CanvasObject) {
			route := rl.routes[id]
			o.(*widget.Label).SetText(fmt.Sprintf("%s/%s via %s", route.Destination, route.Netmask, route.Gateway))
		},
	)

	return rl
}

func (rl *RouteList) UpdateRoutes(routes []route.Route) {
	rl.routes = routes
	rl.list.Refresh()
}

func (rl *RouteList) GetWidget() fyne.CanvasObject {
	return rl.list
}
