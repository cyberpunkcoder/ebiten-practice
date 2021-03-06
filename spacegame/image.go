package main

import (
	_ "image/png" // Required for ebitenutil.NewImageFromFile()
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	shipImage *ebiten.Image
	rcsfl     *ebiten.Image
	rcsfr     *ebiten.Image
	rcsbl     *ebiten.Image
	rcsbr     *ebiten.Image
	err       error
)

func initImages() {

	shipImage, _, err = ebitenutil.NewImageFromFile("assets/player.png")
	rcsfl, _, err = ebitenutil.NewImageFromFile("assets/rcsfl.png")
	rcsfr, _, err = ebitenutil.NewImageFromFile("assets/rcsfr.png")
	rcsbl, _, err = ebitenutil.NewImageFromFile("assets/rcsbl.png")
	rcsbr, _, err = ebitenutil.NewImageFromFile("assets/rcsbr.png")

	if err != nil {
		log.Fatal(err)
	}
}
