package modes

import "connect_four/game"

func PvP_GUI() {
	g := game.StartNewGame(7, 6)
	ui := NewGUI(g, nil, nil)
	RunGUI(ui, "Connect Four – PvP")
}
