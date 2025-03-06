/*
Copyright © 2025 Ernesto Celis <ernesto@patito.ninja>
*/
package platform

const (
	ErrSyslogOpen        = "error opening syslog: %w"
	ErrInsufficientPrivs = "insufficient privileges, please run with sudo"
	ErrAddRoute          = "error adding route (%s): %v\nOutput: %s"

	LogAddDefaultRoute = "Adding default route to %s"
	LogAddRoute        = "Adding route to %s via %s"
	LogRouteAddFailed  = "Route add failed: %w"
	LogUnauthorized    = "Unauthorized route change attempt by %s"
)
