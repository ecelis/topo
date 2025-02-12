package main

import (
	"log"

	"github.com/ecelis/topo/cmd/topo/service"
)

func main() {
	routes, err := service.GetConfiguredRoutes() // Implement this to get routes
	if err != nil {
		log.Fatal(err)
	}

	go service.Run(routes) // Run service logic in a goroutine
	runGUI(routes)         // Start the GUI (blocks)
}
