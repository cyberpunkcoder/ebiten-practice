/*
author: cyberpunkprogrammer
date: 11-10-2020
tutorial: https://medium.com/@chrisandrews_76960/2d-game-development-in-golang-part-1-5e2c11a513ed
Updated for ebiten v2
*/

package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Game basic values
type Game struct{}

// Update the logical state
func (g *Game) Update() error {

	return nil
}

// Draw to the screen
func (g *Game) Draw(screen *ebiten.Image) {

	ebitenutil.DebugPrint(screen, "Hello world!")
}

// Layout the scene
func (g *Game) Layout(ousideWidth, outsideHeight int) (screenWidth, screenHeight int) {

	return 320, 240
}

// main loop
func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello world!")

	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
