/*
Copyright © 2025 Ernesto Celis <ernesto@patito.ninja>
*/
package linux

import "log/syslog"

type SysLogger struct {
	syslog *syslog.Writer
}

func (s *SysLogger) Info(msg string)  { s.syslog.Info(msg) }
func (s *SysLogger) Error(msg string) { s.syslog.Err(msg) }
