package player

import (
	"bytes"
	"github.com/AhEhIOhYou/project2/prj2/internal/actor"
	"github.com/AhEhIOhYou/project2/prj2/internal/resources/images"
	"github.com/AhEhIOhYou/project2/prj2/internal/sprite"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"math"
)

const (
	initPlayerSpeed  = 5
	focusPlayerSpeed = 1
)

// Player represents player of the game
type Player struct {
	sprite *sprite.Sprite
	actor  *actor.Actor
	vx     float64
	vy     float64
	degree int
}

// New returns initialized Player
func New() *Player {
	p := &Player{}

	img, _, _ := image.Decode(bytes.NewReader(images.P_ROBO1))
	sp := sprite.New(&img, 8)
	p.sprite = sp

	p.actor = &actor.Actor{}
	p.actor.SetPosition(120, 160)
	p.actor.SetSpeed(initPlayerSpeed)

	return p
}

// Draw draws this sprite
func (p *Player) Draw(screen *ebiten.Image) {
	p.sprite.SetPosition(p.actor.X, p.actor.Y)
	// Calc frame-sprite index by degree
	spriteIndex := int(float64(((p.actor.Deg+90)+360)%360) / 45)
	p.sprite.SetIndex(spriteIndex)
	p.sprite.Draw(screen)
}

func (p *Player) Move(horizontal float64, vertical float64, isFire, isFocus bool) {

	if isFocus != false {
		p.actor.SetSpeed(focusPlayerSpeed)
	} else {
		p.actor.SetSpeed(initPlayerSpeed)
	}

	if vertical != 0 {
		p.vy = vertical * p.actor.Speed
		p.actor.Y = p.actor.Y + p.vy
	}

	if horizontal != 0 {
		p.vx = horizontal * p.actor.Speed
		p.actor.X = p.actor.X + p.vx
	}

	if vertical != 0 || horizontal != 0 {
		degree := int(math.Atan2(vertical, horizontal) * 180 / math.Pi)
		if isFire == false {
			p.actor.SetDeg(degree)
		}
	}
}
