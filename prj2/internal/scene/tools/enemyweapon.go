package tools

import (
	"github.com/AhEhIOhYou/project2/prj2/internal/scene/actors"
	"time"
)

const (
	lReloadTimeMs = 10
	wspeed        = 10
	wsize         = 4
)

// EnemyWeapon представляет оружие противника
type EnemyWeapon struct {
	lastShotTime time.Time
}

// Shot создает выстрел пулей
func (w *EnemyWeapon) Shot(x, y float64, degree int, enemyShots []*actors.EnemyBullet) {
	if time.Since(w.lastShotTime).Milliseconds() < lReloadTimeMs {
		return
	}
	w.lastShotTime = time.Now()

	for i := 0; i < len(enemyShots); i++ {
		s := enemyShots[i]
		if s.IsActive() {
			continue
		}
		s.Init(degree, wspeed, int(x), int(y), wsize)
	}
}
