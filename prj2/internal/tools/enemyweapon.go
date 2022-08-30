package tools

import "time"

type EnemyWeapon1 struct{ baseWeapon }

func NewEnemyWeapon1(factory shotFactoryFunction) *EnemyWeapon1 {
	w := &EnemyWeapon1{baseWeapon{}}
	w.shotFactory = factory
	return w
}

func (w *EnemyWeapon1) Fire(x, y, speed float64, degree int, angles []int) {
	if time.Since(w.lastShotTime).Milliseconds() < w.cooldown {
		return
	}
	w.lastShotTime = time.Now()
	w.shotFactory(x, y, speed, degree, angles)
}
