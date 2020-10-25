/*
author: cyberpunkprogrammer
date: 10-25-2020
tutorial: https://medium.com/@chrisandrews_76960/2d-game-development-in-golang-part-2-5372e9338488
*/

package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	err        error
	background *ebiten.Image
)

// load background image from assets/space.png
func init() {
	background, _, err = ebitenutil.NewImageFromFile("assets/space.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}
}

// update the screen with background image
func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(background, op)

	return nil
}

// main loop
func main() {
	if err := ebiten.Run(update, 640, 480, 1, "Hello world!"); err != nil {
		log.Fatal(err)
	}
}
