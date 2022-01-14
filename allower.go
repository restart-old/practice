package main

import (
	"net"

	"github.com/RestartFU/practice/custom"

	"github.com/sandertv/gophertunnel/minecraft/protocol/login"
)

type Allower struct {
	s *custom.Server
}

func (a Allower) Allow(addr net.Addr, d login.IdentityData) (string, bool) {
	server := a.s
	if server.Ban().Listed(d.DisplayName) {
		return "You are banned", false
	}

	if !server.Whitelist().Listed(d.DisplayName) && server.Whitelist().Enabled {
		return "server is whitelisted", false
	}
	return "", true
}
