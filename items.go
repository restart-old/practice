package main

import (
	"fmt"
	"strings"

	"github.com/RestartFU/practice/ffas"

	"github.com/RestartFU/practice/custom"

	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/tool"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/form"
	"github.com/df-mc/dragonfly/server/world"
)

var (
	ErrUserNotPlayer        = "FFASword: couldn't use FFASword because user isn't '*player.Player'"
	ErrFFANotRegistered     = "FFASword: ffa with name '%s' isn't registered\n"
	ErrCustomPlayerNotFound = "FFASword: custom player with name '%s' not found\n"
)

// FFASword ...
type FFASword struct {
	server *custom.Server
}

// playersInWorldByName ...
func playersInWorldByName(name string, s *custom.Server) int {
	var n int
	if w, ok := s.WorldManager().World(name); ok {
		for _, e := range w.Entities() {
			if _, ok := e.(*player.Player); ok {
				n++
			}
		}
	}
	return n
}

// newButton ...
func newButton(name, texture string, s *custom.Server) form.Button {
	players := playersInWorldByName(name, s)
	return form.NewButton(fmt.Sprintf(name+"\n%v Player(s)", players), texture)
}

// Name returns the name that the item must have for Use() to be ran.
func (FFASword) Name() string { return "§eFFA - Unranked" }

// Item returns the item that must be used.
func (FFASword) Item() world.Item { return item.Sword{Tier: tool.TierDiamond} }

// Use sends the from for the FFASword item.
func (ffa FFASword) Use(s item.Stack, p *player.Player) {
	nodebuff := newButton("NoDebuff", "textures/items/potion_bottle_splash_heal.png", ffa.server)
	gapple := newButton("Gapple", "textures/items/apple_golden.png", ffa.server)
	fist := newButton("Fist", "textures/items/beef_cooked.png", ffa.server)

	f := form.NewMenu(FFASword{server: ffa.server}, "§eFFA - Unranked").WithButtons(
		nodebuff,
		gapple,
		fist,
	)
	p.SendForm(f)
}

// Submit teleports the player to the chosen FFA.
func (item FFASword) Submit(submitter form.Submitter, pressed form.Button) {
	// Making sure the submitter is a *player.Player.
	p, ok := submitter.(*player.Player)
	if !ok {
		fmt.Println(ErrUserNotPlayer)
		return
	}

	// Making sure the chosen ffa is registered.
	ffaName := strings.Split(pressed.Text, "\n")[0]
	ffa, ok := ffas.ByName(ffaName)
	if !ok {
		fmt.Printf(ErrFFANotRegistered, ffaName)
		return
	}

	// Making sure the custom player exists.
	pl, ok := item.server.PlayerByName(p.Name())
	if !ok {
		fmt.Printf(ErrCustomPlayerNotFound, p.Name())
		return
	}

	// Set the player FFA.
	pl.SetFFA(ffa)

	// Teleport the player to the FFA.
	ffas.TeleportToFFA(p, ffa)
}
