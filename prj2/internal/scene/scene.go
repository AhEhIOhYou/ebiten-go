package scene

import (
	"github.com/AhEhIOhYou/project2/prj2/internal/scene/actors"
	"github.com/AhEhIOhYou/project2/prj2/internal/scene/fields"
	"github.com/AhEhIOhYou/project2/prj2/internal/scene/inputs"
	"github.com/AhEhIOhYou/project2/prj2/internal/scene/tools"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"math/rand"
	"time"
)

const (
	maxPlayerShot = 200
	maxEnemyShot  = 200
)

// PlayerShooter представляет интерфейс оружия игрока
type PlayerShooter interface {
	Shot(x, y float64, degree int, playerShots []*actors.PlayerBullet)
}

// EnemyShooter представляет интерфейс оружия противника
type EnemyShooter interface {
	Shot(x, y float64, degree int, EnemyShooter []*actors.EnemyBullet)
}

var (
	input        *inputs.Input
	field        *fields.Field
	player       *actors.Player
	playerWeapon PlayerShooter
	enemyWeapon  EnemyShooter
	playerShots  [maxPlayerShot]*actors.PlayerBullet
	enemyShots   [maxEnemyShot]*actors.EnemyBullet
)

// Scene представляет сцену
type Scene struct {
	tick time.Time
}

// NewSceneOptions представляет настройки сцены
type NewSceneOptions struct {
	ScreenWidth  int
	ScreenHeight int
}

// NewScene вернет стандартную сцену
func NewScene(options NewSceneOptions) *Scene {
	stg := &Scene{}
	stg.tick = time.Now()
	initGame()
	return stg
}

// initGame инициализирует игру
func initGame() {
	rand.Seed(time.Now().Unix())
	input = inputs.New()
	field = fields.NewField()

	actors.SetBoundary(field)

	player = actors.NewPlayer()
	playerWeapon = &tools.PlayerWeapon{}
	for i := 0; i < len(playerShots); i++ {
		playerShots[i] = actors.NewPlayerShot()
	}

	enemyWeapon = &tools.EnemyWeapon{}
	for i := 0; i < len(enemyShots); i++ {
		enemyShots[i] = actors.NewEnemyShot()
	}
}

// Update обновляет состояние сцены (актеров и окружения)
func (stg *Scene) Update() {
	input.Update()
	checkCollision()
	player.Action(input.Horizontal, input.Vertical, input.Fire, input.Focus)

	enemyWeapon.Shot(100, 100, 90, enemyShots[:])
	if input.Fire {

		x, y := player.GetPosition()
		playerWeapon.Shot(x, y, player.GetNormalizedDegree(), playerShots[:])
	}

	for i := 0; i < len(playerShots); i++ {
		p := playerShots[i]
		if p.IsActive() == false {
			continue
		}
		p.Move()
	}

	for i := 0; i < len(enemyShots); i++ {
		p := enemyShots[i]
		if p.IsActive() == false {
			continue
		}
		p.Move()
	}
}

// Draw отрисовывает всех действиующих лиц сцены
func (stg *Scene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x10, 0x10, 0x30, 0xff})
	field.Draw(screen)
	player.Draw(screen)

	for i := 0; i < len(playerShots); i++ {
		p := playerShots[i]
		if p.IsActive() == false {
			continue
		}
		p.Draw(screen)
	}

	for i := 0; i < len(enemyShots); i++ {
		p := enemyShots[i]
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
