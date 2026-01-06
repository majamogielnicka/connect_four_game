package modes

import (
	"connect_four/game"
	"connect_four/player"
	"time"
)

func AI_vs_AI_GUI(depthX, depthO int) {
	g := game.StartNewGame(7, 6)
	aiX := player.NewMinMaxPlayer(game.X, depthX)
	aiO := player.NewMinMaxPlayer(game.O, depthO)

	ui := NewGUI(g, aiO, aiX)
	ui.delay = 3 * time.Second
	ui.last_move = time.Now()

	RunGUI(ui, "Connect Four – AI vs AI")
}
