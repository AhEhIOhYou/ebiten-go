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

const (
	screenWidth  = 640
	screenHeight = 440
)

var (
	runnerImage *ebiten.Image
	testSprite  *sprite.Sprite
)

type Game struct{}

// Update updates a game by one tick. The given argument represents a screen image.
func (g *Game) Update() error {
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	testSprite.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
func NewGame() (*Game, error) {

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("project2")

	img, _, err := image.Decode(bytes.NewReader(images.EXPLODE_SMALL))
	if err != nil {
		log.Fatal(err)
	}

	runnerImage = ebiten.NewImageFromImage(img)

	testSprite = sprite.New(runnerImage, 1)
	testSprite.SetPosition(screenWidth/2, screenWidth/2)

	game := &Game{}

	return game, nil
}
