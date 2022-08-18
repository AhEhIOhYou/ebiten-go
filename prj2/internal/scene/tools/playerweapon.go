package tools

import (
	"github.com/AhEhIOhYou/project2/prj2/internal/scene/actors"
	"time"
)

const (
	weaponReloadTimeMs = 10
	weaponSpeed        = 20
	weaponSize         = 4
)

// PlayerWeapon представляет оружие игрока
type PlayerWeapon struct {
	lastShotTime time.Time
}

// Shot создает выстрел пулей
func (w *PlayerWeapon) Shot(x, y float64, degree int, playerShots []*actors.PlayerBullet) {
	if time.Since(w.lastShotTime).Milliseconds() < weaponReloadTimeMs {
		return
	}
	w.lastShotTime = time.Now()

	for i := 0; i < len(playerShots); i += 3 {
		s := playerShots[i]
		if s.IsActive() {
			continue
		}
		s.Init(degree, weaponSpeed, int(x), int(y), weaponSize)
		s = playerShots[i+1]
		if s.IsActive() {
			continue
		}
		s.Init(degree, weaponSpeed, int(x+20), int(y), weaponSize)
		s = playerShots[i+2]
		if s.IsActive() {
			continue
		}
		s.Init(degree, weaponSpeed, int(x-20), int(y), weaponSize)
		break
	}
}
