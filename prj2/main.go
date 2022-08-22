package prj2

import (
	"github.com/AhEhIOhYou/project2/prj2/internal/scene"
	"github.com/AhEhIOhYou/project2/prj2/internal/sprite"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

// Scene представляет интерфейс сцены
type Scene interface {
	Update()
	Draw(screen *ebiten.Image)
}

var (
	currentScene    Scene = nil
	screenWidth           = 640
	screenHeight          = 640
	isInitialized         = false
	isWindowSizeSet       = false
)

// Update обновляет сцену каждый тик (tps)
func (g *Game) Update() error {
	if isWindowSizeSet && !isInitialized {
		sprite.LoadSprites()
		currentScene = scene.NewScene()
		isInitialized = true
		return nil
	}

	if currentScene != nil {
		currentScene.Update()
	}

	return nil
}

// Draw отрисовка сцены, вызывается каждый фрейм (fps)
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x10, 0x10, 0x30, 0xff})
	if currentScene != nil {
		currentScene.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	screenHeight = int(float64(screenWidth) / float64(outsideWidth) * float64(outsideHeight))
	isWindowSizeSet = true
	return screenWidth, screenHeight
}

// NewGame создает новую стуктуры игры
func NewGame() (*Game, error) {
	game := &Game{}
	return game, nil
}
