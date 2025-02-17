package platform

import (
	"fmt"

	"github.com/ecelis/topo/internal/route"
)

type RouteChecker struct {
	logger Logger // interface for syslog/eventlog
	routes []route.Route
}

func (rc *RouteChecker) checkAndUpdateRoutes() error {
	for _, expected := range rc.routes {
		exists, err := RouteExists(expected)
		if err != nil {
			rc.logger.Error(fmt.Sprintf("Failed checking route %s: %v", expected.Destination, err))
			continue
		}

		if !exists {
			rc.logger.Info(fmt.Sprintf("Route %s missing, reinstalling", expected.Destination))
			if err := AddRoute(expected); err != nil {
				rc.logger.Error(fmt.Sprintf("Failed reinstalling route %s: %v", expected.Destination, err))
			}
		}
	}
	return nil
}
