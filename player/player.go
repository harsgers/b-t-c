package player

import (
	"bytes"
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

func (p *Player) HandleMovement(keyPressed ebiten.Key, mapData []int) error {
	newPlayerPosition := p.PosIndex

	switch keyPressed {
	case ebiten.KeyDown:
		if mapData[p.PosIndex+30] != 2 {
			newPlayerPosition = p.PosIndex + 30
		}
	case ebiten.KeyUp:
		if mapData[p.PosIndex-30] != 2 {
			newPlayerPosition = p.PosIndex - 30
		}
	case ebiten.KeyLeft:
		if mapData[p.PosIndex-1] != 2 {
			newPlayerPosition = p.PosIndex - 1
		}
	case ebiten.KeyRight:
		if mapData[p.PosIndex+1] != 2 {
			newPlayerPosition = p.PosIndex + 1
		}
	default:
		newPlayerPosition = p.PosIndex
	}
	p.PosIndex = newPlayerPosition

	return nil
}

func (p *Player) Update(mapData []int) error {
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		p.HandleMovement(ebiten.KeyDown, mapData)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		p.HandleMovement(ebiten.KeyUp, mapData)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		p.HandleMovement(ebiten.KeyLeft, mapData)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		p.HandleMovement(ebiten.KeyRight, mapData)
	}

	return nil
}

func (p *Player) Draw(screen *ebiten.Image) {
	const xCount = SCREEN_WIDTH / TILE_SIZE

	pOps := &ebiten.DrawImageOptions{}
	pOps.GeoM.Translate((float64(p.PosIndex%xCount) * TILE_SIZE), float64((p.PosIndex/xCount)*TILE_SIZE))
	screen.DrawImage(playerImage, pOps)

}

func (p *Player) isPathBlocked(mapData []int) {

}
