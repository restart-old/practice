package handler

func (h PlayerHandler) HandleQuit() {
	/*if _, combat := h.p.Combat(); combat {
		if src, ok := h.p.LastHurt(); ok {
			h.p.Kill(damage.SourceEntityAttack{Attacker: src})
		}
	}*/
	h.p.Server().RemovePlayer(h.p)
}
