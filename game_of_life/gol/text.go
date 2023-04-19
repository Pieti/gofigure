package gol

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
    GolFont font.Face
)

func init(){
    tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
    if err != nil {
        log.Fatal(err)
    }
    GolFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
        Size: 15,
        DPI:  72,
        Hinting: font.HintingVertical,
    })
    if err != nil {
        log.Fatal(err)
    }
}

func DrawCenteredText(screen *ebiten.Image, s string, font font.Face, cx, cy int, color color.Color) {
    bounds := text.BoundString(font, s)
    x, y := cx-bounds.Min.X-bounds.Dx()/2, cy-bounds.Min.Y-bounds.Dy()/2
    text.Draw(screen, s, font, x, y, color)
}
