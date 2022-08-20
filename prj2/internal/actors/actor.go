package actors

import (
	"github.com/AhEhIOhYou/project2/prj2/internal/fields"
	"github.com/AhEhIOhYou/project2/prj2/internal/objectpool"
	"github.com/AhEhIOhYou/project2/prj2/internal/sprite"
	"github.com/AhEhIOhYou/project2/prj2/internal/tools"
	"github.com/AhEhIOhYou/project2/prj2/internal/utils"
	"math"
)

type Actor struct {
	x, y          float64
	width, height float64
	currField     *fields.Field
	isActive      bool
	speed         float64
	vx            float64
	vy            float64
	degree        int
	spr           *sprite.Sprite
	sprIndex      int
	life          int
	mainWeapon    tools.Weapon
	movdweTo      struct{ x, y float64 }
	bulletPool    *objectpool.Pool
}

func NewActor() *Actor {
	actor := &Actor{}
	return actor
}

// GetX возвращает позицию X
func (a *Actor) GetX() float64 {
	return a.x
}

// GetY возвращает позицию Y
func (a *Actor) GetY() float64 {
	return a.y
}

// GetPosition возвращает позицию актёра
func (a *Actor) GetPosition() (float64, float64) {
	return a.x, a.y
}

// GetWidth возвращает ширину
func (a *Actor) GetWidth() float64 {
	return a.width
}

// GetHeight возвращает высоту
func (a *Actor) GetHeight() float64 {
	return a.height
}

// GetDegree возвращает угол актёра
func (a *Actor) GetDegree() int {
	return a.degree
}

// IsActive возвращает угол актёра
func (a *Actor) IsActive() bool {
	return a.isActive
}

// GetMainSpriteIndex returns sprite
func (a *Actor) GetMainSpriteIndex() int {
	return a.sprIndex
}

// SetMainSpriteIndex returns sprite
func (a *Actor) SetMainSpriteIndex(index int) {
	a.sprIndex = index
}

// SetSpeed устанавливает скорость актёра
func (a *Actor) SetSpeed(speed float64, degree int) {
	a.speed = speed
	a.degree = degree
	a.vx = math.Cos(utils.DegToRad(a.degree)) * speed
	a.vy = math.Sin(utils.DegToRad(a.degree)) * speed
}

// SetMainWeapon устанавливает пушку
func (a *Actor) SetMainWeapon(mainWeapon tools.Weapon) {
	a.mainWeapon = mainWeapon
}

// SetPosition устанавливает позицию актёра
func (a *Actor) SetPosition(x, y float64) {
	a.x = x
	a.y = y
}

// FireWeapon ведет огонь из оружия
func (a *Actor) FireWeapon() {
	a.mainWeapon.Fire(a.x, a.y, a.degree)
}

// SetField возвращает поле игрока
func (a *Actor) SetField(f *fields.Field) {
	a.currField = f
}

func (a *Actor) setSize(width, height float64) {
	a.width = width
	a.height = height
}
