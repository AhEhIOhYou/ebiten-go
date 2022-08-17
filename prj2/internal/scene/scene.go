package scene

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Scene represents scene interface
type Scene interface {
	Update()
	Draw(screen *ebiten.Image)
}
