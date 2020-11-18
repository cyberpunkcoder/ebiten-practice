package main

import (
	"math"
)

// Ship spaceShip
type Ship struct {
	Object
	lives        int
	rmax         float64
	vmax         float64
	cwThrusters  bool
	ccwThrusters bool
	fwdThrusters bool
	revThrusters bool
}

// NewShip at x, y coordinates
func NewShip(x float64, y float64) Ship {

	obj := Object{
		image: shipImage,
		xPos:  x,
		yPos:  y,
		xSpd:  0,
		ySpd:  0,
		rSpd:  0,
	}

	ship := Ship{
		lives:  3,
		rmax:   10,
		vmax:   1,
		Object: obj,
	}

	return ship
}

// Turns on clockwise thrusters
func (ship *Ship) cwThrustersOn() {
	if !ship.cwThrusters {
		ship.cwThrusters = true
		startRcsSound()
	}

	if ship.rSpd < ship.rmax {
		ship.rSpd += 0.05
	}
}

// Turns off clockwise thrusters
func (ship *Ship) cwThrustersOff() {
	ship.cwThrusters = false
	if !ship.isThrusting() {
		stopRcsSound()
	}
}

// Turns on counter-clockwise thrusters
func (ship *Ship) ccwThrustersOn() {
	if !ship.ccwThrusters {
		ship.ccwThrusters = true
		startRcsSound()
	}

	if ship.rSpd > -ship.rmax {
		ship.rSpd -= 0.05
	}
}

// Turns off counter-clockwise thrusters
func (ship *Ship) ccwThrustersOff() {
	ship.ccwThrusters = false
	if !ship.isThrusting() {
		stopRcsSound()
	}
}

// Turns on forward thrusters
func (ship *Ship) fwdThrustersOn() {
	if !ship.fwdThrusters {
		ship.fwdThrusters = true
		startRcsSound()
	}

	radAng := (ship.rPos + 90) * (math.Pi / 180)
	xSpd := ship.xSpd - 0.01*math.Cos(radAng)
	ySpd := ship.ySpd - 0.01*math.Sin(radAng)

	if math.Abs(xSpd)+math.Abs(ySpd) < ship.vmax {
		ship.xSpd = xSpd
		ship.ySpd = ySpd
	}
}

// Turns off forward thrusters
func (ship *Ship) fwdThrustersOff() {
	ship.fwdThrusters = false
	if !ship.isThrusting() {
		stopRcsSound()
	}
}

// Turns on reverse thrusters
func (ship *Ship) revThrustersOn() {
	if !ship.revThrusters {
		ship.revThrusters = true
		startRcsSound()
	}

	radAng := (ship.rPos + 90) * (math.Pi / 180)
	xSpd := ship.xSpd + 0.01*math.Cos(radAng)
	ySpd := ship.ySpd + 0.01*math.Sin(radAng)

	if math.Abs(xSpd)+math.Abs(ySpd) < ship.vmax {
		ship.xSpd = xSpd
		ship.ySpd = ySpd
	}
}

// Turns off reverse thrusters
func (ship *Ship) revThrustersOff() {
	ship.revThrusters = false
	if !ship.isThrusting() {
		stopRcsSound()
	}
}

func (ship *Ship) isThrusting() bool {
	return ship.fwdThrusters || ship.revThrusters || ship.cwThrusters || ship.ccwThrusters
}
