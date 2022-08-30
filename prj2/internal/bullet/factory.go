package bullet

import "github.com/AhEhIOhYou/project2/prj2/internal/shared"

var (
	playerNormal = &playerNormalController{baseController{}}
	enemyWeapon1 = &enemyWeapon1Controller{baseController{}}
)

// NormalPlayerShot Создает выстрел
func NormalPlayerShot(x, y, speed float64, degree int, angles []int) {
	b1 := (*Bullet)(shared.PlayerBullets.CreateFromPool())
	if b1 == nil {
		return
	}
	b1.init(playerNormal, x, y, speed, degree)

	b2 := (*Bullet)(shared.PlayerBullets.CreateFromPool())
	if b2 == nil {
		return
	}
	b2.init(playerNormal, x+5, y-10, speed, degree+5)

	b3 := (*Bullet)(shared.PlayerBullets.CreateFromPool())
	if b3 == nil {
		return
	}
	b3.init(playerNormal, x-5, y-10, speed, degree-5)
}

func EnemyWeaponShot(x, y, speed float64, weaponDegree int, angles []int) {

	for i := 0; i < len(angles); i++ {
		b := (*Bullet)(shared.EnemyBullets.CreateFromPool())
		if b == nil {
			return
		}
		b.init(enemyWeapon1, x, y, speed, weaponDegree+angles[i])
	}

}
