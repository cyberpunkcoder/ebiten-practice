/*
author: cyberpunkprogrammer
start date: 10-30-2020
*/

package main

import (
	"fmt"
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	screenWidth, screenHeight, scale = 640, 480, 4
)

// Game struct for ebiten
type Game struct {
	count      int
	playerShip Ship
	objects    []*Object
}

func init() {
	initImages()
	initSounds()
}

func newGame() *Game {
	g := &Game{}
	g.init()
	return g
}

func (g *Game) init() {
	g.playerShip = NewShip(float64(screenWidth/2), float64(screenHeight/2))
	g.objects = append(g.objects, &g.playerShip.Object)
}

func (g *Game) control() {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		g.playerShip.fwdThrustersOn()
	} else if !(ebiten.IsKeyPressed(ebiten.KeyUp) && !ebiten.IsKeyPressed(ebiten.KeyW)) {
		g.playerShip.fwdThrustersOff()
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		g.playerShip.revThrustersOn()
	} else if !(ebiten.IsKeyPressed(ebiten.KeyDown) && ebiten.IsKeyPressed(ebiten.KeyS)) {
		g.playerShip.revThrustersOff()
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		g.playerShip.ccwThrustersOn()
	} else if !(ebiten.IsKeyPressed(ebiten.KeyLeft) && ebiten.IsKeyPressed(ebiten.KeyA)) {
		g.playerShip.ccwThrustersOff()
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		g.playerShip.cwThrustersOn()
	} else if !(ebiten.IsKeyPressed(ebiten.KeyRight) && ebiten.IsKeyPressed(ebiten.KeyD)) {
		g.playerShip.cwThrustersOff()
	}
}

// Layout the screen
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

// Update the logical state
func (g *Game) Update() error {
	g.count++
	g.control()

	for i := range g.objects {
		g.objects[i].Update()
	}

	updateSound()

	return nil
}

// Draw the screen
func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	for i := range g.objects {
		g.objects[i].Draw(screen, op, g.count)

		if g.objects[i] == &g.playerShip.Object {

			frame := (g.count / 2) % 2

			if g.playerShip.ccwThrusters {
				screen.DrawImage(rcsfl.SubImage(image.Rect(frame*32, 0, 32+(frame*32), 32)).(*ebiten.Image), op)
				screen.DrawImage(rcsbr.SubImage(image.Rect(frame*32, 0, 32+(frame*32), 32)).(*ebiten.Image), op)
			}

			if g.playerShip.cwThrusters {
				screen.DrawImage(rcsfr.SubImage(image.Rect(frame*32, 0, 32+(frame*32), 32)).(*ebiten.Image), op)
				screen.DrawImage(rcsbl.SubImage(image.Rect(frame*32, 0, 32+(frame*32), 32)).(*ebiten.Image), op)
			}

			if g.playerShip.fwdThrusters {
				if !g.playerShip.cwThrusters {
					screen.DrawImage(rcsbl.SubImage(image.Rect(frame*32, 0, 32+(frame*32), 32)).(*ebiten.Image), op)
				}

				if !g.playerShip.ccwThrusters {
					screen.DrawImage(rcsbr.SubImage(image.Rect(frame*32, 0, 32+(frame*32), 32)).(*ebiten.Image), op)
				}
			}

			if g.playerShip.revThrusters {
				if !g.playerShip.ccwThrusters {
					screen.DrawImage(rcsfl.SubImage(image.Rect(frame*32, 0, 32+(frame*32), 32)).(*ebiten.Image), op)
				}

				if !g.playerShip.cwThrusters {
					screen.DrawImage(rcsfr.SubImage(image.Rect(frame*32, 0, 32+(frame*32), 32)).(*ebiten.Image), op)
				}
			}
		}
	}

	ebitenutil.DebugPrintAt(screen, "Controls = ESC, W, A, S, D", 0, 0)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("xPos = %f", g.playerShip.xPos), 0, 12)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("yPos = %f", g.playerShip.yPos), 0, 24)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("rPos = %f", g.playerShip.rPos), 0, 36)
}

func main() {
	ebiten.SetFullscreen(true)

	// scale up pixel art for aesthetics
	screenWidth, screenHeight = ebiten.ScreenSizeInFullscreen()
	screenWidth /= scale
	screenHeight /= scale

	if err := ebiten.RunGame(newGame()); err != nil {
		panic(err)
	}
}
