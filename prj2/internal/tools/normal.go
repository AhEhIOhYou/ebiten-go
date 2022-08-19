package tools

import (
	"github.com/AhEhIOhYou/project2/prj2/internal/bullet"
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
func (w *Normal) Fire(shooter Shooter, shots []*bullet.Bullet) {
	if time.Since(w.lastShotTime).Milliseconds() < 20 {
		return
	}
	w.lastShotTime = time.Now()

	for i := 0; i < len(shots); i++ {
		s := shots[i]
		if s.IsActive() {
			continue
		}
		s.Init(w.bulletKind, shooter.GetDegree())
		s.SetPosition(shooter.GetX(), shooter.GetY())
		break
	}

}
