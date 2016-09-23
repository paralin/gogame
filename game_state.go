package gogame

import "time"

type GameState struct {
	game         *Game
	updateTicker *time.Ticker
	running      bool
}

func (gs *GameState) Start() {
	if gs.game.GameSettings.Tick <= 0 {
		gs.game.GameSettings.Tick = 30
	}
	period := time.Duration(1000/gs.game.GameSettings.Tick) * time.Millisecond
	gs.updateTicker = time.NewTicker(period)
	gs.running = true
	go gs.tickThread()
}

func (gs *GameState) tickThread() {
	for gs.running {
		// Wait for next tick
		<-gs.updateTicker.C

		// Tick
	}
	gs.updateTicker.Stop()
	gs.updateTicker = nil
}

func (gs *GameState) Stop() {
	gs.running = false
}
