package actors

import (
	"github.com/AhEhIOhYou/project2/prj2/internal/sprite"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

const (
	initPlayerSpeed  = 5
	focusPlayerSpeed = 2
	initPositionX    = 120
	initPositionY    = 160
	playerWidth      = 10
	playerHeight     = 10
)

// Player представляет игрока
type Player struct {
	Actor
	life int
}

// NewPlayer возвращает инициализированного игрока
func NewPlayer() *Player {
	actor := &Actor{}
	p := &Player{Actor: *actor}

	p.width = playerWidth
	p.height = playerHeight
	p.setPosition(initPositionX, initPositionY)
	p.setSpeed(initPlayerSpeed)
	p.deg = 270

	return p
}

// Draw отрисовка игрока
func (p *Player) Draw(screen *ebiten.Image) {
	sprite.Player.SetPosition(p.x, p.y)
	sprite.Player.SetIndex(degreeToDirectionIndex(p.deg))
	sprite.Player.Draw(screen)
}

// Action меняет состояние игрока от выполняемых действий
func (p *Player) Action(horizontal float64, vertical float64, isFire, isFocus bool) {

	if isFocus != false {
		p.setSpeed(focusPlayerSpeed)
	} else {
		p.setSpeed(initPlayerSpeed)
	}

	if vertical != 0 {
		p.vy = vertical * p.speed
		p.y = p.y + p.vy
		p.y = math.Max(float64(bounds.GetTop()+p.height/2), p.y)
		p.y = math.Min(float64(bounds.GetBottom()-p.height/2), p.y)
	}

	if horizontal != 0 {
		p.vx = horizontal * p.speed
		p.x = p.x + p.vx
		p.x = math.Max(float64(bounds.GetLeft()+p.width/2), p.x)
		p.x = math.Min(float64(bounds.GetRight()-p.width/2), p.x)
	}

	if vertical != 0 || horizontal != 0 {
		p.deg = radToDeg(math.Atan2(vertical, horizontal))
	}
}
