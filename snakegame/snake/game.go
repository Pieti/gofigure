package snake

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
    ScreenWidth = 600
    ScreenHeight = 600
)

var (
    backgroundColor = color.RGBA{200, 200, 200, 255}
    snakeColor = color.RGBA{200, 0, 150, 255}
    foodColor = color.RGBA{200, 200, 50, 150}
)

type Direction int

const (
    Up Direction = iota
    Right
    Down
    Left
)

type Position struct {
    x, y int
}

type Game struct{
    input *Input
    board *Board
}

func NewGame() *Game {
    return &Game{
        input: NewInput(),
        board: NewBoard(20, 20),
    }
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
    return ScreenWidth, ScreenHeight
}

func (g *Game) Update() error {
    return g.board.Update(g.input)
}

func (g *Game) Draw(screen *ebiten.Image) {
    screen.Fill(backgroundColor)
    if g.board.gameOver {
        ebitenutil.DebugPrint(screen, fmt.Sprintf("Game Over. Score: %d", g.board.score))
    } else {
        screenWidth, screenHeight := screen.Size()
        rowHeight := screenHeight / g.board.rows
        columnWidth := screenWidth / g.board.cols


        for _, p := range g.board.snake.body{
            ebitenutil.DrawRect(screen, float64(p.y*rowHeight), float64(p.x*columnWidth), float64(rowHeight), float64(columnWidth), snakeColor)
        }

        ebitenutil.DrawRect(screen, float64(g.board.food.y*rowHeight), float64(g.board.food.x*columnWidth), float64(rowHeight), float64(columnWidth), foodColor)
        ebitenutil.DebugPrint(screen, fmt.Sprintf("Score: %d", g.board.score))
    }
}
