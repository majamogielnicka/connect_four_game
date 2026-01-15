package player

import (
	"connect_four/game"
	"fmt"
	"math"
)

const WIN = 1_000_000_000
const INF = math.MaxInt / 2

type Min_max_player struct {
	depth int
	piece game.Cell
}

func NewMinMaxPlayer(piece game.Cell, depth int) *Min_max_player {
	return &Min_max_player{
		piece: piece,
		depth: depth,
	}
}

func (p Min_max_player) algorithm(g *game.Connect4, d int, maximizing bool, alpha int, beta int) int {
	if g.Game_over == true {
		if g.Winner == p.piece {
			return WIN+d
		} else if g.Winner == game.Empty {
			return 0
		} else {
			return -WIN-d
		}
	} else if d == 0 {
		return Heuristics(g, p.piece)
	} else if maximizing {
		val := -INF
		for _, move := range g.Possible_drops() {
			new_state := g.Clone()
			new_state = new_state.Drop_piece(move)
			new_state = new_state.Switch_player()

			score := p.algorithm(new_state, d-1, false, alpha, beta)
			if score > val {
				val = score
			}
			if val > alpha {
				alpha = val
			}
			if alpha >= beta {
				break
			}
		}
		return val
	} else {
		val := INF
		for _, move := range g.Possible_drops() {
			new_state := g.Clone()
			new_state = new_state.Drop_piece(move)
			new_state = new_state.Switch_player()

			score := p.algorithm(new_state, d-1, true, alpha, beta)
			if score < val {
				val = score
			}
			if val < beta {
				beta = val
			}
			if beta <= alpha {
				break
			}
		}
		return val
	}
}

func (p Min_max_player) Decide(g game.Connect4) int {
	if g.Who_moves != p.piece {
		fmt.Println("not my round")
		return 0
	}

	type ms struct {
		move  int
		score int
	}

	var all []ms

	for _, move := range g.Possible_drops() {
		new_state := g.Clone()
		new_state = new_state.Drop_piece(move)
		new_state = new_state.Switch_player()

		score := p.algorithm(new_state, p.depth-1, false, math.MinInt/2, math.MaxInt/2)
		all = append(all, ms{move: move, score: score})
	}

	bestScore := -INF
	for _, x := range all {
		if x.score > bestScore {
			bestScore = x.score
		}
	}

	ties := 0
	for _, x := range all {
		if x.score == bestScore {
			ties++
		}
	}

	if ties > 1 {
		weights := []int{0, 1, 2, 3, 2, 1, 0}
		for i := range all {
			all[i].score += weights[all[i].move]
		}
	}

	bestMove := -1
	bestScore = -INF
	for _, x := range all {
		if bestMove == -1 || x.score > bestScore {
			bestMove = x.move
			bestScore = x.score
		}
	}
	
	//for _, x := range all {
	//	fmt.Println("move:", x.move, "score:", x.score)
	//}
	//fmt.Println("AI chose", bestMove, "score", bestScore)

	if bestMove == -1 {
		return 0
	}
	return bestMove
}
