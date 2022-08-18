package actors

import (
	"github.com/AhEhIOhYou/project2/prj2/internal/sprite"
	"github.com/hajimehoshi/ebiten/v2"
)

// PlayerBullet представление пули игрока
type PlayerBullet struct {
	Actor
	spriteIndex int
	isActive    bool
}

// NewPlayerShot возвращает инициализированную пулю
func NewPlayerShot() *PlayerBullet {
	actor := &Actor{}
	p := &PlayerBullet{Actor: *actor}
	p.isActive = false
	return p
}

// IsActive возвращает состояние пули
func (p *PlayerBullet) IsActive() bool {
	return p.isActive
}

// SetInactive делает пулю неактивной
func (p *PlayerBullet) SetInactive() {
	p.isActive = false
}

// Init инициализация пули
func (p *PlayerBullet) Init(degree int, speed float64, x, y, size int) {
	p.speed = speed
	p.deg = degree

	p.vx = 0
	p.vy = -speed

	p.spriteIndex = degreeToDirectionIndex(degree)

	p.width = size
	p.height = size
	p.x = float64(x)
	p.y = float64(y)

	p.isActive = true
}

// Draw отрисовка пули
func (p *PlayerBullet) Draw(screen *ebiten.Image) {
	sprite.PlayerBullet.SetPosition(p.x, p.y)
	sprite.PlayerBullet.SetIndex(p.spriteIndex)
	sprite.PlayerBullet.Draw(screen)
}

// Move движение пули
func (p *PlayerBullet) Move() {
	p.x = p.x + p.vx
	p.y = p.y + p.vy
	if p.isOutOfBoundary() {
		p.isActive = false
	}
}
