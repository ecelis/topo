/*
Copyright © 2025 Ernesto Celis <ernesto@patito.ninja>
*/
package service

import "fmt"

// Import the route package

func Run() error {
	fmt.Println("Start")
	// routes := []route.Route{
	// 	{Destination: "10.0.0.0", Netmask: "255.255.255.0", Gateway: "10.70.32.193"},
	// 	{Destination: "0.0.0.0", Netmask: "0.0.0.0", Gateway: "10.182.112.1"},
	// }

	// for _, r := range routes {
	// 	if err := platform.AddRoute(r); err != nil {
	// 		return fmt.Errorf("error adding route: %w", err)
	// 	}
	// }

	// // Periodic check
	// ticker := time.NewTicker(5 * time.Minute)
	// defer ticker.Stop()

	// for range ticker.C {
	// 	// Check routes, udpdate if needed
	// }

	// // Notify systemd
	// _, err := daemon.SdNotify(false, "READY=1")
	// if err != nil {
	// 	log.Println("Systemd notify error:", err)
	// }

	return nil
}
