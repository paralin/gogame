package gogame

import "errors"

func BuildGame(gameSettings GameSettings, componentTable ComponentTable, gameRules GameRules, frontend Frontend) (*Game, error) {
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
		GameSettings:   gameSettings,
	}
	game.GameState = &GameState{
		game: game,
	}

	// Test component table
	if err := game.ComponentTable.Validate(); err != nil {
		return nil, err
	}

	// Initialize game rules
	game.GameRules.Init(game)

	// Initialize frontend
	if game.Frontend != nil {
		rules := game.Frontend.Init()
		if rules != nil {
			game.FrontendGameRules = rules
			rules.Init()
		}
	}

	game.setOperatingMode(GameOperatingMode_LOCAL)
	return game, nil
}
