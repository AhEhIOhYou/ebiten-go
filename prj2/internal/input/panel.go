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
	panelWidth   int
	panelHeight  int
	X            int
	Y            int
	color        color.RGBA
	offsetImage  *ebiten.Image
	animateAlpha int
	lastTickTime time.Time
}

// NewPanel returns Panel
func NewPanel(x, y, panelWidth, panelHeight int) *Panel {
	p := &Panel{}

	p.X = x
	p.Y = y
	p.panelWidth = panelWidth
	p.panelHeight = panelHeight

	// color setting
	p.color = color.RGBA{G: 0xff, A: 0xff}
	p.animateAlpha = -2
	p.lastTickTime = time.Now()

	p.preparePanel()

	return p
}

func (p *Panel) preparePanel() {
	offsetImage := ebiten.NewImage(p.panelWidth, p.panelHeight)

	rect := paint.Rect{X: 0, Y: 0, W: p.panelWidth, H: p.panelHeight}
	paint.FillRect(offsetImage, rect, p.color)

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

	op.GeoM.Translate(float64(p.X), float64(p.Y))

	// set color
	c := p.color
	r := float64(c.R) / 0xff
	g := float64(c.G) / 0xff
	b := float64(c.B) / 0xff
	a := float64(c.A) / 0xff * -1
	op.ColorM.Translate(r, g, b, a)

	screen.DrawImage(p.offsetImage, op)
}
