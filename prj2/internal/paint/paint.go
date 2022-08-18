package paint

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

// Rect представляет область с изображением
type Rect struct {
	X, Y, W, H int
}

// FillRect заполняет область изображения
func FillRect(target *ebiten.Image, r Rect, clr color.Color) {
	img := ebiten.NewImage(r.W, r.H)
	img.Fill(clr)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(r.X), float64(r.Y))

	target.DrawImage(img, op)
}

// DrawRect отрисовывает область
func DrawRect(target *ebiten.Image, r Rect, clr color.Color, width int) {
	for i := r.X; i < r.X+r.W; i++ {
		for j := 0; j < width; j++ {
			target.Set(i, r.Y+j, clr)
			target.Set(i, r.Y+r.H-j-1, clr)
		}
	}

	for i := r.Y; i < r.Y+r.H; i++ {
		for j := 0; j < width; j++ {
			target.Set(r.X+j, i, clr)
			target.Set(r.X+r.W-j-1, i, clr)
		}
	}
}
