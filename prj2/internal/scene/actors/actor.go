package actors

import "math"

type Actor struct {
	x      float64
	y      float64
	speed  float64
	vx     float64
	vy     float64
	width  int
	height int
	deg    int
}

// setSpeed устанавливает скорость актёра
func (a *Actor) setSpeed(speed float64) {
	a.speed = speed
}

// setSpeed устанавливает позицию актёра
func (a *Actor) setPosition(x, y float64) {
	a.x = x
	a.y = y
}

// GetPosition возвращает позицию актёра
func (a *Actor) GetPosition() (float64, float64) {
	return a.x, a.y
}

// GetDeg возвращает угол актёра
func (a *Actor) GetDeg() int {
	return a.deg
}

// GetNormalizedDegree возвращает нормализированный угол актера
func (a *Actor) GetNormalizedDegree() int {
	adjust := 22.5
	return int((float64(((a.deg+360)%360))+adjust)/45) * 45
}

// GetX возвращает позицию X
func (a *Actor) GetX() int {
	return int(a.x)
}

// GetY возвращает позицию Y
func (a *Actor) GetY() int {
	return int(a.y)
}

// GetWidth возвращает ширину
func (a *Actor) GetWidth() int {
	return a.width
}

// GetHeight возвращает высоту
func (a *Actor) GetHeight() int {
	return a.height
}

// degreeToDirectionIndex индекс отношения угла к направления
func degreeToDirectionIndex(degree int) int {
	adjust := 22.5
	return int(float64(degree)+90.0+360.0+adjust) % 360 / 45
}

//radToDeg перевод радиан в градусы
func radToDeg(radian float64) int {
	return int(radian * 180 / math.Pi)
}

//degToRad перевод градусов в радианы
func degToRad(degree int) float64 {
	return float64(degree) * math.Pi / 180
}

// Bounds представляют собой ограниченное поле, внутри которого актёр может перемещаться
type Bounds interface {
	GetLeft() int
	GetTop() int
	GetRight() int
	GetBottom() int
}

var (
	bounds Bounds
)

// SetBoundary устанавливает границы поля
func SetBoundary(b Bounds) {
	bounds = b
}

// isOutOfBoundary проверяет вышел ли актёр за границы своего поля
func (a *Actor) isOutOfBoundary() bool {
	if int(a.x)+a.width/2 < bounds.GetLeft() {
		return true
	}
	if int(a.x)-a.width > bounds.GetRight() {
		return true
	}
	if int(a.y)+a.height < bounds.GetTop() {
		return true
	}
	if int(a.y)-a.height/2 > bounds.GetBottom() {
		return true
	}
	return false
}

// Collider стуктура столкновений
type Collider interface {
	GetX() int
	GetY() int
	GetWidth() int
	GetHeight() int
}

// IsCollideWith проверка столкновения с другим коллайдером
func IsCollideWith(c1 Collider, c2 Collider) bool {
	return c1.GetX() <= c2.GetX()+c2.GetWidth() &&
		c2.GetX() <= c1.GetX()+c1.GetWidth() &&
		c1.GetY() <= c2.GetY()+c2.GetHeight() &&
		c2.GetY() <= c1.GetY()+c1.GetHeight()
}
