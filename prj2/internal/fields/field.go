package fields

import (
	"github.com/AhEhIOhYou/project2/prj2/internal/paint"
	"github.com/AhEhIOhYou/project2/prj2/internal/sprite"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	fieldWidth  = 640
	fieldHeight = 640
)

// Field представление игрового поля
type Field struct {
	x             int
	y             int
	width         int
	height        int
	boundaryImage *ebiten.Image
}

// NewField создание нового поля
func NewField() *Field {
	f := &Field{}
	f.x = fieldWidth / 2
	f.y = fieldHeight / 2
	f.width = fieldWidth
	f.height = fieldHeight

	borderColor := color.RGBA{0xff, 0, 0, 0x50}
	offsetImage := ebiten.NewImage(f.width, f.height)
	paint.DrawRect(offsetImage, paint.Rect{X: 0, Y: 0, W: f.width, H: f.height}, borderColor, 1)
	f.boundaryImage = offsetImage

	return f
}

// Draw отрисовывает поле
func (f *Field) Draw(screen *ebiten.Image) {
	sprite.Background.SetPosition(float64(f.x), float64(f.y))
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(f.boundaryImage, op)
}

func (f *Field) GetLeft() int {
	return f.x - f.width/2
}

func (f *Field) GetTop() int {
	return f.y - f.height/2
}

func (f *Field) GetRight() int {
	return f.x + f.width/2
}

func (f *Field) GetBottom() int {
	return f.y + f.height/2
}
