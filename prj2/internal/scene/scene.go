package scene

import (
	"image/color"
	"log"
	"time"
	"unsafe"

	"github.com/AhEhIOhYou/project2/prj2/internal/actors"
	"github.com/AhEhIOhYou/project2/prj2/internal/bullet"
	"github.com/AhEhIOhYou/project2/prj2/internal/eventmanager"
	"github.com/AhEhIOhYou/project2/prj2/internal/fields"
	"github.com/AhEhIOhYou/project2/prj2/internal/inputs"
	"github.com/AhEhIOhYou/project2/prj2/internal/shared"
	"github.com/AhEhIOhYou/project2/prj2/internal/ui"
	"github.com/AhEhIOhYou/project2/prj2/internal/utils"
	"github.com/AhEhIOhYou/project2/prj2/internal/weapon"
	"github.com/hajimehoshi/ebiten/v2"
)

type gameState int

const (
	maxPlayerShot              = 500
	maxEnemyShot               = 100000
	damagePerHit               = 10
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
	enemy2 *actors.Enemy

	eventManager *eventmanager.EventManager
}

// NewScene вернет стандартную сцену
func NewScene() *Scene {
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
	em := eventmanager.NewEventManager()
	stg.input = inputs.New()
	stg.field = field
	stg.eventManager = em

	stg.player = actors.NewPlayer(field, shared.PlayerBullets)
	stg.player.SetMainWeapon(weapon.NewPlayerWeapon(bullet.PlayerWeaponShot))
	stg.player.SetWeaponCooldown(10)

	stg.enemy = actors.NewEnemy(field, shared.EnemyBullets)
	stg.enemy.SetMainWeapon(weapon.NewEnemyWeapon(bullet.EnemyWeaponShot))

	stg.enemy2 = actors.NewEnemy(field, shared.EnemyBullets)
	stg.enemy2.SetMainWeapon(weapon.NewEnemyWeapon(bullet.EnemyWeaponShot))

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
	stg.enemy.Init(300, 170, 1)
	stg.enemy2.Init(300, 170, 1)
	stg.time = time.Now()
	stg.LoadEvents()
	stg.eventManager.Time = time.Now()
}

// Update обновляет состояние сцены (актеров и окружения)
func (stg *Scene) Update() {
	input := stg.input

	stg.checkCollision()

	stg.eventManager.Execute()

	input.Update()
	if input.Reload {
		stg.setupGame()
	}

	stg.player.Action(input.Horizontal, input.Vertical, input.Fire, input.Focus)
	if input.Fire {
		stg.player.FireWeapon(270, 10, []int{-5, -1, 1, 5}, [][]float64{{-10, 0}, {-5, 0}, {5, 0}, {10, 0}})
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

func (stg *Scene) LoadEvents() {

	stg.eventManager.AddEvent(new(eventmanager.Event).
		OnTime(1).
		Actor(stg.enemy2).
		Fire(true).
		Weapon(90, 100, 0, 3).
		AddShot([]int{-6, -3, 0, 3, 6}, [][]float64{{0, -6}, {0, 4}, {0, 6}, {0, 4}, {0, -6}}).
		AddShot([]int{-3, 0, 3}, [][]float64{{0, 0}, {0, 0}, {0, 0}}).
		Duration(10).
		UpdateStatus("new"))

}

// Draw отрисовывает всех действиующих лиц сцены
func (stg *Scene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 0x10, G: 0x10, B: 0x30, A: 0xff})
	stg.field.Draw(screen)
	stg.player.Draw(screen)

	stg.enemy.Draw(screen)
	stg.enemy2.Draw(screen)

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
			e.OnHit()
			log.Println("HIT")
			lifePoints := stg.player.GetLifePoints()
			if lifePoints == 10 {
				log.Println("ПОМЕР")
			}
			stg.player.SetLifePoints(lifePoints - damagePerHit)
		}
	}
}
