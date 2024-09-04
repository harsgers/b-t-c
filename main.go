package main

import (
	"bytes"
	"fmt"
	"github.com/harrisongerst/b-t-c/assets"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	SCREEN_HEIGHT = 490
	SCREEN_WIDTH  = 490

	TILE_SIZE = 16
)

var tilesImage *ebiten.Image
var playerImage *ebiten.Image

func init() {
	var err error
	img, _, err := image.Decode(bytes.NewReader(assets.Shore_png))

	p_img, _, err := image.Decode(bytes.NewReader(assets.MyGuy_png))
	if err != nil {
		log.Fatal(err)
	}
	tilesImage = ebiten.NewImageFromImage(img)
	playerImage = ebiten.NewImageFromImage(p_img)
}

type Coords struct {
	X float64
	Y float64
}
type Game struct {
	layers         [][]int
	playerPosition Coords
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		g.playerPosition.Y += 1 * TILE_SIZE
		fmt.Printf("coords: %v \n", g.playerPosition)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		g.playerPosition.Y -= 1 * TILE_SIZE
		fmt.Printf("coords: %v \n", g.playerPosition)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		g.playerPosition.X -= 1 * TILE_SIZE
		fmt.Printf("coords: %v \n", g.playerPosition)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		g.playerPosition.X += 1 * TILE_SIZE
		fmt.Printf("coords: %v \n", g.playerPosition)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	w := tilesImage.Bounds().Dx()
	tileXCount := w / TILE_SIZE

	const xCount = SCREEN_WIDTH / TILE_SIZE

	for _, l := range g.layers {
		for i, t := range l {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate((float64(i%xCount) * TILE_SIZE), float64((i/xCount)*TILE_SIZE))
			sx := (t % tileXCount) * TILE_SIZE
			sy := (t / tileXCount) * TILE_SIZE
			screen.DrawImage(tilesImage.SubImage(image.Rect(sx, sy, sx+TILE_SIZE, sy+TILE_SIZE)).(*ebiten.Image), op)
		}
	}
	pOps := &ebiten.DrawImageOptions{}
	pOps.GeoM.Translate(g.playerPosition.X, g.playerPosition.Y)
	screen.DrawImage(playerImage, pOps)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %1.2f", ebiten.ActualTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func main() {

	g := &Game{
		playerPosition: Coords{X: 5 * TILE_SIZE, Y: 5 * TILE_SIZE},
		layers:         [][]int{InitMap('0', '3', 8, 4)},
	}
	ebiten.SetWindowSize(961, 961)
	ebiten.SetWindowTitle("beneath the castle")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
