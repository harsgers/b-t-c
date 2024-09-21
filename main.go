package main

import (
	"fmt"
	"github.com/harrisongerst/b-t-c/player"
	"github.com/harrisongerst/b-t-c/scene"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	SCREEN_HEIGHT = 488
	SCREEN_WIDTH  = 488

	TILE_SIZE = 16
)

type Game struct {
	player *player.Player
	scene  *scene.Scene
}

func (g *Game) Update() error {
	//update player based on input
	//see if a turn has happened
	//if so then update other entities on screen
	//update map state
	g.scene.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.scene.Draw(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %1.2f", ebiten.ActualTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}
func main() {
	var layers [][]int
	baseMap := InitMap('0', '3', 8, 4)
	objectMap := InitOLayer(len(baseMap), baseMap)
	layers = append(layers, baseMap, objectMap)
	newPlayer := player.NewPlayer(20, 20)
	g := &Game{
		newPlayer,
		scene.NewScene(newPlayer, layers, nil),
	}

	ebiten.SetWindowSize(961, 961)
	ebiten.SetWindowTitle("beneath the castle")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
