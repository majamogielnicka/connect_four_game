package modes

import (
	"fmt"
	"connect_four/game"
)

func PvP (){
	g := game.StartNewGame(7,6)
	fmt.Println("------PLAYER vs PLAYER mode ----------")
	var players_move int
	for {
		g.Draw()
		if g.Game_over ==true{
			fmt.Println("----GAME OVER-----")
			fmt.Println("The winner is ", game.Cell_to_string(g.Winner))
			break
		} else {
  		fmt.Scan(&players_move)
		g=g.Drop_piece(players_move)
		}
	}
}