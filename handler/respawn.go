package handler

import (
	"time"

	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/tool"
	"github.com/go-gl/mathgl/mgl64"
)

func (h PlayerHandler) HandleRespawn(pos *mgl64.Vec3) {
	*pos = mgl64.Vec3{291, 77, 176}

	time.AfterFunc(1, func() {
		w := h.p.Server().World()
		w.AddEntity(h.p.Player())

		h.p.Player().Inventory().SetItem(0, item.NewStack(item.Sword{Tier: tool.TierDiamond}, 1).WithCustomName("§eFFA - Unranked"))
	})
}
