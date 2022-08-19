package tools

import (
	"github.com/AhEhIOhYou/project2/prj2/internal/bullet"
)

// Shooter represents shooter
type Shooter interface {
	GetX() float64
	GetY() float64
	GetDegree() int
}

// Weapon represents weapon
type Weapon interface {
	Fire(shooter Shooter, shots []*bullet.Bullet)
}
