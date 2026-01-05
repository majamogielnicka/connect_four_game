package player

import (
	"connect_four/game"
	"fmt"
)

func Heuristics (g *game.Connect4, player game.Cell) int {
	count:=0
	center := g.Center_column()
	for _,x := range center {
		if x == player {
			count += 2
		} else if x == game.Empty {
			count += 0
		} else {
			count -= 3
		}
	}
	count*=4
	fours:=g.Iter_fours()
	for _, four := range fours {
		count+=Count(player, four)
		var opponent game.Cell
		if player==1{
			opponent=2
		} else {
			opponent=1
		}
		count-=Count(opponent, four)
	}
	fmt.Print(count)
	return count
}

func Count(searched game.Cell, four []game.Cell) int {
	counter := 0
	for _, v := range four {
		if v == searched {
			counter++
		}
	}
	return counter
}