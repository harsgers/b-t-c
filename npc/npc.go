package npc

import (
	"bytes"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/harrisongerst/b-t-c/assets"
)

type NPC struct {
	PosIndex  int
	Health    int
	Genus     string
	IsHostile bool
}

var npcImage *ebiten.Image

const (
	SCREEN_HEIGHT = 488
	SCREEN_WIDTH  = 488

	TILE_SIZE = 16
)

func init() {
	var err error

	img, _, err := image.Decode(bytes.NewBuffer(assets.Baddie_png))
	if err != nil {
		log.Fatal(err)
	}
	npcImage = ebiten.NewImageFromImage(img)
}
func NewNPC(position int, health int, genus string, isHostile bool) *NPC {
	npc := &NPC{
		PosIndex:  position,
		Health:    health,
		Genus:     genus,
		IsHostile: isHostile,
	}
	return npc
}

func (npc *NPC) Draw(screen *ebiten.Image) {
	const xCount = SCREEN_WIDTH / TILE_SIZE

	pOps := &ebiten.DrawImageOptions{}
	pOps.GeoM.Translate((float64(npc.PosIndex%xCount) * TILE_SIZE), float64((npc.PosIndex/xCount)*TILE_SIZE))
	screen.DrawImage(npcImage, pOps)

}
