package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Player that the user controls
var (
	Player object
)

// CreatePlayer for user to control at coordinates
func CreatePlayer(x float64, y float64) {
	Player.image, _, err = ebitenutil.NewImageFromFile("assets/player.png", ebiten.FilterDefault)
	Player.xPos = x
	Player.yPos = y
	Player.rPos = 0
	Player.xSpd = 0
	Player.ySpd = 0
	Player.rSpd = 0

	// Add player to list of objects
	All = append(All, Player)
}
