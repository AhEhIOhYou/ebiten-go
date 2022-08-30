package tools

import (
	"time"
)

// Normal представляет оружие игрока
type Normal struct{ baseWeapon }

func NewNormal(factory shotFactoryFunction) *Normal {
	w := &Normal{baseWeapon{}}
	w.shotFactory = factory

	return w
}

// Fire создает выстрелы
func (w *Normal) Fire(x, y, speed float64, degree int) {
	if time.Since(w.lastShotTime).Milliseconds() < w.cooldown {
		return
	}
	w.lastShotTime = time.Now()
	w.shotFactory(x, y, speed, degree)
}
