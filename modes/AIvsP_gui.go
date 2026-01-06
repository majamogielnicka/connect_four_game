package modes

import (
	"connect_four/game"
	"connect_four/player"
)

func AI_vs_P_GUI(depth int) {
	g := game.StartNewGame(7, 6)
	aiX := player.NewMinMaxPlayer(game.X, depth)

	_ = depth

	ui := NewGUI(g, nil, aiX)
	RunGUI(ui, "Connect Four – AI vs Player")
}
