package bullet

import (
	"github.com/AhEhIOhYou/project2/prj2/internal/fields"
	"github.com/AhEhIOhYou/project2/prj2/internal/sprite"
	"github.com/AhEhIOhYou/project2/prj2/internal/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

// Bullet представление пули
type Bullet struct {
	x, y          float64
	width, height float64
	currField     *fields.Field
	isActive      bool
	speed         float64
	vx            float64
	vy            float64
	degree        int
	spr           *sprite.Sprite
	sprIndex      int
}

// Kind represetns the kind of the shot
type Kind int

const (
	KindPlayerNormal Kind = iota
	KindEnemyNormal
)

// NewBullet возвращает инициализированную пулю
func NewBullet(f *fields.Field) *Bullet {
	bullet := &Bullet{}
	bullet.currField = f
	return bullet
}

func (b *Bullet) Init(kind Kind, degree int) {
	b.isActive = true

	switch kind {
	case KindPlayerNormal:
		b.spr = sprite.PlayerBullet
		b.setSize(4, 4)
		b.setSpeed(20, degree)
		break
	case KindEnemyNormal:
		b.spr = sprite.PlayerBullet
		b.setSize(10, 10)
		b.setSpeed(1.44, degree)
	}
}

// IsActive returns if this is active
func (b *Bullet) IsActive() bool {
	return b.isActive
}

// GetX returns x
func (b *Bullet) GetX() float64 {
	return b.x
}

// GetY returns y
func (b *Bullet) GetY() float64 {
	return b.y
}

// GetWidth returns width
func (b *Bullet) GetWidth() float64 {
	return b.width
}

// GetHeight returns height
func (b *Bullet) GetHeight() float64 {
	return b.height
}

// GetDegree returns the degree
func (b *Bullet) GetDegree() int {
	return b.degree
}

// SetPosition sets the position
func (b *Bullet) SetPosition(x, y float64) {
	b.x = x
	b.y = y
}

// Draw draws this
func (b *Bullet) Draw(screen *ebiten.Image) {
	spr := b.spr
	spr.SetPosition(b.x, b.y)
	spr.SetIndex(b.sprIndex)
	spr.Draw(screen)
}

// Move moves this
func (b *Bullet) Move() {
	b.SetPosition(b.x+b.vx, b.y+b.vy)
	if utils.IsOutOfArea(b, b.currField) {
		b.isActive = false
	}
}

// SetField returns field
func (b *Bullet) SetField(f *fields.Field) {
	b.currField = f
}

// OnHit should be called on hit something
func (b *Bullet) OnHit() {
	b.isActive = false
}

func (b *Bullet) setSize(width, height float64) {
	b.width = width
	b.height = height
}

func (b *Bullet) setSpeed(speed float64, degree int) {
	b.speed = speed
	b.degree = degree
	b.vx = math.Cos(utils.DegToRad(b.degree)) * speed
	b.vy = math.Sin(utils.DegToRad(b.degree)) * speed
}
