//go:build windows
// +build windows

package platform

import (
	"fmt"
	"os/exec"

	"github.com/ecelis/topo/internal/route"
)

func AddRoute(r route.Route) error {
	cmd := exec.Command("route", "add", r.Destination, "mask", r.Netmask, r.Gateway)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error adding route (Windows): %v\nOutput: %s", err, output)
	}
	return nil
}

// ... other Windows-specific functions
