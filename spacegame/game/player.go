package game

import (
	"log"
	"math"
)

// Player the user controls
var (
	Player player
)

type player struct {
	*object
	lives        int
	rmax         float64
	vmax         float64
	cwThrusters  bool
	ccwThrusters bool
	fwdThrusters bool
	revThrusters bool
}

// Controls the player
type Controls interface {
	CwThrustersOn()
	CwThrustersOff()
	CcwThrustersOn()
	CcwThrustersOff()
	FwdThrustersOn()
	FwdThrustersOff()
	RevThrustersOn()
	RevThrustersOff()
}

// CreatePlayer at coordinates
func CreatePlayer(x float64, y float64) {

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

// Turns on clockwise thrusters
func (plr *player) CwThrustersOn() {
	if !plr.cwThrusters {
		startThrusterSound()
	}

	if plr.rSpd < plr.rmax {
		plr.rSpd += 0.1

		if !plr.cwThrusters {
			startThrusterSound()
		}
	}

	plr.cwThrusters = true
}

// Turns off clockwise thrusters
func (plr *player) CwThrustersOff() {
	if !plr.isThrusting() {
		stopThrusterSound()
	}
	plr.cwThrusters = false
}

// Turns on counter-clockwise thrusters
func (plr *player) CcwThrustersOn() {
	if !plr.ccwThrusters {
		startThrusterSound()
	}

	if plr.rSpd > -plr.rmax {
		plr.rSpd -= 0.1
	}

	plr.ccwThrusters = true
}

// Turns off counter-clockwise thrusters
func (plr *player) CcwThrustersOff() {
	if !plr.isThrusting() {
		stopThrusterSound()
	}
	plr.ccwThrusters = false
}

// Turns on forward thrusters
func (plr *player) FwdThrustersOn() {

	if !plr.fwdThrusters {
		startThrusterSound()
		plr.fwdThrusters = true
	}

	radAng := (plr.rPos + 90) * (math.Pi / 180)
	xSpd := plr.xSpd - 0.01*math.Cos(radAng)
	ySpd := plr.ySpd - 0.01*math.Sin(radAng)

	if math.Abs(xSpd)+math.Abs(ySpd) < plr.vmax {
		plr.xSpd = xSpd
		plr.ySpd = ySpd
	}
}

// Turns off forward thrusters
func (plr *player) FwdThrustersOff() {
	if !plr.isThrusting() {
		stopThrusterSound()
	}
	plr.fwdThrusters = false
}

// Turns on reverse thrusters
func (plr *player) RevThrustersOn() {
	if !plr.revThrusters {
		startThrusterSound()
	}

	radAng := (plr.rPos + 90) * (math.Pi / 180)
	xSpd := plr.xSpd + 0.01*math.Cos(radAng)
	ySpd := plr.ySpd + 0.01*math.Sin(radAng)

	if math.Abs(xSpd)+math.Abs(ySpd) < plr.vmax {
		plr.xSpd = xSpd
		plr.ySpd = ySpd
	}

	plr.revThrusters = true
}

// Turns off reverse thrusters
func (plr *player) RevThrustersOff() {
	if !plr.isThrusting() {
		stopThrusterSound()
	}
	plr.revThrusters = false
}

func (plr *player) isThrusting() bool {
	return plr.fwdThrusters || plr.revThrusters || plr.cwThrusters || plr.ccwThrusters
}
