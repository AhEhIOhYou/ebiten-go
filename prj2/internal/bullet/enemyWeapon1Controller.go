package bullet

import "github.com/AhEhIOhYou/project2/prj2/internal/sprite"

type enemyWeapon1Controller struct {
	baseController
}

func (c *enemyWeapon1Controller) init(b *Bullet) {
	b.spr = sprite.EnemyBullet
	b.setSize(8, 8)
	b.setSpeed(b.speed, b.degree)
}
