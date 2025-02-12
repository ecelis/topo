package service

import (

	// Import the platform package
	"github.com/ecelis/topo/internal/route" // Import the route package
)

func Run(routes []route.Route) error {
	// ... (Your service logic) ...

	// Now routes are available here
	for _, r := range routes {
		// ... Use the routes
	}
	return nil
}

func GetConfiguredRoutes() ([]route.Route, error) {
	//Logic to obtain the current routes, this could be reading from the routing table, a configuration file, etc.
	return []route.Route{}, nil
}
