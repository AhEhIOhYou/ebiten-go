package shooting

import (
	"github.com/AhEhIOhYou/project2/prj2/internal/player"
	"github.com/AhEhIOhYou/project2/prj2/internal/scene/shooting/input"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	myPlayer     *player.Player
	myInput      *input.Input
	screenWidth  = 0
	screenHeight = 0
)

// Shooting represents shooting scene
type Shooting struct {
}

// NewOptions represents options for New func
type NewOptions struct {
	ScreenWidth  int
	ScreenHeight int
}

// New returns new Shooting struct
func New(options NewOptions) *Shooting {
	shooting := &Shooting{}

	myPlayer = player.New()
	myInput = input.New()

	return shooting
}

// Update updates the scene
func (shooting *Shooting) Update() {
	myInput.Update()
	myPlayer.Move(myInput.Horizontal, myInput.Vertical, false, myInput.Focus)
}

// Draw draws the scene
func (shooting *Shooting) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x10, 0x10, 0x30, 0xff})
	myPlayer.Draw(screen)
}
