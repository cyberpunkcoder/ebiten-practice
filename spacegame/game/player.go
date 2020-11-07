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

	rmax float64 = 10
	vmax float64 = 1
)

type player struct {
	object
	lives int
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
	Player.image, _, err = ebitenutil.NewImageFromFile("assets/player.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatal(err)
	}

	Player.xPos = x
	Player.yPos = y
	Player.xSpd = 0
	Player.ySpd = 0
	Player.rSpd = 0
	Player.lives = 3

	All = append(All, Player.object)
}

func (plr *player) IncreaseRspd() {
	if plr.rSpd < rmax {
		plr.rSpd += 0.1
	}
}

func (plr *player) DecreaseRspd() {
	if plr.rSpd > -rmax {
		plr.rSpd -= 0.1
	}
}

func (plr *player) IncreaseVelocity() {
	radAng := (plr.rPos + 90) * (math.Pi / 180)
	xSpd := plr.xSpd - 0.01*math.Cos(radAng)
	ySpd := plr.ySpd - 0.01*math.Sin(radAng)

	if math.Abs(xSpd)+math.Abs(ySpd) < vmax {
		plr.xSpd = xSpd
		plr.ySpd = ySpd
	}
}

func (plr *player) DecreaseVelocity() {
	radAng := (plr.rPos + 90) * (math.Pi / 180)
	xSpd := plr.xSpd + 0.01*math.Cos(radAng)
	ySpd := plr.ySpd + 0.01*math.Sin(radAng)

	if math.Abs(xSpd)+math.Abs(ySpd) < vmax {
		plr.xSpd = xSpd
		plr.ySpd = ySpd
	}
}
