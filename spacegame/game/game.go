package game

import (
	"fmt"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// All objects within the game
var (
	All []object

	err error
)

// Object shown on screen
type Object interface {
	Update()
	Draw(screen *ebiten.Image, op *ebiten.DrawImageOptions)
}

type object struct {
	image            *ebiten.Image
	xPos, yPos, rPos float64
	xSpd, ySpd, rSpd float64
}

func (obj *object) Update() {
	obj.xPos += obj.xSpd
	obj.yPos += obj.ySpd
	obj.rPos = math.Mod(obj.rPos+obj.rSpd, 360)
}

func (obj *object) Draw(screen *ebiten.Image, op *ebiten.DrawImageOptions) {
	imgWidth, imgHeight := obj.image.Size()
	op.GeoM.Translate(-float64(imgWidth)/2, -float64(imgHeight)/2)
	op.GeoM.Rotate(float64(obj.rPos) * 2 * math.Pi / 360)
	op.GeoM.Translate(obj.xPos, obj.yPos)
	screen.DrawImage(obj.image, op)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("xPos = %f", obj.xPos), 0, 0)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("yPos = %f", obj.yPos), 0, 12)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("rPos = %f", obj.rPos), 0, 24)
}
