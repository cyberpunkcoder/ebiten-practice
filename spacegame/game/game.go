package game

import (
	"github.com/hajimehoshi/ebiten"
)

// All objects within the game
var (
	All []object

	err error
)

// Object shown on screen
type Object interface {
	GetPos() (float64, float64)
	Draw(screen *ebiten.Image, op *ebiten.DrawImageOptions)
}

type object struct {
	image            *ebiten.Image
	xPos, yPos, rPos float64
	xSpd, ySpd, rSpd float64
}

func (obj object) Draw(screen *ebiten.Image, op *ebiten.DrawImageOptions) {
	imgWidth, imgHeight := obj.image.Size()
	op.GeoM.Translate(obj.xPos-(float64(imgWidth)/2), obj.yPos-(float64(imgHeight/2)))
	screen.DrawImage(obj.image, op)
}
