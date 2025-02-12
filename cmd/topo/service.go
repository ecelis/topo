package service

import (
    "fmt"
    "log"
    "time"

    "github.com/ecelis/topo/internal/platform" // Import the platform package
    "github.com/ecelis/topo/internal/route" // Import the route package
    "github.com/coreos/go-systemd/daemon"
)

func Run() error {

        //Get the routes to set, from a file, an environment variable, etc.
        routes := []route.Route{
                {Destination: "10.0.0.0", Netmask: "255.0.0.0", Gateway: "10.70.32.193"},
                {Destination: "0.0.0.0", Netmask: "0.0.0.0", Gateway: "10.182.112.1"},
        }

        for _, r := range routes {
                if err := platform.AddRoute(r); err != nil {
                        return fmt.Errorf("error adding route: %w", err)
                }
                fmt.Printf("Route added: %v\n", r)
        }

        // ... Your service logic (e.g., periodic route checks, etc.) ...

        // Example: Periodic check (adapt as needed)
        ticker := time.NewTicker(5 * time.Minute) // Check every 5 minutes
        defer ticker.Stop()

        for range ticker.C {
                // ... (Check routes, update if needed, etc.) ...
        }

        // Notify systemd (Linux)
        _, err := daemon.SdNotify(false, "READY=1")
        if err != nil {
                log.Println("Systemd notify error:", err) // Not fatal on Windows
        }

    return nil
}