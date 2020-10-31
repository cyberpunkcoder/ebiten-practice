/*
author: cyberpunkprogrammer
date: 10-30-2020
*/

package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
)

var (
	err                       error
	screenWidth, screenHeight = 640, 480
	background                *ebiten.Image
	player                    *ebiten.Image
	bullet                    *ebiten.Image
	enemy                     *ebiten.Image
)

func init() {
	//ebiten.SetFullscreen(true)
	screenWidth, screenHeight = ebiten.ScreenSizeInFullscreen()
}

func update(screen *ebiten.Image) error {
	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "Space Game"); err != nil {
		log.Fatal(err)
	}
}
