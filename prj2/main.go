package prj2

import (
	"bytes"
	"github.com/AhEhIOhYou/project2/prj2/resources/images"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
	"log"
)

type Sprites struct {
	frameOX     int
	frameOY     int
	frameWidth  int
	frameHeight int
	frameCount  int
}

var ExplodeSprites = Sprites{
	frameOX:     0,
	frameOY:     0,
	frameWidth:  32,
	frameHeight: 32,
	frameCount:  10,
}

const (
	screenWidth  = 640
	screenHeight = 440
)

var (
	runnerImage *ebiten.Image
)

type Game struct {
	count int
}

// Update updates a game by one tick. The given argument represents a screen image.
func (g *Game) Update() error {
	g.count++
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(-float64(ExplodeSprites.frameWidth)/2, -float64(ExplodeSprites.frameHeight)/2)
	op.GeoM.Translate(screenWidth/2, screenHeight/2)

	i := (g.count / 2) % ExplodeSprites.frameCount
	sx, sy := ExplodeSprites.frameOX+i*ExplodeSprites.frameWidth, ExplodeSprites.frameOY

	img := runnerImage.SubImage(
		image.Rect(
			sx,
			sy,
			sx+ExplodeSprites.frameWidth,
			sy+ExplodeSprites.frameHeight)).(*ebiten.Image)

	screen.DrawImage(img, op)
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

	game := &Game{}

	return game, nil
}
