package handler

import (
	"fmt"
	"github.com/RestartFU/practice/cooldown"
	"github.com/RestartFU/practice/custom"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/chat"
	"time"
)

// playerHandler is the handler used to handle the actions of a player
type playerHandler struct {
	player.NopHandler
	Player *custom.Player
}

// PlayerHandler returns a new *playerHandler
func PlayerHandler(player *custom.Player) *playerHandler {
	return &playerHandler{Player: player}
}

// HandleChat handles the messages sent by a player.
// It sets the player on a 3 seconds chat cooldown each time they send a message.
func (handler *playerHandler) HandleChat(ctx *event.Context, message *string) {
	ctx.Cancel()

	player := handler.Player
	if cd := player.CoolDown(cooldown.Chat()); cd.Expired() {
		_, err := chat.Global.WriteString(fmt.Sprintf("§7%s§f: %s", player.DisplayName(), *message))
		if err != nil {
			player.Messagef("§can error occurred while trying to send a message: %s", err)
		}
		cd.SetCooldown(3 * time.Second)
	} else {
		player.Messagef("§cYou're on §lChat§r§c cooldown for §l%.2f Seconds§r", cd.UntilExpiration().Seconds())
	}
}
