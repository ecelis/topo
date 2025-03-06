/*
Copyright © 2025 Ernesto Celis <ernesto@patito.ninja>
*/
package linux

import (
	"os/exec"
	"strings"

	"github.com/ecelis/topo/internal/route"
)

func isDefaultRoute(elements []string, route route.Route) bool {
	if elements[0] == route.Destination &&
		elements[2] == route.Gateway &&
		elements[4] == route.Device {
		return true
	}
	return false
}

func isStaticRoute(elements []string, route route.Route) bool {
	if elements[0] == route.Destination &&
		elements[2] == route.Gateway {
		return true
	}
	return false
}

func RouteExists(r route.Route) (bool, error) {
	output, err := exec.Command("ip", "route", "show", r.Destination).Output()

	if err != nil {
		return false, err
	}

	elements := strings.Split(string(output), " ")

	if isDefaultRoute(elements, r) {
		return true, nil
	}

	if isStaticRoute(elements, r) {
		return true, nil
	}

	return false, nil
}
