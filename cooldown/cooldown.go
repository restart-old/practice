package cooldown

type CoolDown struct {
	cooldown
}

func (cd CoolDown) String() string { return string(cd.cooldown) }

type cooldown string

func Combat() CoolDown     { return CoolDown{"cooldown_combat"} }
func EnderPearl() CoolDown { return CoolDown{"cooldown_enderPearl"} }
func Chat() CoolDown       { return CoolDown{"cooldown_chat"} }
