package commands

import (
	"math"

	"github.com/RestartFU/practice/custom"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/entity/damage"
)

type SPAWN struct{}

func (SPAWN) Run(src cmd.Source, out *cmd.Output) {
	if p, ok := src.(*custom.Player); ok {
		if cd, ok := p.CombatCD(); ok && !cd.Expired() {
			out.Errorf("You're still in combat for %v seconds", math.Round(cd.UntilExpiration().Seconds()))
			return
		}
		p.Player().Hurt(p.Player().MaxHealth(), damage.SourceCustom{})
	}
}
