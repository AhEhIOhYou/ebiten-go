package input

import (
	"github.com/AhEhIOhYou/project2/prj2/internal/paint"
	"image/color"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	minAlphaTranslate = 0xcc
)

// Panel represents virtual keyboard
type Panel struct {
	screenWidth  int
	screenHeight int
	panelSize    int
	keySize      int
	panelNum     int
	color        color.RGBA
	offsetImage  *ebiten.Image
	animateAlpha int
	lastTickTime time.Time
}

// NewPanel returns Panel
func NewPanel(screenWidth, screenHeight int) *Panel {
	p := &Panel{}

	p.screenWidth = screenWidth
	p.screenHeight = screenHeight

	// Prepare an offset image for Panel
	p.keySize = 20
	p.panelNum = 5
	p.panelSize = p.keySize * p.panelNum

	// color setting
	p.color = color.RGBA{G: 0xff, A: 0xff}
	p.animateAlpha = -2
	p.lastTickTime = time.Now()

	p.preparePanel()

	return p
}

func (p *Panel) preparePanel() {
	offsetImage := ebiten.NewImage(p.panelSize, p.panelSize)

	// draw keys
	for i := 0; i < p.panelNum; i++ {
		for j := 0; j < p.panelNum; j++ {
			x := i * p.keySize
			y := j * p.keySize
			paint.DrawRect(offsetImage, paint.Rect{X: x, Y: y, W: p.keySize, H: p.keySize}, p.color, 1)
		}
	}

	p.offsetImage = offsetImage
}

// UpdatePanel updates the state of the panel
func (p *Panel) UpdatePanel() {
	p.updateColor()
}

func (p *Panel) updateColor() {
	// animate the panel color
	clr := p.color
	a := clr.A
	a = uint8(math.Min(math.Max(float64(a)+float64(p.animateAlpha), 0), minAlphaTranslate))
	clr.A = a
	if a == minAlphaTranslate || a == 0 {
		p.animateAlpha *= -1
	}
	p.color = clr

}

// DrawPanel draws panel
func (p *Panel) DrawPanel(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	// set position
	op.GeoM.Translate(0, float64(p.screenHeight-p.panelSize))

	// set color
	c := p.color
	r := float64(c.R) / 0xff
	g := float64(c.G) / 0xff
	b := float64(c.B) / 0xff
	a := float64(c.A) / 0xff * -1
	op.ColorM.Translate(r, g, b, a)

	screen.DrawImage(p.offsetImage, op)
}
