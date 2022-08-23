package actors

import (
	"github.com/AhEhIOhYou/project2/prj2/internal/fields"
	"github.com/AhEhIOhYou/project2/prj2/internal/objectpool"
	"github.com/AhEhIOhYou/project2/prj2/internal/sprite"
	"github.com/AhEhIOhYou/project2/prj2/internal/tools"
	"github.com/AhEhIOhYou/project2/prj2/internal/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

const (
	enemyPosX   = 0
	enemyPosY   = 100
	enemyWidth  = 10
	enemyHeight = 10
	enemySpeed  = 1
	enemyDegree = 0
)

type Enemy struct {
	Actor
	wep       tools.Weapon
	shotSpeed float64
	shotSize  float64
}

func NewEnemy(f *fields.Field, bulletPool *objectpool.Pool) *Enemy {
	e := &Enemy{Actor: *NewActor()}
	e.currField = f
	e.bulletPool = bulletPool
	return e
}

func (e *Enemy) Init(x, y, speed float64) {
	e.setSize(enemyWidth, enemyHeight)
	e.SetPosition(x, y)
	e.SetSpeed(speed, enemyDegree)
	e.isActive = true
	e.spr = sprite.Player
}

func (e *Enemy) Draw(screen *ebiten.Image) {
	e.spr.SetPosition(e.GetX(), e.GetY())
	e.spr.SetIndex(utils.DegreeToDirectionIndex(e.degree))
	e.spr.Draw(screen)
}

func (e *Enemy) Action(horizontal, vertical float64) {
	x := e.GetX()
	y := e.GetY()
	f := e.currField

	if vertical != 0 {
		e.vy = vertical * e.speed
		y = y + e.vy
		y = math.Max(f.GetTop()+e.GetHeight()/2, y)
		y = math.Min(f.GetBottom()-e.GetHeight()/2, y)
	}

	if horizontal != 0 {
		e.vx = horizontal * e.speed
		x = x + e.vx
		x = math.Max(f.GetLeft()+e.GetWidth()/2, x)
		x = math.Min(f.GetRight()-e.GetWidth()/2, x)
	}

	e.SetPosition(x, y)

	if vertical != 0 || horizontal != 0 {
		e.degree = utils.RadToDeg(math.Atan2(vertical, horizontal))
	}
}
