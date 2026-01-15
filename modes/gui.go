package modes

import (
	"image/color"
	"math"
	"time"

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

var (
	bgColor        = color.RGBA{245, 230, 240, 255}
	boardColor     = color.RGBA{240, 180, 210, 255}
	emptyCellColor = color.RGBA{250, 240, 245, 255}
	player1Color   = color.RGBA{170, 210, 230, 255}
	player2Color   = color.RGBA{200, 80, 130, 255}
	outlineColor   = color.RGBA{120, 90, 110, 255}
	hoverLineColor = color.RGBA{255, 255, 255, 90}
)

var (
	topGoXFont  font.Face = basicfont.Face7x13
	topTextFace           = text.NewGoXFace(topGoXFont)
)

type GUI struct {
	g       *game.Connect4
	playerO *player.Min_max_player
	playerX *player.Min_max_player

	lastMouse bool

	pendingAIMove     bool
	waitDrawAfterMove bool

	delay    time.Duration
	last_move time.Time

	w, h   int
	boardX int
	boardY int
	radius int
}

func NewGUI(g *game.Connect4, playerO, playerX *player.Min_max_player) *GUI {
	return &GUI{
		g:       g,
		playerO: playerO,
		playerX: playerX,
		w:       padding*2 + g.Width()*cellSize,
		h:       padding*2 + topBar + g.Height()*cellSize,
		boardX:  padding,
		boardY:  padding + topBar,
		radius:  cellSize / 3,
	}
}

func RunGUI(ui *GUI, title string) {
	ebiten.SetWindowSize(ui.w, ui.h)
	ebiten.SetWindowTitle(title)
	if err := ebiten.RunGame(ui); err != nil {
		panic(err)
	}
}

func (ui *GUI) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyR) {
		ui.g = game.StartNewGame(ui.g.Width(), ui.g.Height())
		ui.pendingAIMove = false
		ui.waitDrawAfterMove = false
		ui.last_move = time.Time{}
		ui.lastMouse = false
		return nil
	}

	if ui.g.Game_over {
		return nil
	}

	if ui.waitDrawAfterMove {
		return nil
	}

	if ui.delay > 0 && !ui.isHumanTurn() {
		if ui.last_move.IsZero() {
			ui.last_move = time.Now()
			return nil
		}
		if time.Since(ui.last_move) < ui.delay {
			return nil
		}
	}

	if ui.pendingAIMove {
		ui.pendingAIMove = false
		if !ui.g.Game_over {
			_ = ui.tryAIMove()
			ui.waitDrawAfterMove = true
			if ui.delay > 0 {
				ui.last_move = time.Now()
			}
			if !ui.g.Game_over && !ui.isHumanTurn() && ui.delay > 0 {
				ui.pendingAIMove = true
			}
		}
		return nil
	}

	if ui.isHumanTurn() {
		mouseDown := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
		if mouseDown && !ui.lastMouse {
			mx, my := ebiten.CursorPosition()
			if col, ok := ui.pickColumn(mx, my); ok && ui.isAllowed(col) {
				ui.applyMove(col)
				if !ui.g.Game_over {
					ui.waitDrawAfterMove = true
					if !ui.isHumanTurn() {
						if ui.delay > 0 {
							ui.last_move = time.Now()
						}
						ui.pendingAIMove = true
					}
				}
			}
		}
		ui.lastMouse = mouseDown
		return nil
	}

	_ = ui.tryAIMove()
	ui.waitDrawAfterMove = true
	if ui.delay > 0 {
		ui.last_move = time.Now()
	}
	if !ui.g.Game_over && !ui.isHumanTurn() && ui.delay > 0 {
		ui.pendingAIMove = true
	}
	return nil
}

func (ui *GUI) applyMove(col int) {
	ui.g = ui.g.Drop_piece(col)
	ui.g = ui.g.Switch_player()
}

func (ui *GUI) tryAIMove() bool {
	var bot *player.Min_max_player

	switch ui.g.Who_moves {
	case game.O:
		bot = ui.playerO
	case game.X:
		bot = ui.playerX
	}

	if bot == nil {
		return false
	}

	move := bot.Decide(*ui.g)
	if !ui.isAllowed(move) {
		pd := ui.g.Possible_drops()
		if len(pd) == 0 {
			return true
		}
		move = pd[0]
	}

	ui.applyMove(move)
	return true
}

func (ui *GUI) Draw(screen *ebiten.Image) {
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

	if !ui.g.Game_over && ui.isHumanTurn() && !ui.pendingAIMove {
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

	if ui.waitDrawAfterMove {
		ui.waitDrawAfterMove = false
	}
}

func (ui *GUI) isHumanTurn() bool {
	if ui.g.Who_moves == game.O && ui.playerO != nil {
		return false
	}
	if ui.g.Who_moves == game.X && ui.playerX != nil {
		return false
	}
	return true
}

func drawTopText(screen *ebiten.Image, ui *GUI) {
	var msg string
	if ui.g.Game_over {
		msg = "GAME OVER WINNER IS: " + game.Cell_to_string(ui.g.Winner) + "   (R = restart)"
	} else {
		msg = "TURN: " + game.Cell_to_string(ui.g.Who_moves)
	}

	scale := 2.0
	textWidth := font.MeasureString(topGoXFont, msg).Ceil()
	x := float64((ui.w - int(float64(textWidth)*scale)) / 2)

	op := &text.DrawOptions{}
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(x, 25)
	op.ColorScale.ScaleWithColor(player2Color)
	text.Draw(screen, msg, topTextFace, op)
}

func (ui *GUI) Layout(_, _ int) (int, int) { return ui.w, ui.h }

func (ui *GUI) pickColumn(mx, my int) (int, bool) {
	bw := ui.g.Width() * cellSize
	bh := ui.g.Height() * cellSize
	if mx < ui.boardX || mx >= ui.boardX+bw || my < ui.boardY || my >= ui.boardY+bh {
		return 0, false
	}
	return (mx - ui.boardX) / cellSize, true
}

func (ui *GUI) isAllowed(col int) bool {
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
