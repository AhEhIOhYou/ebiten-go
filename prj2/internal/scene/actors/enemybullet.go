package actors

import (
	"github.com/AhEhIOhYou/project2/prj2/internal/sprite"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

// EnemyBullet представление пули игрока
type EnemyBullet struct {
	Actor
	spriteIndex int
	isActive    bool
}

// NewEnemyShot возвращает инициализированную пулю
func NewEnemyShot() *EnemyBullet {
	actor := &Actor{}
	b := &EnemyBullet{Actor: *actor}
	b.isActive = false
	return b
}

// IsActive возвращает состояние пули
func (b *EnemyBullet) IsActive() bool {
	return b.isActive
}

// SetInactive делает пулю неактивной
func (b *EnemyBullet) SetInactive() {
	b.isActive = false
}

// Init инициализация пули
func (b *EnemyBullet) Init(degree int, speed float64, x, y, size int) {
	b.speed = speed
	b.deg = degree

	b.vx = math.Cos(degToRad(degree)) * speed
	b.vy = math.Sin(degToRad(degree)) * speed

	b.spriteIndex = degreeToDirectionIndex(degree)

	b.width = size
	b.height = size
	b.x = float64(x)
	b.y = float64(y)

	b.isActive = true
}

// Draw отрисовка пули
func (b *EnemyBullet) Draw(screen *ebiten.Image) {
	sprite.EnemyBullet.SetPosition(b.x, b.y)
	sprite.EnemyBullet.SetIndex(b.spriteIndex)
	sprite.EnemyBullet.Draw(screen)
}

// Move движение пули
func (b *EnemyBullet) Move() {
	b.x = b.x + b.vx
	b.y = b.y + b.vy
	if b.isOutOfBoundary() {
		b.isActive = false
	}
}
