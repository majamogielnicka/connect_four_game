package player

import (
	"connect_four/game"
	"fmt"
)

type Min_max_player struct {
	depth int
	piece game.Cell
}

func NewMinMaxPlayer(piece game.Cell) *Min_max_player {
	return &Min_max_player{
		piece: piece,
		depth: 5,
	}
}

func (p Min_max_player) algorithm(g *game.Connect4, d int, maximizing bool, alpha int, beta int) int {
	if g.Game_over == true {
		if g.Winner == p.piece {
			return 999
		} else if g.Winner == game.Empty {
			return 0
		} else {
			return -999
		}
	}

	if d == 0 {
		fmt.Println("heurystyka")
		return 0
	}

	if maximizing {
		val := -999999
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
		val := 999999
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

	moves := make(map[int]int)

	for _, move := range g.Possible_drops() {
		new_state := g.Clone()
		new_state = new_state.Drop_piece(move)
		new_state = new_state.Switch_player()
		moves[move] = p.algorithm(new_state, p.depth-1, false, -99999, 99999)
	}

	bestMove := -1
	bestScore := -999999
	for move, score := range moves {
		if bestMove == -1 || score > bestScore {
			bestMove = move
			bestScore = score
		}
	}

	if bestMove == -1 {
		return 0
	}
	return bestMove
}
