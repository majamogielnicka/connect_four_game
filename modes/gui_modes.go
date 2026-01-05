package modes

import (
	"image/color"
	"math"

	"connect_four/game"
	"connect_four/player"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
)


const (
	cellSize = 90
	padding  = 20
	topBar   = 90
)

// Pastel theme
var (
	bgColor        = color.RGBA{245, 230, 240, 255}
	boardColor     = color.RGBA{240, 180, 210, 255}
	emptyCellColor = color.RGBA{250, 240, 245, 255}
	player1Color   = color.RGBA{170, 210, 230, 255}
	player2Color   = color.RGBA{200, 80, 130, 255}
	outlineColor   = color.RGBA{120, 90, 110, 255}
	hoverLineColor = color.RGBA{255, 255, 255, 90}
)

func PvP_GUI()     { runGUI(false) }
func AI_vs_P_GUI() { runGUI(true) }

type guiUI struct {
	g         *game.Connect4
	vsAI      bool
	ai        *player.Min_max_player
	lastMouse bool

	w, h   int
	boardX int
	boardY int
	radius int
}

var (
	topGoXFont font.Face = basicfont.Face7x13
	topTextFace          = text.NewGoXFace(topGoXFont)
)

func runGUI(vsAI bool) {
	g := game.StartNewGame(7, 6)

	ui := &guiUI{
		g:      g,
		vsAI:   vsAI,
		w:      padding*2 + g.Width()*cellSize,
		h:      padding*2 + topBar + g.Height()*cellSize,
		boardX: padding,
		boardY: padding + topBar,
		radius: cellSize / 3,
	}

	if vsAI {
		ui.ai = player.NewMinMaxPlayer(game.X)
	}

	ebiten.SetWindowSize(ui.w, ui.h)
	if vsAI {
		ebiten.SetWindowTitle("Connect Four – AI vs Player")
	} else {
		ebiten.SetWindowTitle("Connect Four – PvP")
	}

	if err := ebiten.RunGame(ui); err != nil {
		panic(err)
	}
}

func (ui *guiUI) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyR) {
		ui.g = game.StartNewGame(7, 6)
		if ui.vsAI {
			ui.ai = player.NewMinMaxPlayer(game.X)
		}
		return nil
	}

	if ui.g.Game_over {
		return nil
	}

	if ui.vsAI && ui.g.Who_moves == game.X {
		move := ui.ai.Decide(*ui.g)
		if !ui.isAllowed(move) {
			pd := ui.g.Possible_drops()
			if len(pd) == 0 {
				return nil
			}
			move = pd[0]
		}
		ui.g = ui.g.Drop_piece(move)
		ui.g = ui.g.Switch_player()
		return nil
	}

	mouseDown := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	if mouseDown && !ui.lastMouse {
		if !ui.vsAI || ui.g.Who_moves == game.O {
			mx, my := ebiten.CursorPosition()
			if col, ok := ui.pickColumn(mx, my); ok && ui.isAllowed(col) {
				ui.g = ui.g.Drop_piece(col)
				ui.g = ui.g.Switch_player()
			}
		}
	}
	ui.lastMouse = mouseDown
	return nil
}

func (ui *guiUI) Draw(screen *ebiten.Image) {
	screen.Fill(bgColor)

	drawTopText(screen, ui)

	bw := ui.g.Width() * cellSize
	bh := ui.g.Height() * cellSize

	vector.FillRect(
		screen,
		float32(ui.boardX), float32(ui.boardY),
		float32(bw), float32(bh),
		boardColor,
		false,
	)

	if !ui.g.Game_over && (!ui.vsAI || ui.g.Who_moves == game.O) {
		mx, my := ebiten.CursorPosition()
		if col, ok := ui.pickColumn(mx, my); ok {
			x := float32(ui.boardX + col*cellSize)
			vector.FillRect(
				screen,
				x, float32(ui.boardY),
				3, float32(bh),
				hoverLineColor,
				false,
			)
		}
	}

	for r := 0; r < ui.g.Height(); r++ {
		for c := 0; c < ui.g.Width(); c++ {
			cx := ui.boardX + c*cellSize + cellSize/2
			cy := ui.boardY + r*cellSize + cellSize/2
			fill := cellColor(ui.g.CellAt(r, c))
			drawCircle(screen, cx, cy, ui.radius+2, outlineColor)
			drawCircle(screen, cx, cy, ui.radius, fill)
		}
	}
}

var topFont font.Face = basicfont.Face7x13
func drawTopText(screen *ebiten.Image, ui *guiUI) {
	var msg string
	if ui.g.Game_over {
		msg = "GAME OVER — WINNER: " + game.Cell_to_string(ui.g.Winner) + "   (R = restart)"
	} else {
		msg = "TURN: " + game.Cell_to_string(ui.g.Who_moves)
	}

	scale := 2.0

	// szerokość liczymy na GoX foncie
	textWidth := font.MeasureString(topGoXFont, msg).Ceil()
	x := float64((ui.w - int(float64(textWidth)*scale)) / 2)
	y := 28.0

	op := &text.DrawOptions{}
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(x, y)
	op.ColorScale.ScaleWithColor(player2Color)

	// a rysujemy już face'm z text/v2
	text.Draw(screen, msg, topTextFace, op)
}


func (ui *guiUI) Layout(_, _ int) (int, int) { return ui.w, ui.h }

func (ui *guiUI) pickColumn(mx, my int) (int, bool) {
	bw := ui.g.Width() * cellSize
	bh := ui.g.Height() * cellSize
	if mx < ui.boardX || mx >= ui.boardX+bw || my < ui.boardY || my >= ui.boardY+bh {
		return 0, false
	}
	return (mx - ui.boardX) / cellSize, true
}

func (ui *guiUI) isAllowed(col int) bool {
	for _, c := range ui.g.Possible_drops() {
		if c == col {
			return true
		}
	}
	return false
}

func cellColor(c game.Cell) color.Color {
	switch c {
	case game.O:
		return player1Color
	case game.X:
		return player2Color
	default:
		return emptyCellColor
	}
}

func drawCircle(dst *ebiten.Image, cx, cy, r int, col color.Color) {
	rr := float64(r)
	for y := -r; y <= r; y++ {
		for x := -r; x <= r; x++ {
			if math.Hypot(float64(x), float64(y)) <= rr {
				dst.Set(cx+x, cy+y, col)
			}
		}
	}
}
