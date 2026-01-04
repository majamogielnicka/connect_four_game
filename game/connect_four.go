package game

import "fmt"

type cell int

const (
	empty cell = iota
	o
	x
)

func Switch_player (player cell) cell{
	if player==x{
		return o
	} else {
		return x
	}
}

type Connect4 struct{
	width int
	height int 
	Who_moves cell 
	Game_over bool
	Winner cell
	board [][]cell 
}

func StartNewGame(width, height int) *Connect4 {
	g := &Connect4{
		width:     width,
		height:    height,
		Who_moves: o,
		Game_over: false,
		board:     make([][]cell, height),
	}
	for r := 0; r < height; r++ {
		g.board[r] = make([]cell, width)
	}
	return g
}

func (g Connect4) possible_drops () []int{ 
	var possible_drops []int
	for i:=0; i<g.width; i++ {
		if g.board[0][i]==empty {
			possible_drops = append(possible_drops, i)
		}
	}
	return possible_drops //slice
}

func (g Connect4) Drop_piece(column int) *Connect4{
	if g.Game_over {
		//TODO game exception
		fmt.Println("GAME OVER")
	}
	if n_is_in_list(column, g.possible_drops())==false {
		//TODO game exception
		fmt.Print("Invalid move")
	}
	n_row := g.height - 1
	for n_row >= 0 && g.board[n_row][column] != empty {
		n_row--
	}
	if n_row < 0 {
		fmt.Println("Invalid move")
		return &g
	}

	g.board[n_row][column] = g.Who_moves
	g.Game_over=g.check_game_over()
	g.Who_moves=Switch_player(g.Who_moves)
	return &g
}

func (g Connect4) center_column() []cell{
	col:=g.width/2
	center:=make([]cell, g.height)//slice
	for row := 0; row < g.height; row++ {
		center[row] = g.board[row][col]
	}
	fmt.Println(center)
	return center
}

func (g *Connect4) iter_fours() [][]cell {
	fours := make([][]cell, 0)

	// horizontal
	for row := 0; row < g.height; row++ {
		for col := 0; col <= g.width-4; col++ {
			four := make([]cell, 4)
			copy(four, g.board[row][col:col+4])
			fours = append(fours, four)
		}
	}

	// vertical
	for col := 0; col < g.width; col++ {
		for row := 0; row <= g.height-4; row++ {
			four := make([]cell, 4)
			for i := 0; i < 4; i++ {
				four[i] = g.board[row+i][col]
			}
			fours = append(fours, four)
		}
	}

	// diagonal 
	for row := 0; row <= g.height-4; row++ {
		for col := 0; col <= g.width-4; col++ {
			four := make([]cell, 4)
			for i := 0; i < 4; i++ {
				four[i] = g.board[row+i][col+i]
			}
			fours = append(fours, four)
		}
	}

	// diagonal 
	for row := 0; row <= g.height-4; row++ {
		for col := 0; col <= g.width-4; col++ {
			four := make([]cell, 4)
			for i := 0; i < 4; i++ {
				four[i] = g.board[row+3-i][col+i]
			}
			fours = append(fours, four)
		}
	}

	return fours
}

func (g *Connect4) check_game_over() bool {
	if len(g.possible_drops()) == 0 {
		g.Game_over = true
		fmt.Println("tie")
		return true
	}

	for _, four := range g.iter_fours() {
		if same_value(four, o) {
			g.Game_over = true
			g.Winner=o
			return true
		}
		if same_value(four, x) {
			g.Game_over = true
			g.Winner=x
			return true
		}
		same_value(four, x)
	}

	g.Game_over = false
	return false
}

func same_value(four []cell, who cell) bool {
	return len(four) == 4 &&
		four[0] == who &&
		four[1] == who &&
		four[2] == who &&
		four[3] == who
}

func (g *Connect4) Draw() {
	for i := range g.width {
		fmt.Print(" ",i, " ")
	}
	fmt.Println()
	for _, row := range g.board {
		fmt.Print("[")
		for i, c := range row {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(Cell_to_string(c))
		}
		fmt.Println("]")
	}

	if g.Game_over {
		return
	} else {
		fmt.Println("now moves:", Cell_to_string(g.Who_moves))
		fmt.Println("possible drops:", g.possible_drops())
	}
}

func Cell_to_string(c cell) string {
	switch c {
	case o:
		return "o"
	case x:
		return "x"
	default:
		return "_"
	}
}

func n_is_in_list (n int, elements []int) bool {
	for i := range elements {
		if elements[i] == n {
			return true
		}
	}
	return false
}