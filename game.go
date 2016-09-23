package gogame

/* An instance of a game. */
type Game struct {
	// All registered components
	ComponentTable ComponentTable

	// All known entities
	EntityTable EntityTable

	// Game rules instance
	GameRules GameRules

	// Optional frontend instance
	Frontend Frontend

	// Optional frontend game rules
	FrontendGameRules FrontendGameRules

	// Game current operating mode
	OperatingMode GameOperatingMode

	// Game settings
	GameSettings GameSettings

	// Game state
	GameState *GameState
}

// Add an entity. This should happen AFTER it's initialized.
func (g *Game) AddEntity(ent *Entity) {
	g.EntityTable[ent.Id] = ent
	if g.Frontend != nil {
		ent.FrontendEntity = g.Frontend.AddEntity(ent)
		ent.InitFrontendEntity()
	}
}

// Propogate the operating mode change
func (g *Game) setOperatingMode(opMode GameOperatingMode) {
	g.OperatingMode = opMode
	if g.FrontendGameRules != nil {
		g.FrontendGameRules.SetGameOperatingMode(opMode)
	}
	if g.GameRules != nil {
		g.GameRules.SetGameOperatingMode(opMode)
	}
}

func (g *Game) Destroy() {
	// Delete all entities
	// Unregister all components
	// etc...

	g.GameRules.Destroy()

	if g.FrontendGameRules != nil {
		g.FrontendGameRules.Destroy()
	}

	if g.Frontend != nil {
		g.Frontend.Destroy()
	}
}
