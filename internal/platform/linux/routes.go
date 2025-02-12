package linux

import (
	"os/exec"
	"strings"

	"github.com/ecelis/topo/internal/route"
)

func RouteExists(r route.Route) (bool, error) {
	output, err := exec.Command("ip", "route", "show", r.Destination).Output()
	if err != nil {
		return false, err
	}
	return strings.Contains(string(output), r.Gateway), nil
}
