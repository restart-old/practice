package commands

import "github.com/df-mc/dragonfly/server/cmd"

func Register(cmds ...cmd.Command) {
	for _, c := range cmds {
		cmd.Register(c)
	}
}
