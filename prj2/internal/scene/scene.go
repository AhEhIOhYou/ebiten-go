package scene

import (
	"github.com/AhEhIOhYou/project2/prj2/internal/actors"
	"github.com/AhEhIOhYou/project2/prj2/internal/bullet"
	"github.com/AhEhIOhYou/project2/prj2/internal/fields"
	"github.com/AhEhIOhYou/project2/prj2/internal/inputs"
	"github.com/AhEhIOhYou/project2/prj2/internal/objectpool"
	"github.com/AhEhIOhYou/project2/prj2/internal/tools"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"time"
	"unsafe"
)

const (
	maxPlayerShot = 200
)

var (
	input        *inputs.Input
	currentField *fields.Field

	player        *actors.Player
	playerBullets *objectpool.Pool
)

// Scene представляет сцену
type Scene struct {
	time         time.Time
	screenWidth  int
	screenHeight int
}

// NewScene вернет стандартную сцену
func NewScene(screenWidth, screenHeight int) *Scene {
	stg := &Scene{}
	stg.screenWidth = screenWidth
	stg.screenHeight = screenHeight
	stg.time = time.Now()
	stg.initGame()
	return stg
}

// initGame инициализирует игру
func (stg *Scene) initGame() {
	input = inputs.New()
	currentField = fields.NewField()

	player = actors.NewPlayer(currentField)
	player.Init()
	player.SetMainWeapon(tools.NewNormal(bullet.KindPlayerNormal))

	playerBullets = objectpool.NewPool()
	for i := 0; i < maxPlayerShot; i++ {
		playerBullets.AddToPool(unsafe.Pointer(bullet.NewBullet(currentField)))
	}
}

// Update обновляет состояние сцены (актеров и окружения)
func (stg *Scene) Update() {
	input.Update()
	checkCollision()
	player.Action(input.Horizontal, input.Vertical, input.Fire, input.Focus)

	player.Action(input.Horizontal, input.Vertical, input.Fire, input.Focus)
	if input.Fire {
		player.FireWeapon(playerBullets)
	}

	for ite := playerBullets.GetIterator(); ite.HasNext(); {
		obj := ite.Next()
		p := (*bullet.Bullet)(obj.GetData())
		if p.IsActive() == false {
			obj.SetInactive()
			continue
		}
		p.Move()
	}
	playerBullets.Sweep()
}

// Draw отрисовывает всех действиующих лиц сцены
func (stg *Scene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x10, 0x10, 0x30, 0xff})
	currentField.Draw(screen)
	player.Draw(screen)

	for ite := playerBullets.GetIterator(); ite.HasNext(); {
		p := (*bullet.Bullet)(ite.Next().GetData())
		if p.IsActive() == false {
			continue
		}
		p.Draw(screen)
	}
}

// checkCollision Проверка столкновений
func checkCollision() {
	// В разрабокте
}
