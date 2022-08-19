package actors

import (
	"github.com/AhEhIOhYou/project2/prj2/internal/sprite"
	"github.com/AhEhIOhYou/project2/prj2/internal/tools"
	"github.com/AhEhIOhYou/project2/prj2/internal/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

const (
	initPlayerSpeed  = 3
	focusPlayerSpeed = 1.5
	initPositionX    = 320
	initPositionY    = 320
	playerWidth      = 10
	playerHeight     = 10
)

// Player представляет игрока
type Player struct {
	Actor
	wep       tools.Weapon
	shotSpeed float64
	shotSize  float64
}

// NewPlayer возвращает инициализированного игрока
func NewPlayer() *Player {
	p := &Player{Actor: *NewActor()}
	return p
}

func (p *Player) Init() {
	p.life = 10
	p.setSize(10, 10)
	p.SetPosition(320, 320)
	p.SetSpeed(2, 270)
	p.isActive = true
	p.spr = sprite.Player
}

// Draw отрисовка игрока
func (p *Player) Draw(screen *ebiten.Image) {
	p.spr.SetPosition(p.GetX(), p.GetY())
	p.spr.SetIndex(utils.DegreeToDirectionIndex(p.degree))
	p.spr.Draw(screen)
}

// Action меняет состояние игрока от выполняемых действий
func (p *Player) Action(horizontal, vertical float64, isFire, isFocus bool) {
	x := p.GetX()
	y := p.GetY()
	f := p.currField

	if isFocus != false {
		p.SetSpeed(focusPlayerSpeed, 270)
	} else {
		p.SetSpeed(initPlayerSpeed, 270)
	}

	if vertical != 0 {
		p.vy = vertical * p.speed
		y = y + p.vy
		y = math.Max(f.GetTop()+p.GetHeight()/2, y)
		y = math.Min(f.GetBottom()-p.GetHeight()/2, y)
	}

	if horizontal != 0 {
		p.vx = horizontal * p.speed
		x = x + p.vx
		x = math.Max(f.GetLeft()+p.GetWidth()/2, x)
		x = math.Min(f.GetRight()-p.GetWidth()/2, x)
	}

	p.SetPosition(x, y)

	if vertical != 0 || horizontal != 0 {
		if isFire == false {
			p.degree = utils.RadToDeg(math.Atan2(vertical, horizontal))
		}
	}
}
