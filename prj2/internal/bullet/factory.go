package bullet

import "github.com/AhEhIOhYou/project2/prj2/internal/shared"

var (
	playerNormal = &playerNormalController{baseController{}}
	enemyWeapon1 = &enemyWeapon1Controller{baseController{}}
)

// NormalPlayerShot Создает выстрел
func NormalPlayerShot(x, y float64, degree int) {
	b1 := (*Bullet)(shared.PlayerBullets.CreateFromPool())
	if b1 == nil {
		return
	}
	b1.init(playerNormal, x, y, degree)

	b2 := (*Bullet)(shared.PlayerBullets.CreateFromPool())
	if b2 == nil {
		return
	}
	b2.init(playerNormal, x+5, y-10, degree+5)

	b3 := (*Bullet)(shared.PlayerBullets.CreateFromPool())
	if b3 == nil {
		return
	}
	b3.init(playerNormal, x-5, y-10, degree-5)
}

func EnemyWeapon1Shot(x, y float64, degree int) {

	var data []float64

	for i := 0.0; i < 360; i += 22.5 {
		data = append(data, i)
	}
	//var data = []int{0, 45, 90, 135, 180, 225, 270, 305}

	//var data = []int{45, 135, 220, 305}

	for i := 0; i < len(data); i++ {
		b := (*Bullet)(shared.EnemyBullets.CreateFromPool())
		if b == nil {
			return
		}
		b.init(enemyWeapon1, x, y, degree+int(data[i]))
	}
}
