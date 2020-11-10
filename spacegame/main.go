/*
author: cyberpunkprogrammer
start date: 10-30-2020
*/

package main

import (
	"log"
	"os"

	"github.com/cyberpunkprogrammer/ebiten-practice/spacegame/game"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Err error created by the game
var (
	Err                              error
	screenWidth, screenHeight, scale = 640, 480, 3
)

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

func update(screen *ebiten.Image) error {
	game.Count++

	control()

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	op := &ebiten.DrawImageOptions{}

	for i := range game.All {
		game.All[i].Draw(screen, op)
		game.All[i].Update()
	}

	game.UpdateSound()

	ebitenutil.DebugPrintAt(screen, "Controls = ESC, W, A, S, D", 0, 0)

	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, float64(scale), "Space Game"); err != nil {
		log.Fatal(err)
	}
}
