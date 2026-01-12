package game

import "fmt"

type Cell int

const (
	Empty Cell = iota
	O
	X
)

func (g Connect4) Switch_player () *Connect4{
	if g.Who_moves==X{
		g.Who_moves=O
		return &g
	} else {
		g.Who_moves=X
		return &g
	}
}

type Connect4 struct{
	width int
	height int 
	Who_moves Cell 
	Game_over bool
	Winner Cell
	Board [][]Cell 
}

func StartNewGame(width, height int) *Connect4 {
	g := &Connect4{
		width:     width,
		height:    height,
		Who_moves: O,
		Game_over: false,
		Board:     make([][]Cell, height),
	}
	for r := 0; r < height; r++ {
		g.Board[r] = make([]Cell, width)
	}
	return g
}

func (g Connect4) Possible_drops () []int{ 
	var Possible_drops []int
	for i:=0; i<g.width; i++ {
		if g.Board[0][i]==Empty {
			Possible_drops = append(Possible_drops, i)
		}
	}
	return Possible_drops //slice
}

func (g Connect4) Drop_piece(column int) *Connect4{
	if g.Game_over {
		fmt.Println("GAME OVER")
		return &g
	}
	if n_is_in_list(column, g.Possible_drops())==false {
		fmt.Print("Invalid move")
		return &g
	}
	n_row := g.height - 1
	for n_row >= 0 && g.Board[n_row][column] != Empty {
		n_row--
	}
	if n_row < 0 {
		fmt.Println("Invalid move")
		return &g
	}

	g.Board[n_row][column] = g.Who_moves
	g.Game_over=g.check_game_over()
	return &g
}

func (g Connect4) Center_column() []Cell{
	col:=g.width/2
	center:=make([]Cell, g.height)//slice
	for row := 0; row < g.height; row++ {
		center[row] = g.Board[row][col]
	}
	return center
}

func (g *Connect4) Iter_fours() [][]Cell {
	fours := make([][]Cell, 0)

	// horizontal
	for row := 0; row < g.height; row++ {
		for col := 0; col <= g.width-4; col++ {
			four := make([]Cell, 4)
			copy(four, g.Board[row][col:col+4])
			fours = append(fours, four)
		}
	}

	// vertical
	for col := 0; col < g.width; col++ {
		for row := 0; row <= g.height-4; row++ {
			four := make([]Cell, 4)
			for i := 0; i < 4; i++ {
				four[i] = g.Board[row+i][col]
			}
			fours = append(fours, four)
		}
	}

	// diagonal 
	for row := 0; row <= g.height-4; row++ {
		for col := 0; col <= g.width-4; col++ {
			four := make([]Cell, 4)
			for i := 0; i < 4; i++ {
				four[i] = g.Board[row+i][col+i]
			}
			fours = append(fours, four)
		}
	}

	// diagonal 
	for row := 0; row <= g.height-4; row++ {
		for col := 0; col <= g.width-4; col++ {
			four := make([]Cell, 4)
			for i := 0; i < 4; i++ {
				four[i] = g.Board[row+3-i][col+i]
			}
			fours = append(fours, four)
		}
	}

	return fours
}

func (g *Connect4) check_game_over() bool {
	if len(g.Possible_drops()) == 0 {
		g.Game_over = true
		fmt.Println("tie")
		return true
	}

	for _, four := range g.Iter_fours() {
		if same_value(four, O) {
			g.Game_over = true
			g.Winner=O
			return true
		}
		if same_value(four, X) {
			g.Game_over = true
			g.Winner=X
			return true
		}
		same_value(four, X)
	}

	g.Game_over = false
	return false
}

func (g *Connect4) Clone() *Connect4 {
	ng := &Connect4{
		width:     g.width,
		height:    g.height,
		Who_moves: g.Who_moves,
		Game_over: g.Game_over,
		Winner:    g.Winner,
		Board:     make([][]Cell, g.height),
	}

	for r := 0; r < g.height; r++ {
		ng.Board[r] = make([]Cell, g.width)
		copy(ng.Board[r], g.Board[r])
	}

	return ng
}

func same_value(four []Cell, who Cell) bool {
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
	for _, row := range g.Board {
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
		fmt.Println("possible drops:", g.Possible_drops())
	}
}

func Cell_to_string(c Cell) string {
	switch c {
	case O:
		return "O"
	case X:
		return "X"
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

func (g *Connect4) Width() int  { return g.width }
func (g *Connect4) Height() int { return g.height }

func (g *Connect4) CellAt(row, col int) Cell {
	return g.Board[row][col]
}