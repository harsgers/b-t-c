package scene

import (
	"bytes"
	"github.com/harrisongerst/b-t-c/assets"
	"github.com/harrisongerst/b-t-c/npc"
	"github.com/harrisongerst/b-t-c/player"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
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

type Scene struct {
	player        *player.Player
	screenMap     [][]int
	npcList       []*npc.NPC
	prevPlayerPos int
}

func NewScene(p *player.Player, sMap [][]int, npcList []*npc.NPC) *Scene {
	s := &Scene{
		player:    p,
		screenMap: sMap,
		npcList:   npcList,
	}
	return s
}

func (s *Scene) Update() {
	didMove, _ := s.player.Update(s.screenMap)
	log.Print(didMove)
}

func (s *Scene) Draw(screen *ebiten.Image) {
	w := tilesImage.Bounds().Dx()
	tileXCount := w / TILE_SIZE
	const xCount = SCREEN_WIDTH / TILE_SIZE

	for m, l := range s.screenMap {
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
	for _, npc := range s.npcList {
		npc.Draw(screen)
	}
	s.player.Draw(screen)
}
