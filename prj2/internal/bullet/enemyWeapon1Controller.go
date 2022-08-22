package bullet

import "github.com/AhEhIOhYou/project2/prj2/internal/sprite"

type enemyWeapon1Controller struct {
	baseController
}

func (c *enemyWeapon1Controller) init(b *Bullet) {
	b.spr = sprite.EnemyBullet
	b.setSize(2, 2)
	b.setSpeed(0.5, b.degree)
}
