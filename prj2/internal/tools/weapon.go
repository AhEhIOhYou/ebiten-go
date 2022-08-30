package tools

import (
	"time"
)

type Weapon interface {
	Fire(x, y, speed float64, degree int)
	SetDegree(degree int)
	GetDegree() int
	SetCooldown(cooldown int64)
	GetCooldown() int64
	SetBulletSpeed(speed float64)
	GetBulletSpeed() float64
}

type shotFactoryFunction func(x, y, speed float64, degree int)

type baseWeapon struct {
	shotFactory  shotFactoryFunction
	lastShotTime time.Time
	degree       int
	cooldown     int64
	bulletSpeed  float64
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
