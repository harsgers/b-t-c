package player

import (
	"bytes"
	"fmt"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/harrisongerst/b-t-c/assets"
)

const (
	SCREEN_HEIGHT = 488
	SCREEN_WIDTH  = 488

	TILE_SIZE = 16
)

type Player struct {
	PosIndex int
	Health   int
}

var playerImage *ebiten.Image

func init() {
	var err error

	img, _, err := image.Decode(bytes.NewBuffer(assets.MyGuy_png))
	if err != nil {
		log.Fatal(err)
	}
	playerImage = ebiten.NewImageFromImage(img)
}

func NewPlayer(startingPosition, startingHealth int) *Player {
	return &Player{
		PosIndex: startingPosition,
		Health:   startingHealth,
	}
}

func (p *Player) HandleMovement(keyPressed ebiten.Key, mapData [][]int) error {
	newPlayerPosition := p.PosIndex

	switch keyPressed {
	case ebiten.KeyDown:
		if mapData[0][p.Below()] == 0 {
			newPlayerPosition = p.Below()
		}
	case ebiten.KeyUp:
		if mapData[0][p.Above()] == 0 {
			newPlayerPosition = p.Above()
		}
	case ebiten.KeyLeft:
		if mapData[0][p.Left()] == 0 {
			newPlayerPosition = p.Left()
		}
	case ebiten.KeyRight:
		if mapData[0][p.Right()] == 0 {
			newPlayerPosition = p.Right()
		}
	default:
		newPlayerPosition = p.PosIndex
	}
	p.PosIndex = newPlayerPosition

	return nil
}

// TODO: it probably makes more sense to pass the player state down into the scene or map or whatever and have all actions handled there
func (p *Player) Update(mapData [][]int) (bool, error) {
	didMove := false
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		switch mapData[1][p.Below()] {
		case 3:
			//handle combat
			fmt.Printf("encountered enemy")

		default:
			p.HandleMovement(ebiten.KeyDown, mapData)
			didMove = true
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		switch mapData[1][p.Above()] {
		case 3:
			//handle combat
			fmt.Printf("encountered enemy")

			p.HandleMovement(ebiten.KeyUp, mapData)
			didMove = true
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
			p.HandleMovement(ebiten.KeyLeft, mapData)
			didMove = true
		}

		if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
			p.HandleMovement(ebiten.KeyRight, mapData)
			didMove = true
		}

		return didMove, nil
	}
	return true, nil
}

func (p *Player) Draw(screen *ebiten.Image) {
	const xCount = SCREEN_WIDTH / TILE_SIZE

	pOps := &ebiten.DrawImageOptions{}
	pOps.GeoM.Translate((float64(p.PosIndex%xCount) * TILE_SIZE), float64((p.PosIndex/xCount)*TILE_SIZE))
	screen.DrawImage(playerImage, pOps)

}

func (p *Player) isPathBlocked(mapData []int) {

}

func (p *Player) Above() int {
	return p.PosIndex - 30
}
func (p *Player) Below() int {
	return p.PosIndex + 30
}
func (p *Player) Right() int {
	return p.PosIndex + 1
}
func (p *Player) Left() int {
	return p.PosIndex - 1
}
