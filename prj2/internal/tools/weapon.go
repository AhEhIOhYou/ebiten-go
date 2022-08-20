package tools

import (
	"time"
)

type Weapon interface {
	Fire(x, y float64, degree int)
}

type shotFactoryFunction func(x, y float64, degree int)

type baseWeapon struct {
	shotFactory  shotFactoryFunction
	lastShotTime time.Time
}
