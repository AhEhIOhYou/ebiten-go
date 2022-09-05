package weapon

import (
	"time"
)

type Weapon interface {
	Fire(x, y, speed float64, degree int, angles []int, pos [][]float64)
	SetDegree(degree int)
	GetDegree() int
	SetCooldown(cooldown int64)
	GetCooldown() int64
	SetBulletSpeed(speed float64)
	GetBulletSpeed() float64
	SetAdjustAngles(angles []int)
	GetAdjustAngles() []int
	SetAdjustPos(positions [][]float64)
	GetAdjustPos() [][]float64
}

type shotFactoryFunction func(x, y, speed float64, degree int, angles []int, pos [][]float64)

type baseWeapon struct {
	shotFactory           shotFactoryFunction
	lastShotTime          time.Time
	degree                int
	cooldown              int64
	bulletSpeed           float64
	adjustBulletAngles    []int
	adjustBulletPositions [][]float64
}

// PlayerWeapon представляет оружие игрока
type PlayerWeapon struct{ baseWeapon }

type EnemyWeapon struct{ baseWeapon }

func NewPlayerWeapon(factory shotFactoryFunction) *PlayerWeapon {
	w := &PlayerWeapon{baseWeapon{}}
	w.shotFactory = factory
	return w
}

func NewEnemyWeapon(factory shotFactoryFunction) *EnemyWeapon {
	w := &EnemyWeapon{baseWeapon{}}
	w.shotFactory = factory
	return w
}

// Fire создает выстрелы
func (w *PlayerWeapon) Fire(x, y, speed float64, degree int, angles []int, pos [][]float64) {
	if time.Since(w.lastShotTime).Milliseconds() < w.cooldown {
		return
	}
	w.lastShotTime = time.Now()
	w.shotFactory(x, y, speed, degree, angles, pos)
}

func (w *EnemyWeapon) Fire(x, y, speed float64, degree int, angles []int, pos [][]float64) {
	if time.Since(w.lastShotTime).Milliseconds() < w.cooldown {
		return
	}
	w.lastShotTime = time.Now()
	w.shotFactory(x, y, speed, degree, angles, pos)
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

func (w *baseWeapon) SetAdjustPos(positions [][]float64) {
	w.adjustBulletPositions = positions
}

func (w *baseWeapon) GetAdjustPos() [][]float64 {
	return w.adjustBulletPositions
}
