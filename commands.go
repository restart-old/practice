package main

import (
	"github.com/RestartFU/practice/custom"

	"github.com/RestartFU/practice/commands"

	"github.com/RestartFU/list"
	"github.com/df-mc/dragonfly/server/cmd"
)

// RegisterAllCommands registers all commands
func RegisterAllCommands(s *custom.Server) {
	commands.Register(
		cmd.New("gamemode", "", []string{"gm"}, commands.GamemodeRunnable(s)),
		cmd.New("id", "get the id if the item you're currently holding", nil, commands.IDRunnable(s)),
		cmd.New("op", "add someone to the operators list", nil, commands.OPRunnable(s)),
		cmd.New("deop", "remove someone from the operators list", nil, commands.DEOPRunnable(s)),
		cmd.New("spawn", "teleport to spawn", []string{"hub"}, commands.SPAWN{}),
		cmd.New("ban", "ban someone", nil, commands.BanRunnable(s)),
		cmd.New("unban", "unban someone", nil, commands.UNBANRunnable(s)),
		cmd.New("whitelist", "enable or disable the whitelist", []string{"wl"}, list.NewRunnable(s.Whitelist(), func(src cmd.Source) bool { return s.Operators().Listed(src.Name()) })),
	)
}
