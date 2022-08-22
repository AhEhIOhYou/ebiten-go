package scene

import (
	"github.com/AhEhIOhYou/project2/prj2/internal/actors"
	"github.com/AhEhIOhYou/project2/prj2/internal/bullet"
	"github.com/AhEhIOhYou/project2/prj2/internal/fields"
	"github.com/AhEhIOhYou/project2/prj2/internal/inputs"
	"github.com/AhEhIOhYou/project2/prj2/internal/shared"
	"github.com/AhEhIOhYou/project2/prj2/internal/ui"
	"github.com/AhEhIOhYou/project2/prj2/internal/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"time"
	"unsafe"
)

type gameState int

const (
	maxPlayerShot              = 500
	maxEnemyShot               = 5000
	gameStateLoading gameState = iota
	gameStatePlaying
)

// Scene представляет сцену
type Scene struct {
	state gameState

	input      *inputs.Input
	field      *fields.Field
	time       time.Time
	viewCenter struct{ x, y float64 }

	player *actors.Player
	enemy  *actors.Enemy
}

// NewScene вернет стандартную сцену
func NewScene(screenWidth, screenHeight int) *Scene {
	stg := &Scene{}

	stg.state = gameStateLoading

	stg.viewCenter.x = float64(ui.GetScreenWidth() / 2)
	stg.viewCenter.y = float64(ui.GetScreenHeight() / 2)
	stg.time = time.Now()

	stg.initGame()
	stg.setupGame()

	stg.state = gameStatePlaying

	return stg
}

// initGame инициализирует игру
func (stg *Scene) initGame() {
	field := fields.NewField()
	stg.input = inputs.New()
	stg.field = field
	stg.player = actors.NewPlayer(field, shared.PlayerBullets)

	stg.enemy = actors.NewEnemy(field, shared.EnemyBullets)

	for i := 0; i < maxPlayerShot; i++ {
		shared.PlayerBullets.AddToPool(unsafe.Pointer(bullet.NewBullet(field)))
	}

	for i := 0; i < maxEnemyShot; i++ {
		shared.EnemyBullets.AddToPool(unsafe.Pointer(bullet.NewBullet(field)))
	}
}

func (stg *Scene) setupGame() {
	shared.PlayerBullets.Clean()
	shared.EnemyBullets.Clean()
	stg.player.Init()
	stg.enemy.Init(320, 200, 1)
}

// Update обновляет состояние сцены (актеров и окружения)
func (stg *Scene) Update() {
	input := stg.input

	stg.checkCollision()

	input.Update()
	if input.Reload {
		stg.setupGame()
	}

	if time.Since(stg.time).Seconds() > 1 {
		stg.enemy.SetDegree(stg.enemy.GetDegree() + 6)
		stg.enemy.FireWeapon()
	}

	stg.player.Action(input.Horizontal, input.Vertical, input.Fire, input.Focus)
	if input.Fire {
		stg.player.FireWeapon()
	}

	for ite := shared.EnemyBullets.GetIterator(); ite.HasNext(); {
		obj := ite.Next()
		p := (*bullet.Bullet)(obj.GetData())
		if p.IsActive() == false {
			obj.SetInactive()
			continue
		}
		p.Update()
	}
	shared.EnemyBullets.Sweep()

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
	screen.Fill(color.RGBA{R: 0x10, G: 0x10, B: 0x30, A: 0xff})
	stg.field.Draw(screen)
	stg.player.Draw(screen)

	stg.enemy.Draw(screen)

	for ite := shared.PlayerBullets.GetIterator(); ite.HasNext(); {
		p := (*bullet.Bullet)(ite.Next().GetData())
		p.Draw(screen)
	}

	for ite := shared.EnemyBullets.GetIterator(); ite.HasNext(); {
		p := (*bullet.Bullet)(ite.Next().GetData())
		p.Draw(screen)
	}
}

// checkCollision Проверка столкновений
func (stg *Scene) checkCollision() {

	for ite := shared.EnemyBullets.GetIterator(); ite.HasNext(); {
		e := (*bullet.Bullet)(ite.Next().GetData())
		if utils.IsCollideWith(stg.player, e) == true {
			//do something
		}
	}
}
