package tools

import (
	"github.com/AhEhIOhYou/project2/prj2/internal/objectpool"
)

// Shooter represents shooter
type Shooter interface {
	GetX() float64
	GetY() float64
	GetDegree() int
}

// Weapon represents weapon
type Weapon interface {
	Fire(shooter Shooter, shots *objectpool.Pool)
}
