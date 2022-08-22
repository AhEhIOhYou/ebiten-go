package tools

import "time"

type EnemyWeapon1 struct{ baseWeapon }

func NewEnemyWeapon1(factory shotFactoryFunction) *EnemyWeapon1 {
	w := &EnemyWeapon1{baseWeapon{}}
	w.shotFactory = factory

	return w
}

func (w *EnemyWeapon1) Fire(x, y float64, degree int) {
	if time.Since(w.lastShotTime).Milliseconds() < 36 {
		return
	}
	w.lastShotTime = time.Now()
	w.shotFactory(x, y, degree)
}
