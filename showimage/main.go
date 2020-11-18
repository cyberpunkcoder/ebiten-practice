/*
author: cyberpunkprogrammer
date: 11-10-2020
tutorial: https://medium.com/@chrisandrews_76960/2d-game-development-in-golang-part-2-5372e9338488
Updated for ebiten v2
*/

package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	err                       error
	screenWidth, screenHeight = 640, 480
	background                *ebiten.Image
)

// Game struct
type Game struct{}

// Layout the scene
func (g *Game) Layout(ousideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

// Update the logical state
func (g *Game) Update() error {

	return nil
}

// Draw the screen
func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(background, op)
}

// load background image from assets/space.png
func init() {
	background, _, err = ebitenutil.NewImageFromFile("assets/space.png")

	if err != nil {
		log.Fatal(err)
	}
}

// main loop
func main() {
	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
