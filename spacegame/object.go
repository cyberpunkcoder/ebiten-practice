package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// Object in the game
type Object struct {
	image            *ebiten.Image
	xPos, yPos, rPos float64
	xSpd, ySpd, rSpd float64
}

// Update the object's state
func (obj *Object) Update() {
	obj.xPos += obj.xSpd
	obj.yPos += obj.ySpd
	obj.rPos = math.Mod(obj.rPos+obj.rSpd, 360)
}

// Draw the object on screen
func (obj *Object) Draw(screen *ebiten.Image, op *ebiten.DrawImageOptions, count int) {
	imgWidth, imgHeight := obj.image.Size()
	op.GeoM.Translate(-float64(imgWidth)/2, -float64(imgHeight)/2)
	op.GeoM.Rotate(float64(obj.rPos) * 2 * math.Pi / 360)
	op.GeoM.Translate(obj.xPos, obj.yPos)
	screen.DrawImage(obj.image, op)
}
