/*
author: cyberpunkprogrammer
date: 10-30-2020
tutorial: https://medium.com/@chrisandrews_76960/2d-game-development-in-golang-part-3-a296aedea77b
*/

package main

import "github.com/hajimehoshi/ebiten"

var (
	err        error
	background *ebiten.Image
	spaceShip  *ebiten.Image
)

type player struct {
	image      *ebiten.Image
	xPos, yPos float64
	speed      float64
}

func main() {

}
