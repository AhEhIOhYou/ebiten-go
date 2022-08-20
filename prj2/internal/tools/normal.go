package tools

import (
	"github.com/AhEhIOhYou/project2/prj2/internal/bullet"
	"github.com/AhEhIOhYou/project2/prj2/internal/objectpool"
	"time"
)

// Normal represents player's weapon
type Normal struct {
	lastShotTime time.Time
	bulletKind   bullet.Kind
}

// NewNormal creates new struct
func NewNormal(bulletKind bullet.Kind) *Normal {
	w := &Normal{}
	w.bulletKind = bulletKind

	return w
}

// Fire create shots
func (w *Normal) Fire(shooter Shooter, bullets *objectpool.Pool) {
	if time.Since(w.lastShotTime).Milliseconds() < 20 {
		return
	}
	w.lastShotTime = time.Now()

	b := (*bullet.Bullet)(bullets.CreateFromPool())
	if b == nil {
		return
	}
	b.Init(w.bulletKind, shooter.GetDegree())
	b.SetPosition(shooter.GetX(), shooter.GetY())
}
