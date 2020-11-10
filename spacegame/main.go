/*
author: cyberpunkprogrammer
start date: 10-30-2020
*/

package main

import (
	"os"

	"github.com/cyberpunkprogrammer/ebiten-practice/spacegame/game"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	screenWidth, screenHeight, scale = 640, 480, 4
)

// Game struct for ebiten
type Game struct{}

func init() {
	ebiten.SetFullscreen(true)

	// scale up pixel art for aesthetics
	screenWidth, screenHeight = ebiten.ScreenSizeInFullscreen()
	screenWidth /= scale
	screenHeight /= scale

	game.CreatePlayer(float64(screenWidth/2), float64(screenHeight/2))
}

func control() {
	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		game.Player.FwdThrustersOn()
	} else if !(ebiten.IsKeyPressed(ebiten.KeyUp) && !ebiten.IsKeyPressed(ebiten.KeyW)) {
		game.Player.FwdThrustersOff()
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		game.Player.RevThrustersOn()
	} else if !(ebiten.IsKeyPressed(ebiten.KeyDown) && ebiten.IsKeyPressed(ebiten.KeyS)) {
		game.Player.RevThrustersOff()
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		game.Player.CcwThrustersOn()
	} else if !(ebiten.IsKeyPressed(ebiten.KeyLeft) && ebiten.IsKeyPressed(ebiten.KeyA)) {
		game.Player.CcwThrustersOff()
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		game.Player.CwThrustersOn()
	} else if !(ebiten.IsKeyPressed(ebiten.KeyRight) && ebiten.IsKeyPressed(ebiten.KeyD)) {
		game.Player.CwThrustersOff()
	}
}

// Layout the screen
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

// Update the logical state
func (g *Game) Update() error {

	return nil
}

// Draw the screen
func (g *Game) Draw(screen *ebiten.Image) {
	game.Count++

	control()

	op := &ebiten.DrawImageOptions{}

	for i := range game.All {
		game.All[i].Draw(screen, op)
		game.All[i].Update()
	}

	game.UpdateSound()

	ebitenutil.DebugPrintAt(screen, "Controls = ESC, W, A, S, D", 0, 0)
}

func main() {
	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
