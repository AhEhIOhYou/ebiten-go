package tools

import (
	"github.com/AhEhIOhYou/project2/prj2/internal/scene/actors"
	"time"
)

const (
	weapon1ReloadTimeMs = 30
	weapon1Speed        = 10
	weapon1Size         = 4
)

// PlayerWeapon представляет оружие игрока
type PlayerWeapon struct {
	lastShotTime time.Time
}

// Shot создает выстрел пулей
func (w *PlayerWeapon) Shot(x, y float64, degree int, playerShots []*actors.PlayerBullet) {
	if time.Since(w.lastShotTime).Milliseconds() < weapon1ReloadTimeMs {
		return
	}
	w.lastShotTime = time.Now()

	for i := 0; i < len(playerShots); i++ {
		s := playerShots[i]
		if s.IsActive() {
			continue
		}
		s.Init(degree, weapon1Speed, int(x), int(y), weapon1Size)
		break
	}
}
