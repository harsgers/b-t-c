package main

import (
	"bytes"
	"fmt"
	"github.com/harrisongerst/b-t-c/assets"
	"github.com/harrisongerst/b-t-c/player"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// HAVE EACH OBJECT HAVE AN UPDATE AND A DRAW
const (
	SCREEN_HEIGHT = 488
	SCREEN_WIDTH  = 488

	TILE_SIZE = 16
)

var tilesImage *ebiten.Image

func init() {
	var err error
	img, _, err := image.Decode(bytes.NewReader(assets.Wallandfloor_png))
	if err != nil {
		log.Fatal(err)
	}
	tilesImage = ebiten.NewImageFromImage(img)
}

type Coords struct {
	X float64
	Y float64
}
type Game struct {
	layers        [][]int
	player        *player.Player
	prevPlayerPos int
	//TODO: keep track of previous positions of all objects on screen (map state?)
}

func (g *Game) Update() error {
	//update player based on input
	//see if a turn has happened
	//if so then update other entities on screen
	//update map state
	didMove, _ := g.player.Update(g.layers[0])
	if didMove {
		g.updateMap()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		fmt.Printf("%v \n", g.layers[1])
		fmt.Printf("player: %v \n", g.player.PosIndex)
		fmt.Printf("prevplayer: %v \n", g.prevPlayerPos)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	w := tilesImage.Bounds().Dx()
	tileXCount := w / TILE_SIZE
	const xCount = SCREEN_WIDTH / TILE_SIZE

	for m, l := range g.layers {
		//render tile layer
		if m == 0 {
			for i, t := range l {
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate((float64(i%xCount) * TILE_SIZE), float64((i/xCount)*TILE_SIZE))
				sx := (t % tileXCount) * TILE_SIZE
				sy := (t / tileXCount) * TILE_SIZE
				screen.DrawImage(tilesImage.SubImage(image.Rect(sx, sy, sx+TILE_SIZE, sy+TILE_SIZE)).(*ebiten.Image), op)
			}
		}
	}

	g.player.Draw(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %1.2f", ebiten.ActualTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func (g *Game) updateMap() {
	g.layers[1][g.prevPlayerPos] = 0
	g.layers[1][g.player.PosIndex] = 1

	g.prevPlayerPos = g.player.PosIndex
}
func main() {

	baseMap := InitMap('0', '3', 8, 4)
	objectMap := InitOLayer(len(baseMap), baseMap)
	g := &Game{
		//TODO: add the objects/player layer to track state of pc position
		layers: [][]int{
			baseMap,
			objectMap,
		},
		player:        player.NewPlayer(20, 20),
		prevPlayerPos: 20,
	}
	ebiten.SetWindowSize(961, 961)
	ebiten.SetWindowTitle("beneath the castle")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
