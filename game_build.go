package gogame

import "errors"

func BuildGame(componentTable ComponentTable, gameRules GameRules, frontend Frontend) (*Game, error) {
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
		Frontend:       frontend,
	}

	// Test component table
	if err := game.ComponentTable.Validate(); err != nil {
		return nil, err
	}

	// Initialize game rules
	game.GameRules.Init(game)

	// Initialize frontend
	if game.Frontend != nil {
		game.Frontend.Init()
	}

	return game, nil
}
