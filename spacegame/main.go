/*
author: cyberpunkprogrammer
start date: 10-30-2020
*/

package main

import (
	"log"

	"github.com/cyberpunkprogrammer/ebiten-practice/spacegame/game"
	"github.com/hajimehoshi/ebiten"
)

// Err error created by the game
var (
	Err                              error
	screenWidth, screenHeight, scale = 640, 480, 2
)

func init() {
	//ebiten.SetFullscreen(true)

	// scale up pixel art for aesthetics
	screenWidth, screenHeight = ebiten.ScreenSizeInFullscreen()
	screenWidth /= scale
	screenHeight /= scale

	game.CreatePlayer(float64(screenWidth/2), float64(screenHeight/2))
	game.Player.DecreaseRspd()
}

func control() {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		game.Player.DecreaseRspd()
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		game.Player.IncreaseRspd()
	}
}

func update(screen *ebiten.Image) error {
	control()

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	op := &ebiten.DrawImageOptions{}

	game.Player.Draw(screen, op)

	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, float64(scale), "Space Game"); err != nil {
		log.Fatal(err)
	}
}
