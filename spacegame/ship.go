package main

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// Ship spaceShip
type Ship struct {
	Object
	image            *ebiten.Image
	xPos, yPos, rPos float64
	xSpd, ySpd, rSpd float64
	rmax             float64
	vmax             float64
	cwThrusters      bool
	ccwThrusters     bool
	fwdThrusters     bool
	revThrusters     bool
}

// NewShip at x, y coordinates
func NewShip(x float64, y float64) *Ship {
	return &Ship{
		image: shipImage,
		xPos:  x,
		yPos:  y,
		xSpd:  0,
		ySpd:  0,
		rSpd:  0,
		rmax:  10,
		vmax:  1,
	}
}

func (ship *Ship) Update() {
	ship.xPos += ship.xSpd
	ship.yPos += ship.ySpd
	ship.rPos = math.Mod(ship.rPos+ship.rSpd, 360)
}

func (ship *Ship) Draw(screen *ebiten.Image, op *ebiten.DrawImageOptions, g *Game) {
	imgWidth, imgHeight := ship.image.Size()
	op.GeoM.Translate(-float64(imgWidth)/2, -float64(imgHeight)/2)
	op.GeoM.Rotate(float64(ship.rPos) * 2 * math.Pi / 360)
	op.GeoM.Translate(ship.xPos, ship.yPos)
	screen.DrawImage(ship.image, op)

	frame := (g.count / 2) % 2

	if ship.ccwThrusters {
		screen.DrawImage(rcsfl.SubImage(image.Rect(frame*32, 0, 32+(frame*32), 32)).(*ebiten.Image), op)
		screen.DrawImage(rcsbr.SubImage(image.Rect(frame*32, 0, 32+(frame*32), 32)).(*ebiten.Image), op)
	}

	if ship.cwThrusters {
		screen.DrawImage(rcsfr.SubImage(image.Rect(frame*32, 0, 32+(frame*32), 32)).(*ebiten.Image), op)
		screen.DrawImage(rcsbl.SubImage(image.Rect(frame*32, 0, 32+(frame*32), 32)).(*ebiten.Image), op)
	}

	if ship.fwdThrusters {
		if !ship.cwThrusters {
			screen.DrawImage(rcsbl.SubImage(image.Rect(frame*32, 0, 32+(frame*32), 32)).(*ebiten.Image), op)
		}

		if !ship.ccwThrusters {
			screen.DrawImage(rcsbr.SubImage(image.Rect(frame*32, 0, 32+(frame*32), 32)).(*ebiten.Image), op)
		}
	}

	if ship.revThrusters {
		if !ship.ccwThrusters {
			screen.DrawImage(rcsfl.SubImage(image.Rect(frame*32, 0, 32+(frame*32), 32)).(*ebiten.Image), op)
		}

		if !ship.cwThrusters {
			screen.DrawImage(rcsfr.SubImage(image.Rect(frame*32, 0, 32+(frame*32), 32)).(*ebiten.Image), op)
		}
	}
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
