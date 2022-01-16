package main

import (
	"fmt"
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
		fmt.Printf("%s[%v] couldn't join the server: player is banned\n", d.DisplayName, addr.String())
		return "You are banned", false
	}

	if !server.Whitelist().Listed(d.DisplayName) && server.Whitelist().Enabled {
		fmt.Printf("%s[%v] couldn't join the server: player is not whitelisted\n", d.DisplayName, addr.String())
		return "server is whitelisted", false
	}
	fmt.Printf("%s[%v] has joined the server\n", d.DisplayName, addr.String())
	return "", true
}
