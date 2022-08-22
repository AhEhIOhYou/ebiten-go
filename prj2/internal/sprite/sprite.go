package sprite

import (
	"bytes"
	"github.com/AhEhIOhYou/project2/prj2/internal/resources/images"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	Background   *Sprite
	Player       *Sprite
	PlayerBullet *Sprite
	EnemyBullet  *Sprite
)

type frame struct {
	w int
	h int
}

type position struct {
	x float64
	y float64
}

type size struct {
	w int
	h int
}

// Sprite представляет спрайт
type Sprite struct {
	image     *ebiten.Image
	subImages []*ebiten.Image
	size      size
	frame     frame
	position  position
	index     int
	length    int
}

// NewSprite создает новый спрайт из изображения и количества кадров
func NewSprite(img *image.Image, frameNum int) *Sprite {
	originalImage := ebiten.NewImageFromImage(*img)

	sprite := &Sprite{}
	sprite.image = originalImage
	sprite.size.w, sprite.size.h = originalImage.Size()
	sprite.length = frameNum
	sprite.index = 0
	sprite.frame.w = sprite.size.w / frameNum
	sprite.frame.h = sprite.size.h

	subImages := []*ebiten.Image{}
	for i := 0; i < frameNum; i++ {
		x := sprite.frame.w * i
		y := 0
		rect := image.Rect(x, y, x+sprite.frame.w, y+sprite.frame.h)
		sub := originalImage.SubImage(rect)
		ebitenImage := ebiten.NewImageFromImage(sub)
		subImages = append(subImages, ebitenImage)
	}
	sprite.subImages = subImages

	return sprite
}

// Size возвращает ширину и высоту кадров спрайта
func (s *Sprite) Size() (int, int) {
	return s.frame.w, s.frame.h
}

// SetPosition устанавливает позицию спрайта
func (s *Sprite) SetPosition(x, y float64) {
	s.position.x = x
	s.position.y = y
}

// SetIndex устанавливает текущий фрейм спрайта
func (s *Sprite) SetIndex(index int) {
	s.index = index
}

// Index возвращает индекс текущего фрейма
func (s *Sprite) Index() int {
	return s.index
}

// Length возвращает количество фреймов у спрайта
func (s *Sprite) Length() int {
	return s.length
}

// Draw отрисовывает спрайт
func (s *Sprite) Draw(screen *ebiten.Image) {
	w, h := s.Size()
	x := s.position.x
	y := s.position.y
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(x-float64(w)/2, y-float64(h)/2)

	screen.DrawImage(s.subImages[s.index], op)
}

// LoadSprites загружает спрайты из изображений
func LoadSprites() {
	Player = createSprite(&images.P_ROBO1, 8)
	Background = createSprite(&images.SPACE5, 1)
	PlayerBullet = createSprite(&images.SHOT2, 8)
	EnemyBullet = createSprite(&images.ESHOT10_1, 1)
}

// создает изображения из байтового массива
func createSprite(rawImage *[]byte, frameNum int) *Sprite {
	img, _, _ := image.Decode(bytes.NewReader(*rawImage))
	return NewSprite(&img, frameNum)
}
