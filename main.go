package main

import (
	"github.com/AhEhIOhYou/project2/prj2"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game, err := prj2.NewGame()
	if err != nil {
		panic(err)
	}
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
