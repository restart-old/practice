package custom

import (
	"fmt"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-plus/cooldown"
)

// Player embeds *player.Player, so we can have more information into the struct.
type Player struct {
	// This is the embedded *player.Player.
	*player.Player

	// The display name is the string that will be shown as the name of the player
	// in chat and in the name tag.
	displayName string

	// The coolDown manager manages the coolDowns for the Player
	coolDownManager *cooldown.Manager
}

// NewPlayer returns a new *Player.
func NewPlayer(p *player.Player) *Player {
	return &Player{
		Player:          p,
		displayName:     p.Name(),
		coolDownManager: cooldown.NewManager(),
	}
}

// DisplayName returns the display name of the player.
func (player *Player) DisplayName() (displayName string) { return player.displayName }

// SetDisplayName sets a new display name for the player.
// This display name will be shown in chat and in the name tag.
func (player *Player) SetDisplayName(displayName string) { player.displayName = displayName }

// CoolDown returns the *cooldown.Cooldown of the argument passed.
func (player *Player) CoolDown(cd fmt.Stringer) *cooldown.Cooldown {
	return player.coolDownManager.Cooldown(cd.String())
}
