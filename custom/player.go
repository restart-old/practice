package custom

import (
	"time"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-plus/ffa"
	"github.com/go-gl/mathgl/mgl64"
)

type Player struct {
	player  *player.Player
	s       *Server
	pearlCD time.Time
	cps     int
	chatCD  time.Time
	combat  time.Time
	ffa     ffa.FFA
}

// NewPlayer returns a new *Player.
func NewPlayer(p *player.Player, s *Server) *Player {
	cPlayer := &Player{player: p, s: s}
	s.AddPlayer(cPlayer)
	return cPlayer
}

// Name ...
func (p *Player) Name() string { return p.player.Name() }

// Position ...
func (p *Player) Position() mgl64.Vec3 { return p.player.Position() }

// SendCommandOutput ...
func (p *Player) SendCommandOutput(output *cmd.Output) {
	p.player.SendCommandOutput(output)
}

// World ...
func (p *Player) World() *world.World { return p.player.World() }

// Server returns the *custom.Server of the player.
func (p *Player) Server() *Server { return p.s }

// Combat returns the last time the user was in combat,
// and if the player is still in combat cooldown.
func (p *Player) Combat() (time.Time, bool) {
	return p.combat, p.combat.After(time.Now())
}

// Combat sets the combat for the player with the duration passed.
func (p *Player) SetCombat(d time.Duration) { p.combat = time.Now().Add(d) }

// PearlCD returns the last time the user threw a pearl,
// and if the player is still on pearl cooldown.
func (p *Player) PearlCD() (time.Time, bool) {
	return p.pearlCD, p.pearlCD.After(time.Now())
}

// SetPearlCD sets the pearl cooldown for the player with the duration passed.
func (p *Player) SetPearlCD(d time.Duration) { p.pearlCD = time.Now().Add(d) }

// SetChatCD sets the chat cooldown for the player with the duration passed.
func (p *Player) SetChatCD(d time.Duration) { p.chatCD = time.Now().Add(d) }

// ChatCD returns the last time the user sent a message,
// and if the player is still on chat cooldown.
func (p *Player) ChatCD() (time.Time, bool) {
	return p.chatCD, p.chatCD.After(time.Now())
}

// AddCPS adds the int passed to the player cps,
// and removes that same amount a second after.
func (p *Player) AddCPS(n int) {
	p.cps += n
	time.AfterFunc(1*time.Second, func() {
		p.cps -= n
	})
}

// CPS returns the current amount of cps of the player.
func (p *Player) CPS() int { return p.cps }

// Player returns the *player.Player of the *Player.
func (p *Player) Player() *player.Player { return p.player }

// Operator returns a bool of if the player is an operator or not.
func (p *Player) Operator() bool { return p.Server().operators.Listed(p.Name()) }

// Whitelisted returns a bool of if the player is whitelisted or not.
func (p *Player) Whitelisted() bool { return p.Server().whitelist.Listed(p.Name()) }

// Banned returns a bool of if the player is banned or not.
func (p *Player) Banned() bool { return p.Server().ban.Listed(p.Name()) }

// SetFFA ...
func (p *Player) SetFFA(ffa ffa.FFA) { p.ffa = ffa }

// FFA ...
func (p *Player) FFA() (ffa.FFA, bool) { return p.ffa, p.ffa != nil }
