package gogame

import "errors"

/* An instance of a game. */
type Game struct {
	// All registered components
	ComponentTable ComponentTable

	// All known entities
	EntityTable EntityTable

	// Game rules instance
	GameRules GameRules
}

func (g *Game) Destroy() {
	// Delete all entities
	// Unregister all components
	// etc...
}

func BuildGame(componentTable ComponentTable, gameRules GameRules) (*Game, error) {
	if componentTable == nil {
		return nil, errors.New("Component table must not be nil.")
	}

	if gameRules == nil {
		return nil, errors.New("Game rules instance must not be nil.")
	}

	game := &Game{
		EntityTable:    NewEntityTable(),
		GameRules:      gameRules,
		ComponentTable: componentTable,
	}

	// Test component table
	if err := game.ComponentTable.Validate(); err != nil {
		return nil, err
	}

	// Initialize game rules
	game.GameRules.Init(game)

	return game, nil
}
