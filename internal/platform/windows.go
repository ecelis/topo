//go:build windows
// +build windows

package platform

import (
	"fmt"
	"net"

	"github.com/ecelis/topo/internal/route"
)

// Windows Implementation (using IP Helper API)
// This example is very basic and requires more error handling and proper interface selection.
// Consider using a third party library.
func AddRoute(r route.Route) error {
	dst := net.ParseIP(r.Destination)
	if dst == nil {
		return fmt.Errorf("invalid destination IP: %s", r.Destination)
	}

	gateway := net.ParseIP(r.Gateway)
	if gateway == nil {
		return fmt.Errorf("invalid gateway IP: %s", r.Gateway)
	}

	mask := net.ParseIPMask(r.Netmask)

	// Convert IP addresses and mask to uint32
	dst32 := ipToUint32(dst)
	gw32 := ipToUint32(gateway)
	mask32 := ipMaskToUint32(mask)

	// Call the AddIPRoute function (requires appropriate privileges)
	err := addIPRoute(dst32, mask32, gw32, 1) // Metric is 1
	if err != nil {
		return fmt.Errorf("error adding route (Windows): %w", err)
	}

	return nil
}

func ipToUint32(ip net.IP) uint32 {
	return uint32(ip[0]) | uint32(ip[1])<<8 | uint32(ip[2])<<16 | uint32(ip[3])<<24
}

func ipMaskToUint32(ip net.IPMask) uint32 {
	return uint32(ip[0]) | uint32(ip[1])<<8 | uint32(ip[2])<<16 | uint32(ip[3])<<24
}

// This is a stub, you'll have to implement the actual call to the windows API
// See https://learn.microsoft.com/en-us/windows/win32/api/iphlpapi/nf-iphlpapi-addiproute
func addIPRoute(destination uint32, mask uint32, gateway uint32, metric uint32) error {
	fmt.Println("Adding route", destination, mask, gateway, metric)
	return nil
}
