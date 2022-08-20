package inputs

import (
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

// Input представляет управление игроком
type Input struct {
	Horizontal   float64
	Vertical     float64
	Fire         bool
	Focus        bool
	Reload       bool
	prevTickTime time.Time
}

// New инициализирует стуктуру управления
func New() *Input {
	gameInput := &Input{}
	gameInput.prevTickTime = time.Now()
	return gameInput
}

// Update обновляет состояние управления
func (i *Input) Update() {
	i.Reload = ebiten.IsKeyPressed(ebiten.KeyR)

	if time.Since(i.prevTickTime).Milliseconds() < 60 {
		return
	}
	i.prevTickTime = time.Now()

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		i.Vertical = 1
	} else if ebiten.IsKeyPressed(ebiten.KeyUp) {
		i.Vertical = -1
	} else {
		i.Vertical = 0
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		i.Horizontal = 1
	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		i.Horizontal = -1
	} else {
		i.Horizontal = 0
	}

	i.Focus = ebiten.IsKeyPressed(ebiten.KeyShiftLeft)
	i.Fire = ebiten.IsKeyPressed(ebiten.KeyZ)
}
