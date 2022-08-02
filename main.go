package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/colornames"
	"log"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Red)
	ebitenutil.DebugPrint(screen, "TEST")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(1280, 960)
	ebiten.SetWindowTitle("ЙУХ")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
