package commands

import (
	"github.com/RestartFU/list"
	"github.com/df-mc/dragonfly/server/cmd"
)

// WhiteList returns the whitelist command.
func WhiteList(wl *list.List) cmd.Command {
	return cmd.New("whitelist", "", []string{"wl"},
		WhiteListAddOffline{whitelist: wl},
		WhiteListAddOnline{whitelist: wl},
		WhiteListStatus{whitelist: wl},
		WhiteListRemoveOnline{whitelist: wl},
		WhiteListRemoveOffline{whitelist: wl},
	)
}

// status is the enum argument used to either disable or enable the whitelist.
type status string

// Type ...
func (status) Type() string { return "status" }

// Options ...
func (status) Options(src cmd.Source) []string {
	return []string{"enable", "disable"}
}

// add is the SubCommand used to add a player to the whitelist.
type add string

// SubName ...
func (add) SubName() string { return "add" }

// remove is the SubCommand used to remove a player to the whitelist.
type remove string

// SubName ...
func (remove) SubName() string { return "remove" }

// WhiteListRemoveOnline is the command used to remove an online player from the whitelist.
type WhiteListRemoveOnline struct {
	Sub    remove
	Player []cmd.Target

	whitelist *list.List
}

// Run ...
func (c WhiteListRemoveOnline) Run(src cmd.Source, out *cmd.Output) {
	err := c.whitelist.Remove(c.Player[0].Name())
	if err != nil {
		out.Errorf("whitelist error: %s", err)
	}
}

// WhiteListRemoveOffline is the command used to remove an offline player from the whitelist.
type WhiteListRemoveOffline struct {
	Sub    remove
	Player string

	whitelist *list.List
}

// Run ...
func (c WhiteListRemoveOffline) Run(src cmd.Source, out *cmd.Output) {
	err := c.whitelist.Remove(c.Player)
	if err != nil {
		out.Errorf("whitelist error: %s", err)
	}
}

// WhiteListStatus is the command used to either disable or enable the whitelist.
type WhiteListStatus struct {
	Status status

	whitelist *list.List
}

// Run ...
func (c WhiteListStatus) Run(src cmd.Source, out *cmd.Output) {
	switch c.Status {
	case "enable":
		if c.whitelist.Enabled {
			out.Error("server is already whitelisted")
		} else {
			c.whitelist.Enabled = true
		}
	case "disable":
		if !c.whitelist.Enabled {
			out.Error("server is not whitelisted")
		} else {
			c.whitelist.Enabled = false
		}
	}
}

// WhiteListAddOnline is the command used to add an online player to the whitelist.
type WhiteListAddOnline struct {
	Sub    add
	Player []cmd.Target

	whitelist *list.List
}

// Run ...
func (c WhiteListAddOnline) Run(src cmd.Source, out *cmd.Output) {
	err := c.whitelist.Add(c.Player[0].Name())
	if err != nil {
		out.Errorf("whitelist error: %s", err)
	}
}

// WhiteListAddOffline is the command used to add an offline player to the whitelist.
type WhiteListAddOffline struct {
	Sub    add
	Player string

	whitelist *list.List
}

// Run ...
func (c WhiteListAddOffline) Run(src cmd.Source, out *cmd.Output) {
	err := c.whitelist.Add(c.Player)
	if err != nil {
		out.Errorf("whitelist error: %s", err)
	}
}
