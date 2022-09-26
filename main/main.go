package main

import (
	"RPG"
	"fmt"

	"github.com/hajimehoshi/ebiten"
)

// "github.com/hajimehoshi/ebiten/ebitenutil"

type Game struct {
	Name            string
	Version         string
	PlayerX         int
	PlayerY         int
}

var (
	ScreenHeight = 144
	ScreenWidth = 256
	ScreenResHeight = 144
	ScreenResWidth = 256
)

// MainMenu is the main menu of the game
func (g *Game) MainMenu() {
	// set the game name
	g.Name = "YLock's"
	// set the game version
	g.Version = "0.0.1"
	
	// run the game
	ebiten.SetWindowSize(ScreenHeight, ScreenWidth)
	ebiten.SetWindowTitle(g.Name + " " + g.Version)
	ebiten.SetWindowResizable(true)
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}

// Update updates the game state.
func (g *Game) Update(screen *ebiten.Image) error {
	//update screen
	g.Draw(screen)
	return nil
}

// Draw draws the game screen.
func (g *Game) Draw(screen *ebiten.Image) {
	// draw the background and set the position to 0:0
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	// background size
	op.GeoM.Scale(1, 1)
	screen.DrawImage(RPG.BackgroundImage, op)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (ScreenWidth, ScreenHeight int) {
	return ScreenResWidth, ScreenResHeight
}

func main() {
	RPG.Textures_init()
	game := Game{}
	game.MainMenu()

	fmt.Println(RPG.Path)
	if RPG.CheckSave() {
		fmt.Println("Save file exists")
	} else {
		fmt.Println("Save file does not exist")
		fmt.Println("Creating save file...")
		RPG.CreateSaveFile()
	}
}
