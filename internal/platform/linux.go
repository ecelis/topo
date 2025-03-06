//go:build linux
// +build linux

/*
Copyright © 2025 Ernesto Celis <ernesto@patito.ninja>
*/
package platform

import (
	"fmt"
	"log/syslog"
	"os/exec"
	"os/user"

	"github.com/ecelis/topo/internal/route"
)

func AddRoute(r route.Route) error {
	logger, err := syslog.New(syslog.LOG_INFO|syslog.LOG_DAEMON, "topo")
	if err != nil {
		return fmt.Errorf(ErrSyslogOpen, err)
	}
	defer logger.Close()

	var cmd *exec.Cmd
	if r.Destination == "0.0.0.0" {
		cmd = exec.Command("route", "add", "default", "gw", r.Gateway)
		logger.Info(fmt.Sprintf(LogAddDefaultRoute, r.Gateway))
	} else {
		cmd = exec.Command("route", "add", "-net", r.Destination, "netmask", r.Netmask, "gw", r.Gateway)
		logger.Info(fmt.Sprintf(LogAddRoute, r.Destination, r.Gateway))
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		logger.Err(fmt.Sprintf(LogRouteAddFailed, err))
		if string(output) == "SIOCADDRT: Operation not permitted\n" {
			currentUser, _ := user.Current()
			logger.Warning(fmt.Sprintf(LogUnauthorized, currentUser.Username))
			return fmt.Errorf(ErrInsufficientPrivs)
		}
		return fmt.Errorf(ErrAddRoute, "Linux", err, output)
	}
	return nil
}
