package snake

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Input struct{}

func NewInput() *Input {
    return &Input{}
}


func (i *Input) GetDirection() (Direction, bool) {
    if inpututil.IsKeyJustReleased(ebiten.KeyArrowUp) {
        return Up, true
    }
    if inpututil.IsKeyJustReleased(ebiten.KeyArrowLeft) {
        return Left, true
    }
    if inpututil.IsKeyJustReleased(ebiten.KeyArrowRight) {
        return Right, true
    }
    if inpututil.IsKeyJustReleased(ebiten.KeyArrowDown) {
        return Down, true
    }

    return 0, false
}
