package ffas

import (
	"fmt"

	"github.com/RestartFU/practice/custom"
	"github.com/RestartFU/practice/kits"

	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-plus/kit"
	"github.com/go-gl/mathgl/mgl64"
)

func NoDebuffFFA(s *custom.Server) NoDebuff { return NoDebuff{server: s} }

type NoDebuff struct {
	server *custom.Server
}

func (NoDebuff) Name() string { return "NoDebuff" }
func (ndbf NoDebuff) World() *world.World {
	worldName := ndbf.Name()
	w, ok := ndbf.server.WorldManager().World(worldName)
	if !ok {
		fmt.Printf("ffa: world '%s' not loaded\n", worldName)
	}
	return w.World
}
func (NoDebuff) Position() mgl64.Vec3 { return mgl64.Vec3{220, 67, 210} }
func (ndbf NoDebuff) Kit() kit.Kit {
	kitName := ndbf.Name()
	k, ok := kits.ByName(kitName)
	if !ok {
		fmt.Printf("ffa: kit '%s' not registered\n", kitName)
	}
	return k
}
