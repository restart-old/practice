package main

type gamemode struct{}

func (gamemode) AllowsEditing() bool      { return false }
func (gamemode) AllowsTakingDamage() bool { return true }
func (gamemode) CreativeInventory() bool  { return false }
func (gamemode) AllowsFlying() bool       { return false }
func (gamemode) HasCollision() bool       { return true }
func (gamemode) AllowsInteraction() bool  { return true }
func (gamemode) Visible() bool            { return true }
