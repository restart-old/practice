package main

import (
	"github.com/RestartFU/practice/custom"
	"github.com/sandertv/gophertunnel/minecraft/protocol/login"
	"net"
)

type Allower struct {
	s *custom.Server
}

func (a Allower) Allow(addr net.Addr, d login.IdentityData, c login.ClientData) (string, bool) {
	if wl := a.s.Whitelist(); !wl.Listed(d.DisplayName) && wl.Enabled {
		return "the server is whitelisted", false
	}
	if conn, ok := a.s.AllowedData()[c.DeviceModel]; !ok || conn.IdentityData().DisplayName != d.DisplayName {
		return "bruh", false
	} else {
		a.s.SetConn(conn, addr)
	}
	return "", true
}
