/*
Copyright © 2025 Ernesto Celis <ernesto@patito.ninja>
*/
package linux

import (
	"fmt"
	"net"
	"os/exec"
	"strings"

	"github.com/ecelis/topo/internal/route"
)

func isDefaultRoute(elements []string, route route.Route) bool {
	if elements[0] == route.Destination &&
		elements[2] == route.Gateway &&
		elements[4] == route.Device {
		return true
	}
	return false
}

func isStaticRoute(elements []string, destination string, gw string) bool {
	if elements[0] == destination &&
		elements[8] == gw {
		return true
	}
	return false
}

func netmaskToBits(netmask string) (bits int, err error) {
	ip := net.ParseIP(netmask)
	if ip == nil {
		return 0, fmt.Errorf("invalid netmask: %s", netmask)
	}

	mask := ip.To4()
	if mask == nil {
		return 0, fmt.Errorf("invalid netmask: %s", netmask)
	}

	bits = 0
	for _, b := range mask {
		for i := 7; i >= 0; i-- {
			if b>>uint(i)&1 == 1 {
				bits++
			} else {
				return bits, nil
			}
		}
	}
	return bits, nil
}

func validateDestination(r route.Route) (destination string, err error) {
	ip := net.ParseIP(r.Destination)
	if r.Destination != "default" && ip == nil {
		return "", fmt.Errorf("invalid destination: %s", r.Destination)
	}
	if r.Destination == "default" {
		return r.Destination, nil
	}
	bits, err := netmaskToBits(r.Netmask)
	if err != nil {
		return "", err
	}

	return ip.String() + "/" + fmt.Sprint(bits), err
}

func RouteExists(r route.Route) (bool, error) {
	destination, err := validateDestination(r)
	if err != nil {
		return false, err
	}
	output, err := exec.Command("ip", "route", "show", destination).Output()
	if err != nil {
		return false, err
	}

	elements := strings.Split(string(output), " ")

	if isDefaultRoute(elements, r) {
		return true, nil
	}

	if isStaticRoute(elements, destination, r.Gateway) {
		return true, nil
	}

	return false, nil
}
