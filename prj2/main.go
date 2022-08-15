package prj2

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

const (
	screenWidth  = 640
	screenHeight = 400
)

type Game struct{}

func (g Game) Update() error {
	return nil
}

func (g Game) Draw(screen *ebiten.Image) {

}

func (g Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenWidth, screenHeight
}

func NewGame() (*Game, error) {

	game := &Game{}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("project2")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

	return game, nil
}
