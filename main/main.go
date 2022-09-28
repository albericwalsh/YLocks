package main

import (
	"RPG"
	"fmt"
	"image/color"
	"os"
	// "time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// "github.com/hajimehoshi/ebiten/ebitenutil"

type Game struct {
	Name    string
	Version string
	PlayerX int
	PlayerY int
	screen  *ebiten.Image
}

var (
	ScreenHeight    = 144
	ScreenWidth     = 256
	ScreenResHeight = 144
	ScreenResWidth  = 256
	paragraph1      = "BIENVENUE ! \nAujourd'hui c'est votre rentrée au sein de \nl'école Ynov sur le campus parisien."
	paragraph2      = "Nous sommes ravis de vous accueillir pour \nles cinq prochaines années de vos études \net nous espérons du fond du coeur que vous \nvous épanouirez. "
	paragraph3      = "Nous allons vous remettre votre badge \nd'accès et nous vous ferons visiter \nle campus. Pour le bon déroulement \nde cette journée vous allez être \nrépartis par filières."
	paragraph4      = "Je vous invite donc à entrer et \nà attendre que vos mentors viennent vous \nchercher. \nBonne journée !"
	Next            = false
)

func print(s string, screen *ebiten.Image) {
	// draw a white text
	ebitenutil.DebugPrint(screen, s)
}

func NewGame(screen *ebiten.Image, s RPG.Save) {
	switch s.Chapter {
	case 0:
		RPG.MainMenuID = "Int_1_P"
	case 1:
		RPG.MainMenuID = "Chp_1_0"
	case 2:
		//start 2nd chapter (Souk)
	case 3:
		//start 3rd chapter (Classes)
	case 4:
		//start Final chapter (Final Dungeon)
	}

}

func CheckButtonID(ID string, screen *ebiten.Image, s RPG.Save) {
	switch ID {
	case "":
		// draw the background and set the position to 0:0
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(0, 0)
		// background size
		op.GeoM.Scale(1, 1)
		screen.DrawImage(RPG.BackgroundImage, op)
		// draw the buttons at the center of the screen
		RPG.Button(screen, false, RPG.Center(RPG.ButtonImage, ScreenWidth), 10, "New Game", "New_Game")
		RPG.Button(screen, !RPG.CanLoad(RPG.Save{}), RPG.Center(RPG.ButtonImage, ScreenWidth), 42, "Load Game", "Load_Game")
		RPG.Button(screen, false, RPG.Center(RPG.ButtonImage, ScreenWidth), 74, "Settings", "Settings")
		RPG.Button(screen, false, RPG.Center(RPG.ButtonImage, ScreenWidth), 106, "Quit", "Quit")
	case "New_Game":
		//RPG.CreateSave(RPG.Save{})
		NewGame(screen, RPG.Save{})
		//RPG.SetCanLoad(RPG.Save{}, true)
	case "Load_Game":
		//RPG.LoadSave(RPG.Save{})
		NewGame(screen, RPG.Save{})
		fmt.Print(RPG.Save{})
	case "Settings":
		fmt.Println("Settings")
	case "Quit":
		os.Exit(0)
	case "Int_1_P":
		screen.Fill(color.RGBA{0, 0, 0, 0})
		print(paragraph1, screen)
		RPG.Button(screen, false, 256-66, 144-18, "Next", "Int_2_P")
	case "Int_2_P":
		screen.Fill(color.RGBA{0, 0, 0, 0})
		print(paragraph2, screen)
		RPG.Button(screen, false, 256-66, 144-18, "Next", "Int_3_P")
		RPG.Button(screen, false, 2, 144-18, "Previous", "Int_1_P")
	case "Int_3_P":
		screen.Fill(color.RGBA{0, 0, 0, 0})
		print(paragraph3, screen)
		RPG.Button(screen, false, 256-66, 144-18, "Next", "Int_4_P")
		RPG.Button(screen, false, 2, 144-18, "Previous", "Int_2_P")
	case "Int_4_P":
		screen.Fill(color.RGBA{0, 0, 0, 0})
		print(paragraph4, screen)
		RPG.Button(screen, false, 256-66, 144-18, "Next", "Int_Next_Chapter")
		RPG.Button(screen, false, 2, 144-18, "Previous", "Int_3_P")
	case "Int_Next_Chapter":
		s.Chapter += 1
		NewGame(screen, s)
	}
}

// MainMenu is the main menu of the game
func (g *Game) MainMenu() {
	// set the game name
	g.Name = "YLock's"
	// set the game version
	g.Version = "0.0.1"

	// run the game
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle(g.Name + " " + g.Version)
	ebiten.SetWindowResizable(true)
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}

// Update updates the game state.
func (g *Game) Update(screen *ebiten.Image) error {
	fmt.Println(RPG.MainMenuID)
	RPG.SetMousePosition()
	//CheckButtonID(RPG.MainMenuID, screen, RPG.Save{})
	//update screen
	g.Draw(screen)
	// set frame rate
	ebiten.SetMaxTPS(60)
	return nil
}

// Draw draws the game screen.
func (g *Game) Draw(screen *ebiten.Image) {
	CheckButtonID(RPG.MainMenuID, screen, RPG.Save{})
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
}
