package gogame

/* An instance of a game. */
type GameRules struct {
	// All registered components
	ComponentTable ComponentTable

	// All known entities
	EntityTable EntityTable
}
