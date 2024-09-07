package player

import "github.com/hajimehoshi/ebiten/v2"

type Player struct {
	PosIndex int
	Health   int
}

func (p *Player) HandleMovement(keyPressed ebiten.Key) error {
	var newPlayerPosition int
	switch keyPressed {
	case ebiten.KeyDown:
		newPlayerPosition = p.PosIndex + 30
	case ebiten.KeyUp:
		newPlayerPosition = p.PosIndex - 30
	case ebiten.KeyLeft:
		newPlayerPosition = p.PosIndex - 1
	case ebiten.KeyRight:
		newPlayerPosition = p.PosIndex + 1
	default:
		newPlayerPosition = p.PosIndex
	}
	p.PosIndex = newPlayerPosition

	return nil
}
