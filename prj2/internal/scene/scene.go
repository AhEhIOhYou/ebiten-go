package scene

import (
	"github.com/AhEhIOhYou/project2/prj2/internal/actors"
	"github.com/AhEhIOhYou/project2/prj2/internal/bullet"
	"github.com/AhEhIOhYou/project2/prj2/internal/fields"
	"github.com/AhEhIOhYou/project2/prj2/internal/inputs"
	"github.com/AhEhIOhYou/project2/prj2/internal/shared"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"time"
	"unsafe"
)

const (
	maxPlayerShot = 500
)

// Scene представляет сцену
type Scene struct {
	input  *inputs.Input
	field  *fields.Field
	player *actors.Player

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
	stg.setupGame()
	return stg
}

// initGame инициализирует игру
func (stg *Scene) initGame() {
	field := fields.NewField()
	stg.input = inputs.New()
	stg.field = field
	stg.player = actors.NewPlayer(field, shared.PlayerBullets)

	for i := 0; i < maxPlayerShot; i++ {
		shared.PlayerBullets.AddToPool(unsafe.Pointer(bullet.NewBullet(field)))
	}
}

func (stg *Scene) setupGame() {
	shared.PlayerBullets.Clean()
	stg.player.Init()
}

// Update обновляет состояние сцены (актеров и окружения)
func (stg *Scene) Update() {
	input := stg.input

	input.Update()
	if input.Reload {
		stg.setupGame()
	}

	stg.player.Action(input.Horizontal, input.Vertical, input.Fire, input.Focus)
	if input.Fire {
		stg.player.FireWeapon()
	}

	for ite := shared.PlayerBullets.GetIterator(); ite.HasNext(); {
		obj := ite.Next()
		p := (*bullet.Bullet)(obj.GetData())
		if p.IsActive() == false {
			obj.SetInactive()
			continue
		}
		p.Update()
	}
	shared.PlayerBullets.Sweep()
}

// Draw отрисовывает всех действиующих лиц сцены
func (stg *Scene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x10, 0x10, 0x30, 0xff})
	stg.field.Draw(screen)
	stg.player.Draw(screen)

	for ite := shared.PlayerBullets.GetIterator(); ite.HasNext(); {
		p := (*bullet.Bullet)(ite.Next().GetData())
		p.Draw(screen)
	}
}

// checkCollision Проверка столкновений
func checkCollision() {
	// В разрабокте
}
