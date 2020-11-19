/*
author: cyberpunkprogrammer
start date: 10-30-2020
*/

package main

import (
	"fmt"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	screenWidth, screenHeight, scale = 640, 480, 3
)

// Game struct for ebiten
type Game struct {
	count      int
	playerShip *Ship
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
	objects = append(objects, g.playerShip)
}

func (g *Game) control() {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		g.playerShip.fwdThrustersOn()
	} else if inpututil.IsKeyJustReleased(ebiten.KeyW) {
		g.playerShip.fwdThrustersOff()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		g.playerShip.revThrustersOn()
	} else if inpututil.IsKeyJustReleased(ebiten.KeyS) {
		g.playerShip.revThrustersOff()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		g.playerShip.ccwThrustersOn()
	} else if inpututil.IsKeyJustReleased(ebiten.KeyA) {
		g.playerShip.ccwThrustersOff()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		g.playerShip.cwThrustersOn()
	} else if inpututil.IsKeyJustReleased(ebiten.KeyD) {
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

	for _, o := range objects {
		o.Update()
	}

	updateSound()

	return nil
}

// Draw the screen
func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	for _, o := range objects {
		o.Draw(screen, op, g)
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
