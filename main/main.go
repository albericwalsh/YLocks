package main

import (
	"RPG/game"
	"RPG/tools"
)

func main() {
	tools.Textures_init()
	game := game.Game{}
	game.MainMenu()
}