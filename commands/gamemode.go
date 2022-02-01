package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/world"
)

// GameMode returns the gamemode command
func GameMode() cmd.Command {
	return cmd.New("gamemode", "update your gamemode", []string{"gm"},
		GameModeWord{},
		GameModeInt{},
		GameModePlayerInt{},
		GameModePlayerWord{},
	)
}

<<<<<<< Updated upstream
func GamemodeRunnable(server *custom.Server) cmd.Runnable { return Gamemode{server: server} }

func (g Gamemode) Run(src cmd.Source, o *cmd.Output) {
	if p, ok := src.(*custom.Player); ok {
		if gm, ok := gms[g.Gamemode]; ok {
			p.Player().SetGameMode(gm)
		}
	}
=======
// modeInt is a map of the n enum pointing to a gamemode.
var modeInt = map[n]world.GameMode{
	"0": world.GameModeSurvival,
	"1": world.GameModeCreative,
	"2": world.GameModeAdventure,
	"3": world.GameModeSpectator,
>>>>>>> Stashed changes
}

// modeWord is a map of the gamemode enum pointing to a gamemode.
var modeWord = map[gameMode]world.GameMode{
	"survival":  world.GameModeSurvival,
	"creative":  world.GameModeCreative,
	"adventure": world.GameModeAdventure,
	"spectator": world.GameModeSpectator,
}

// n is the enum argument used to change your gamemode using an integer.
type n string

// Type ...
func (n) Type() string { return "GameModeInt" }

// Options ...
func (n) Options(_ cmd.Source) []string {
	return []string{
		"0",
		"1",
		"2",
		"3",
	}
}

// gameMode is the enum argument used to change your gamemode with the name of the gamemode.
type gameMode string

// Type ...
func (gameMode) Type() string { return "GameModeWord" }

// Options ...
func (gameMode) Options(_ cmd.Source) []string {
	return []string{
		"survival",
		"creative",
		"adventure",
		"spectator",
	}
}

// GameModeInt is the command used to change your gamemode with an integer.
type GameModeInt struct {
	GameMode n
}

// Run ...
func (g GameModeInt) Run(src cmd.Source, _ *cmd.Output) {
	if gm, ok := src.(interface{ SetGameMode(mode world.GameMode) }); ok {
		gm.SetGameMode(modeInt[g.GameMode])
	}
}

// GameModePlayerInt is the command used to change the gamemode of an online player with an integer.
type GameModePlayerInt struct {
	Target   []cmd.Target
	GameMode n
}

// Run ...
func (g GameModePlayerInt) Run(_ cmd.Source, _ *cmd.Output) {
	if gm, ok := g.Target[0].(interface{ SetGameMode(mode world.GameMode) }); ok {
		gm.SetGameMode(modeInt[g.GameMode])
	}
}

// GameModePlayerWord is the command used to change the gamemode of an online player with the name of the gamemode.
type GameModePlayerWord struct {
	Target   []cmd.Target
	GameMode gameMode
}

// Run ...
func (g GameModePlayerWord) Run(_ cmd.Source, _ *cmd.Output) {
	if gm, ok := g.Target[0].(interface{ SetGameMode(mode world.GameMode) }); ok {
		gm.SetGameMode(modeWord[g.GameMode])
	}
}

// GameModeWord is the command used to change your gamemode with the name of the gamemode.
type GameModeWord struct {
	GameMode gameMode
}

// Run ...
func (g GameModeWord) Run(src cmd.Source, out *cmd.Output) {
	if gm, ok := src.(interface{ SetGameMode(mode world.GameMode) }); ok {
		gm.SetGameMode(modeWord[g.GameMode])
	}
}
