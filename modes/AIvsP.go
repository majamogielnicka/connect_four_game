package modes

import (
	"fmt"
	"connect_four/game"
	"connect_four/player"
)

func AI_vs_P(depth int) {
	g := game.StartNewGame(7, 6)
	ai := player.NewMinMaxPlayer(game.X, depth)

	fmt.Println("------PLAYER vs AI mode ----------")
	var start int
	var players_move int
	fmt.Println("Who starts?: \n Type *1* for ai\n Type *2* for player\n")
	fmt.Scan(&start)
	if start == 1{
		g.Who_moves = game.X
	} else{
		g.Who_moves = game.O
	}

	for {
		g.Draw()
		if g.Game_over == true {
			fmt.Println("----GAME OVER-----")
			fmt.Println("The winner is", game.Cell_to_string(g.Winner))
			break
		}
		if g.Who_moves == game.O {
			fmt.Println("Players move!!!")

			if _, err := fmt.Scan(&players_move); err != nil {
				fmt.Println("\nINVALID INPUT, try once again:")
				continue
			}

			if players_move > 6 || players_move < 0{
				fmt.Println("\nINVALID INPUT, try once again: ")
				continue
			}

			g = g.Drop_piece(players_move)
			g = g.Switch_player()
		} else if g.Who_moves == game.X {
			g.Draw()
			fmt.Println("AI's Move")
			move := ai.Decide(*g)
			g = g.Drop_piece(move)
			g = g.Switch_player()
		}
	}
}
