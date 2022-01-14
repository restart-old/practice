package handler

import (
	"strings"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/event"
)

func (h PlayerHandler) HandleCommandExecution(ctx *event.Context, command cmd.Command, args []string) {
	ctx.Cancel()
	command.Execute(strings.Join(args, " "), h.p)
}
