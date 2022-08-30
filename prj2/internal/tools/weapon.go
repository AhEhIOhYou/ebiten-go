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
