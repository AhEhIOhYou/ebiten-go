package bullet

import (
	"github.com/AhEhIOhYou/project2/prj2/internal/shared"
	"github.com/AhEhIOhYou/project2/prj2/internal/sprite"
)

type playerNormalController struct {
	baseController
}
type enemyWeaponController struct {
	baseController
}

func (c *playerNormalController) init(b *Bullet) {
	b.spr = sprite.PlayerBullet
	b.setSize(4, 4)
	b.setSpeed(30, b.degree)
}

func (c *enemyWeaponController) init(b *Bullet) {
	b.spr = sprite.EnemyBullet
	b.setSize(8, 8)
	b.setSpeed(b.speed, b.degree)
}

var (
	playerNormal = &playerNormalController{baseController{}}
	enemyWeapon  = &enemyWeaponController{baseController{}}
)

// PlayerWeaponShot Создает выстрел
func PlayerWeaponShot(x, y, speed float64, weaponDegree int, angles []int, pos [][]float64) {

	for i := 0; i < len(angles); i++ {
		b := (*Bullet)(shared.PlayerBullets.CreateFromPool())
		if b == nil {
			return
		}
		b.init(playerNormal, x+pos[i][0], y+pos[i][1], speed, weaponDegree+angles[i])
	}

}

func EnemyWeaponShot(x, y, speed float64, weaponDegree int, angles []int, pos [][]float64) {

	for i := 0; i < len(angles); i++ {
		b := (*Bullet)(shared.EnemyBullets.CreateFromPool())
		if b == nil {
			return
		}
		b.init(enemyWeapon, x+pos[i][0], y+pos[i][1], speed, weaponDegree+angles[i])
	}

}
