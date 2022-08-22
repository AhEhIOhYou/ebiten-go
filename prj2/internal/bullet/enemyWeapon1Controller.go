package bullet

import "github.com/AhEhIOhYou/project2/prj2/internal/sprite"

type enemyWeapon1Controller struct {
	baseController
}

func (c *enemyWeapon1Controller) init(b *Bullet) {
	b.spr = sprite.PlayerBullet
	b.setSize(2, 2)
	b.setSpeed(20, b.degree)
}
