package game

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	playerImage *ebiten.Image
	rcsfl       *ebiten.Image
	rcsfr       *ebiten.Image
	rcsbl       *ebiten.Image
	rcsbr       *ebiten.Image
)

func init() {
	playerImage, _, err = ebitenutil.NewImageFromFile("assets/player.png", ebiten.FilterDefault)
	rcsfl, _, err = ebitenutil.NewImageFromFile("assets/rcsfl.png", ebiten.FilterDefault)
	rcsfr, _, err = ebitenutil.NewImageFromFile("assets/rcsfr.png", ebiten.FilterDefault)
	rcsbl, _, err = ebitenutil.NewImageFromFile("assets/rcsbl.png", ebiten.FilterDefault)
	rcsbr, _, err = ebitenutil.NewImageFromFile("assets/rcsbr.png", ebiten.FilterDefault)

	if err != nil {
		log.Fatal(err)
	}
}
