package platform

import (
	"fmt"
	"os/exec"

	"github.com/ecelis/topo/internal/route"
)

func AddRoute(r route.Route) error {
	cmd := exec.Command("route", "add", r.Destination, "netmask", r.Netmask, "gw", r.Gateway)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error adding route (Linux): %v\nOutput: %s", err, output)
	}
	return nil
}

// ... other Linux-specific functions (e.g., delete route)
