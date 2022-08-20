package bullet

import (
	"github.com/AhEhIOhYou/project2/prj2/internal/sprite"
)

// playerNormalController контроллер для инициализации пуль игрока
type playerNormalController struct {
	baseController
}

func (c *playerNormalController) init(b *Bullet) {
	b.spr = sprite.PlayerBullet
	b.setSize(4, 4)
	b.setSpeed(30, b.degree)
}
