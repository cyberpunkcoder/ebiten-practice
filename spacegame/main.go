/*
author: cyberpunkprogrammer
date: 10-30-2020
*/

package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	err                              error
	screenWidth, screenHeight, scale = 640, 480, 2
	backgroundImage                  *ebiten.Image
	playerImage                      *ebiten.Image
	bulletImage                      *ebiten.Image
	enemyImage                       *ebiten.Image
	player                           object
)

type object struct {
	image            *ebiten.Image
	xPos, yPos       float64
	xSpd, ySpd, rSpd float64
}

func init() {
	//ebiten.SetFullscreen(true)

	// scale up pixel art for aesthetics
	screenWidth, screenHeight = ebiten.ScreenSizeInFullscreen()
	screenWidth /= scale
	screenHeight /= scale

	playerImage, _, err = ebitenutil.NewImageFromFile("assets/player.png", ebiten.FilterDefault)
	player.image = playerImage

	// position player at center of screen
	playerImageX, playerImageY := player.image.Size()
	player.xPos = float64((screenWidth / 2) - (playerImageX / 2))
	player.yPos = float64((screenHeight / 2) - (playerImageY / 2))

	// make player still
	player.xSpd, player.ySpd, player.rSpd = 0, 0, 0
}

func update(screen *ebiten.Image) error {

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	playerOp := &ebiten.DrawImageOptions{}
	playerOp.GeoM.Translate(player.xPos, player.yPos)
	screen.DrawImage(player.image, playerOp)

	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, float64(scale), "Space Game"); err != nil {
		log.Fatal(err)
	}
}
