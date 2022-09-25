package main

import (
	"fmt"
	"RPG"
)

func main() {
	//game := RPG.Game{}
	//game.Master()

	fmt.Println(RPG.Path)
	if RPG.CheckSave() {
		fmt.Println("Save file exists")
	} else {
		fmt.Println("Save file does not exist")
		fmt.Println("Creating save file...")
		RPG.CreateSaveFile()
	}
}