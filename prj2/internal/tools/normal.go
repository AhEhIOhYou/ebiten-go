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
	if time.Since(w.lastShotTime).Milliseconds() < 200 {
		return
	}
	w.lastShotTime = time.Now()

	b := (*bullet.Bullet)(bullets.CreateFromPool())
	if b == nil {
		return
	}
	b.Init(w.bulletKind, shooter.GetDegree())
	b.SetPosition(shooter.GetX(), shooter.GetY())

	b2 := (*bullet.Bullet)(bullets.CreateFromPool())
	if b2 == nil {
		return
	}
	b2.Init(w.bulletKind, shooter.GetDegree())
	b2.SetPosition(shooter.GetX()+20, shooter.GetY())

	b3 := (*bullet.Bullet)(bullets.CreateFromPool())
	if b3 == nil {
		return
	}
	b3.Init(w.bulletKind, shooter.GetDegree())
	b3.SetPosition(shooter.GetX()-20, shooter.GetY())
}
