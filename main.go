package main

import (
	"github.com/AhEhIOhYou/project2/prj2"
	"github.com/hajimehoshi/ebiten/v2"
)

type size struct {
	width  int
	height int
}

func main() {
	window := &size{640, 640}
	ebiten.SetWindowTitle("project2")
	ebiten.SetWindowSize(window.width, window.height)
	game, err := prj2.NewGame()
	if err != nil {
		panic(err)
	}
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
