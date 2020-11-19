package main

import "github.com/hajimehoshi/ebiten/v2"

var objects []Object

type Object interface {
	Update()
	Draw(*ebiten.Image, *ebiten.DrawImageOptions, *Game)
}
