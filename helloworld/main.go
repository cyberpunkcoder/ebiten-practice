/*
author: cyberpunkprogrammer
date: 10-25-2020
tutorial: https://medium.com/@chrisandrews_76960/2d-game-development-in-golang-part-1-5e2c11a513ed
*/

package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// update the text hello world on the screen
func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	ebitenutil.DebugPrint(screen, "Hello world!")
	return nil
}

// main loop
func main() {
	if err := ebiten.Run(update, 320, 240, 2, "Hello world!"); err != nil {
		log.Fatal(err)
	}
}
