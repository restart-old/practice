package commands

import (
	"Practice/custom"
	"math"
	"time"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/entity/damage"
)

type SPAWN struct{}

func (SPAWN) Run(src cmd.Source, out *cmd.Output) {
	if p, ok := src.(*custom.Player); ok {
		if t, ok := p.Combat(); ok {
			out.Errorf("You're still in combat for %v seconds", math.Round(time.Until(t).Seconds()))
			return
		}
		p.Player().Hurt(p.Player().MaxHealth(), damage.SourceCustom{})
	}
}
