package platform

import (
	"fmt"

	"github.com/ecelis/topo/internal/platform/linux"
	"github.com/ecelis/topo/internal/route"
)

type RouteChecker struct {
	// Logger Logger // interface for syslog/eventlog
	Routes []route.Route
}

func (rc *RouteChecker) CheckAndUpdateRoutes() error {
	for _, expected := range rc.Routes {
		exists, err := linux.RouteExists(expected)
		if err != nil {
			fmt.Println(fmt.Sprintf("Failed checking route %s: %v", expected.Destination, err))
			// rc.logger.Error(fmt.Sprintf("Failed checking route %s: %v", expected.Destination, err))
			continue
		}

		if !exists {
			fmt.Println(fmt.Sprintf("Route %s missing, reinstalling", expected.Destination))
			// rc.logger.Info(fmt.Sprintf("Route %s missing, reinstalling", expected.Destination))
			// if err := AddRoute(expected); err != nil {
			// 	rc.logger.Error(fmt.Sprintf("Failed reinstalling route %s: %v", expected.Destination, err))
			// }
		}
	}
	return nil
}
