package gogame

type GameRules interface {
	// Called just after Game is finished being built
	Init(game *Game)
	SetGameOperatingMode(opMode GameOperatingMode)
	Update()
	Destroy()
}
