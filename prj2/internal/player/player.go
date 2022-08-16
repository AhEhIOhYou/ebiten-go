package player

import (
	"bytes"
	"github.com/AhEhIOhYou/project2/prj2/internal/actor"
	"github.com/AhEhIOhYou/project2/prj2/internal/sprite"
	"github.com/AhEhIOhYou/project2/prj2/resources/images"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

// Player represents player of the game
type Player struct {
	sprite *sprite.Sprite
	actor  *actor.Actor
}

// New returns initialized Player
func New() *Player {
	p := &Player{}

	img, _, _ := image.Decode(bytes.NewReader(images.P_ROBO1))
	sp := sprite.New(&img, 8)
	p.sprite = sp

	p.actor = &actor.Actor{}
	p.actor.SetPosition(120, 160)

	return p
}

// Draw draws this sprite
func (p *Player) Draw(screen *ebiten.Image) {
	p.sprite.SetPosition(p.actor.X, p.actor.Y)
	p.sprite.Draw(screen)
}
