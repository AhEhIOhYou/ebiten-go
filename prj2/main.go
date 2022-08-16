package prj2

import (
	"github.com/AhEhIOhYou/project2/prj2/internal/player"
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
)

var (
	myPlayer      *player.Player
	screenWidth   = 640
	screenHeight  = 440
	isInitialized = false
)

type Game struct{}

// Update updates a game by one tick. The given argument represents a screen image.
func (g *Game) Update() error {
	if !isInitialized {
		myPlayer = player.New()
		isInitialized = true
		return nil
	}

	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	myPlayer.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	screenHeight = int(float64(screenWidth) / float64(outsideWidth) * float64(outsideHeight))
	return screenWidth, screenHeight
}

func NewGame() (*Game, error) {

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("project2")

	game := &Game{}

	return game, nil
}
