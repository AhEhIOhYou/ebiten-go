package tools

import (
	"time"
)

type Weapon interface {
	Fire(x, y, speed float64, degree int, angles []int)
	SetDegree(degree int)
	GetDegree() int
	SetCooldown(cooldown int64)
	GetCooldown() int64
	SetBulletSpeed(speed float64)
	GetBulletSpeed() float64
	SetAdjustAngles(angles []int)
	GetAdjustAngles() []int
}

type shotFactoryFunction func(x, y, speed float64, degree int, angles []int)

type baseWeapon struct {
	shotFactory        shotFactoryFunction
	lastShotTime       time.Time
	degree             int
	cooldown           int64
	bulletSpeed        float64
	adjustBulletAngles []int
}

// Normal представляет оружие игрока
type Normal struct{ baseWeapon }

func NewNormal(factory shotFactoryFunction) *Normal {
	w := &Normal{baseWeapon{}}
	w.shotFactory = factory

	return w
}

// Fire создает выстрелы
func (w *Normal) Fire(x, y, speed float64, degree int, angles []int) {
	if time.Since(w.lastShotTime).Milliseconds() < w.cooldown {
		return
	}
	w.lastShotTime = time.Now()
	w.shotFactory(x, y, speed, degree, angles)
}

type wp struct{ baseWeapon }

func NewEnemyWeapon1(factory shotFactoryFunction) *wp {
	w := &wp{baseWeapon{}}
	w.shotFactory = factory
	return w
}

func (w *wp) Fire(x, y, speed float64, degree int, angles []int) {
	if time.Since(w.lastShotTime).Milliseconds() < w.cooldown {
		return
	}
	w.lastShotTime = time.Now()
	w.shotFactory(x, y, speed, degree, angles)
}

func (w *baseWeapon) SetBulletSpeed(speed float64) {
	w.bulletSpeed = speed
}

func (w *baseWeapon) GetBulletSpeed() float64 {
	return w.bulletSpeed
}

func (w *baseWeapon) SetDegree(degree int) {
	w.degree = degree
}

func (w *baseWeapon) GetDegree() int {
	return w.degree
}

func (w *baseWeapon) SetCooldown(cooldown int64) {
	w.cooldown = cooldown
}

func (w *baseWeapon) GetCooldown() int64 {
	return w.cooldown
}

func (w *baseWeapon) SetAdjustAngles(angles []int) {
	w.adjustBulletAngles = angles
}

func (w *baseWeapon) GetAdjustAngles() []int {
	return w.adjustBulletAngles
}
