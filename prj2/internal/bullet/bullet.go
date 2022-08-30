package bullet

import (
	"github.com/AhEhIOhYou/project2/prj2/internal/fields"
	"github.com/AhEhIOhYou/project2/prj2/internal/sprite"
	"github.com/AhEhIOhYou/project2/prj2/internal/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"math"
)

// Bullet представляет пулю
type Bullet struct {
	controller    controller
	x, y          float64
	width, height float64
	field         *fields.Field
	isActive      bool
	speed         float64
	vx            float64
	vy            float64
	degree        int
	spr           *sprite.Sprite
	sprIndex      int
}

func NewBullet(f *fields.Field) *Bullet {
	s := &Bullet{}
	s.field = f

	return s
}

func (b *Bullet) IsActive() bool {
	return b.isActive
}

func (b *Bullet) GetX() float64 {
	return b.x
}

func (b *Bullet) GetY() float64 {
	return b.y
}

func (b *Bullet) GetWidth() float64 {
	return b.width
}

func (b *Bullet) GetHeight() float64 {
	return b.height
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	b.controller.draw(b, screen)
}

// Update Во время обновления пули вызывается контроллер пуль
func (b *Bullet) Update() {
	b.controller.update(b)
}

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

func (b *Bullet) init(controller controller, x, y, speed float64, degree int) {
	b.isActive = true
	b.x = x
	b.y = y
	b.degree = degree
	b.controller = controller
	b.speed = speed
	controller.init(b)
}

func (b *Bullet) setPosition(x, y float64) {
	b.x = x
	b.y = y
}
