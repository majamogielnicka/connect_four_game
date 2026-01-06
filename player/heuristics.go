package player

import (
	"connect_four/game"
)

func Heuristics(g *game.Connect4, player_id game.Cell) int {
	score := 0
	opp := get_opponent(player_id)

	center := g.Center_column()
	for _, x := range center {
		if x == player_id {
			score += 6
		} else if x == opp {
			score -= 6
		}
	}

	for _, four := range g.Iter_fours() {
		score += score_window(four, player_id, opp)
	}

	return score
}

func get_opponent(player_id game.Cell) game.Cell {
	if player_id == game.O {
		return game.X
	}
	return game.O
}

func Count(four []game.Cell, player_id, opp game.Cell) (player_count, opp_count, empty_count int) {
	for _, v := range four {
		switch v {
		case player_id:
			player_count++
		case opp:
			opp_count++
		default:
			empty_count++
		}
	}
	return
}

func score_window(four []game.Cell, player_id, opp game.Cell) int {
	x, o, e := Count(four, player_id, opp)

	if x > 0 && o > 0 {
		return 0
	}

	if x == 4 {
		return 10000
	}
	if o == 4 {
		return -10000
	}

	if x == 3 && e == 1 {
		return 6000
	}
	if x == 2 && e == 2 {
		return 250
	}
	if x == 1 && e == 3 {
		return 10
	}

	if o == 3 && e == 1 {
		return -6000
	}
	if o == 2 && e == 2 {
		return -250
	}
	if o == 1 && e == 3 {
		return -10
	}

	return 0
}
