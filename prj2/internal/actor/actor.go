package actor

import "math"

var (
	directionDegreeMap = map[int]int{
		1: 0,
		2: 45,
		3: 90,
		4: 135,
		5: 180,
		6: 225,
		7: 270,
		8: 315,
	}
)

// Actor represents actor
type Actor struct {
	X               float64
	Y               float64
	Speed           float64
	NSpd            float64
	Width           int
	Height          int
	Direction       int
	DirectionDegree int
}

// SetSpeed sets the speed to the actor
func (a *Actor) SetSpeed(speed float64) {
	a.Speed = speed
	a.NSpd = math.Cos(math.Pi/4) * a.Speed
}

// SetPosition sets the position
func (a *Actor) SetPosition(x, y float64) {
	a.X = x
	a.Y = y
}

// SetDirection sets the direction
func (a *Actor) SetDirection(direction int) {
	a.Direction = direction
	a.DirectionDegree = directionDegreeMap[direction]
}
