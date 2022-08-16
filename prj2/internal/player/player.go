package player

import (
	"bytes"
	"github.com/AhEhIOhYou/project2/prj2/internal/actor"
	"github.com/AhEhIOhYou/project2/prj2/internal/input"
	"github.com/AhEhIOhYou/project2/prj2/internal/sprite"
	"github.com/AhEhIOhYou/project2/prj2/resources/images"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
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
	p.sprite.Draw(screen)
}

// Update updates the player's state
func (p *Player) Update(input *input.GameInput) {

	if input.Focus != false {
		p.actor.SetSpeed(focusPlayerSpeed)
	} else {
		p.actor.SetSpeed(initPlayerSpeed)
	}

	isMoving := false

	// Up
	if input.Up != 0 && (input.Right+input.Left == 0 || input.Right+input.Left == 2) {
		p.vx = 0
		p.vy = -p.actor.Speed
		p.actor.Y = p.actor.Y - p.actor.Speed
		isMoving = true
	}

	// Down
	if input.Down != 0 && (input.Right+input.Left == 0 || input.Right+input.Left == 2) {
		p.vx = 0
		p.vy = p.actor.Speed
		p.actor.Y = p.actor.Y + p.actor.Speed
		isMoving = true
	}

	// Left
	if input.Left != 0 && (input.Up+input.Down == 0 || input.Up+input.Down == 2) {
		p.vx = -p.actor.Speed
		p.vy = 0
		p.actor.X = p.actor.X - p.actor.Speed
		isMoving = true
	}

	// Right
	if input.Right != 0 && (input.Up+input.Down == 0 || input.Up+input.Down == 2) {
		p.vx = p.actor.Speed
		p.vy = 0
		p.actor.X = p.actor.X + p.actor.Speed
		isMoving = true
	}

	// Diagonal
	if isMoving == false {
		if input.Up != 0 && input.Right != 0 {
			p.vx = p.actor.NSpd
			p.vy = -p.actor.NSpd
			p.actor.X = p.actor.X + p.actor.NSpd
			p.actor.Y = p.actor.Y - p.actor.NSpd
			isMoving = true
		}
		if input.Up != 0 && input.Left != 0 {
			p.vx = -p.actor.NSpd
			p.vy = -p.actor.NSpd
			p.actor.X = p.actor.X - p.actor.NSpd
			p.actor.Y = p.actor.Y - p.actor.NSpd
			isMoving = true
		}
		if input.Down != 0 && input.Right != 0 {
			p.vx = p.actor.NSpd
			p.vy = p.actor.NSpd
			p.actor.X = p.actor.X + p.actor.NSpd
			p.actor.Y = p.actor.Y + p.actor.NSpd
			isMoving = true
		}
		if input.Down != 0 && input.Left != 0 {
			p.vx = -p.actor.NSpd
			p.vy = p.actor.NSpd
			p.actor.X = p.actor.X - p.actor.NSpd
			p.actor.Y = p.actor.Y + p.actor.NSpd
			isMoving = true
		}
	}

	if input.Fire == false {
		p.actor.SetDirection(2)
	}
}
