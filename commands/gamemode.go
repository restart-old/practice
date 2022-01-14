package commands

import (
	"Practice/custom"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/world"
)

type Gamemode struct {
	Gamemode gm
	server   *custom.Server
}

func GamemodeRunnable(server *custom.Server) cmd.Runnable { return Gamemode{server: server} }

func (g Gamemode) Run(src cmd.Source, o *cmd.Output) {
	if p, ok := src.(*custom.Player); ok {
		if gm, ok := gms[g.Gamemode]; ok {
			p.Player().SetGameMode(gm)
		}
	}
}

var gms = map[gm]world.GameMode{
	"creative":  world.GameModeCreative,
	"survival":  world.GameModeSurvival,
	"spectator": world.GameModeSpectator,
	"adventure": world.GameModeAdventure,
}

type gm string

func (gm) Type() string { return "gamemode" }

func (gm) Options(cmd.Source) []string {
	return []string{"creative", "survival", "spectator", "adventure"}
}

func (g Gamemode) Allow(src cmd.Source) bool {
	return g.server.Operators().Listed(src.Name())
}
