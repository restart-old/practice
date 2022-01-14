package ffas

import (
	"sync"

	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-plus/ffa"
)

func init() {
	ffas = make(map[string]ffa.FFA)
}

var ffas map[string]ffa.FFA
var ffasMu sync.RWMutex

func Register(f ffa.FFA) {
	ffasMu.Lock()
	defer ffasMu.Unlock()
	ffas[f.Name()] = f
}
func ByName(name string) (ffa.FFA, bool) {
	ffasMu.RLock()
	defer ffasMu.RUnlock()
	f, ok := ffas[name]
	return f, ok
}

func TeleportToFFA(p *player.Player, f ffa.FFA) {
	ffa.TeleportToFFA(p, f)
}
