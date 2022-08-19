package scene

import (
	"github.com/AhEhIOhYou/project2/prj2/internal/actors"
	"github.com/AhEhIOhYou/project2/prj2/internal/bullet"
	"github.com/AhEhIOhYou/project2/prj2/internal/fields"
	"github.com/AhEhIOhYou/project2/prj2/internal/inputs"
	"github.com/AhEhIOhYou/project2/prj2/internal/tools"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"time"
)

const (
	maxPlayerShot = 200
)

var (
	input        *inputs.Input
	currentField *fields.Field
	player       *actors.Player
	playerShots  [maxPlayerShot]*bullet.Bullet
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

	player = actors.NewPlayer()
	player.Init()
	player.SetMainWeapon(tools.NewNormal(bullet.KindPlayerNormal))
	player.SetField(currentField)

	for i := 0; i < len(playerShots); i++ {
		playerShots[i] = bullet.NewBullet()
		playerShots[i].SetField(currentField)
	}
}

// Update обновляет состояние сцены (актеров и окружения)
func (stg *Scene) Update() {
	input.Update()
	checkCollision()
	player.Action(input.Horizontal, input.Vertical, input.Fire, input.Focus)

	player.Action(input.Horizontal, input.Vertical, input.Fire, input.Focus)
	if input.Fire {
		player.FireWeapon(playerShots[:])
	}

	for i := 0; i < len(playerShots); i++ {
		p := playerShots[i]
		if p.IsActive() == false {
			continue
		}
		p.Move()
	}
}

// Draw отрисовывает всех действиующих лиц сцены
func (stg *Scene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x10, 0x10, 0x30, 0xff})
	currentField.Draw(screen)
	player.Draw(screen)

	for i := 0; i < len(playerShots); i++ {
		p := playerShots[i]
		if p.IsActive() == false {
			continue
		}
		p.Draw(screen)
	}
}

// checkCollision Проверка столкновений
func checkCollision() {
	for i := 0; i < len(playerShots); i++ {
		p := playerShots[i]
		if p.IsActive() == false {
			continue
		}
	}
}
