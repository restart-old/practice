package handler

func (h PlayerHandler) HandleQuit() {
	h.p.Server().RemovePlayer(h.p)
}
