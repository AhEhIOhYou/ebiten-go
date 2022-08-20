package bullet

import (
	"github.com/AhEhIOhYou/project2/prj2/internal/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

// controller интерфейс контроллера пуль
type controller interface {
	init(b *Bullet)
	update(b *Bullet)
	draw(b *Bullet, screen *ebiten.Image)
}

type baseController struct{}

func (c *baseController) init(b *Bullet) {}

func (c *baseController) update(b *Bullet) {
	b.setPosition(b.x+b.vx, b.y+b.vy)
	if utils.IsOutOfArea(b, b.field) {
		b.isActive = false
	}
}

func (c *baseController) draw(b *Bullet, screen *ebiten.Image) {
	spr := b.spr
	spr.SetPosition(b.x, b.y)
	spr.SetIndex(b.sprIndex)
	spr.Draw(screen)
}
