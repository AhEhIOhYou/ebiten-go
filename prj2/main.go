package prj2

import (
	"bytes"
	"github.com/AhEhIOhYou/project2/prj2/internal/sprite"
	"github.com/AhEhIOhYou/project2/prj2/resources/images"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
	"log"
)

var (
	testSprite    *sprite.Sprite
	screenWidth   = 1280
	screenHeight  = 960
	isInitialized = false
)

type Game struct{}

// Update updates a game by one tick. The given argument represents a screen image.
func (g *Game) Update() error {
	if !isInitialized {
		img, _, err := image.Decode(bytes.NewReader(images.ESHOT10_1))
		if err != nil {
			log.Fatal(err)
		}
		testSprite = sprite.New(&img, 1)
		testSprite.SetIndex(0)
		isInitialized = true
		return nil
	}

	testSprite.SetPosition(float64(screenWidth/2), float64(screenHeight/2))
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	testSprite.Draw(screen)
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
