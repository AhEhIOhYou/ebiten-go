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
	actor.Actor
	sprite *sprite.Sprite
	vx     float64
	vy     float64
	degree int
}

// New returns initialized Player
func New() *Player {

	actor := &actor.Actor{}
	p := &Player{Actor: *actor}

	img, _, _ := image.Decode(bytes.NewReader(images.P_ROBO1))
	sp := sprite.New(&img, 8)
	p.sprite = sp

	p.SetPosition(120, 160)

	p.SetSpeed(initPlayerSpeed)

	return p
}

// Draw draws this sprite
func (p *Player) Draw(screen *ebiten.Image) {
	p.sprite.SetPosition(p.X, p.Y)
	adjust := 22.5
	spriteIndex := int(float64(p.Deg)+90.0+360.0+adjust) % 360 / 45
	p.sprite.SetIndex(spriteIndex)
	p.sprite.Draw(screen)
}

func (p *Player) Move(horizontal float64, vertical float64, isFire, isFocus bool) {

	if isFocus != false {
		p.SetSpeed(focusPlayerSpeed)
	} else {
		p.SetSpeed(initPlayerSpeed)
	}

	if vertical != 0 {
		p.vy = vertical * p.Speed
		p.Y = p.Y + p.vy
	}

	if horizontal != 0 {
		p.vx = horizontal * p.Speed
		p.X = p.X + p.vx
	}

	if vertical != 0 || horizontal != 0 {
		degree := int(math.Atan2(vertical, horizontal) * 180 / math.Pi)
		p.SetDeg(degree)
	}
}
