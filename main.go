package main

import (
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)
var img *ebiten.Image

func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("assets/Shore.png")
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	screen.DrawImage(img, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("beneath the castle")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}