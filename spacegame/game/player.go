package game

import (
	"log"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Player the user controls
var (
	Player player
)

type player struct {
	*object
	lives int
	rmax  float64
	vmax  float64
}

// Controls the player
type Controls interface {
	IncreaseRspd()
	DecreaseRspd()
	IncreaseVelocity()
	DecreaseVelocity()
}

// CreatePlayer at coordinates
func CreatePlayer(x float64, y float64) {

	playerImage, _, err := ebitenutil.NewImageFromFile("assets/player.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatal(err)
	}

	newObject := object{
		image: playerImage,
		xPos:  x,
		yPos:  y,
		xSpd:  0,
		ySpd:  0,
		rSpd:  0,
	}

	Player.lives = 3
	Player.rmax = 10
	Player.vmax = 1

	All = append(All, newObject)
	Player.object = &All[len(All)-1]
}

func (plr *player) IncreaseRspd() {
	if plr.rSpd < plr.rmax {
		plr.rSpd += 0.1
	}
}

func (plr *player) DecreaseRspd() {
	if plr.rSpd > -plr.rmax {
		plr.rSpd -= 0.1
	}
}

func (plr *player) IncreaseVelocity() {
	radAng := (plr.rPos + 90) * (math.Pi / 180)
	xSpd := plr.xSpd - 0.01*math.Cos(radAng)
	ySpd := plr.ySpd - 0.01*math.Sin(radAng)

	if math.Abs(xSpd)+math.Abs(ySpd) < plr.vmax {
		plr.xSpd = xSpd
		plr.ySpd = ySpd
	}
}

func (plr *player) DecreaseVelocity() {
	radAng := (plr.rPos + 90) * (math.Pi / 180)
	xSpd := plr.xSpd + 0.01*math.Cos(radAng)
	ySpd := plr.ySpd + 0.01*math.Sin(radAng)

	if math.Abs(xSpd)+math.Abs(ySpd) < plr.vmax {
		plr.xSpd = xSpd
		plr.ySpd = ySpd
	}
}
