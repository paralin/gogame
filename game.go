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
}

// Add an entity. This should happen AFTER it's initialized.
func (g *Game) AddEntity(ent *Entity) {
	g.EntityTable[ent.Id] = ent
	if g.Frontend != nil {
		ent.FrontendEntity = g.Frontend.AddEntity(ent)
		ent.InitFrontendEntity()
	}
}

func (g *Game) Destroy() {
	// Delete all entities
	// Unregister all components
	// etc...

	if g.Frontend != nil {
		g.Frontend.Destroy()
	}
}
