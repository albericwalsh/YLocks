package main

import (
	"fmt"


	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"

	"RPG/chapter"
	"RPG/button"
	"RPG/fight"
	"RPG/game"
	"RPG/mob"
	"RPG/player"
	"RPG/tools"
)

var (
	ScreenHeight    = 144
	ScreenWidth     = 256
	ScreenResHeight = 144
	ScreenResWidth  = 256
)

func NewGame(screen *ebiten.Image, s *tools.Save,g *game.Game) {
	game.NewGame(screen, s)
}

func Fight(screen *ebiten.Image, v string, m map[string]player.Player, PV *int, s *tools.Save, g *game.Game) {
	fight.Fight(screen, v, m, PV, s)
}


// MainMenu is the main menu of the game
func MainMenu(g *game.Game) {
	game := new(game.Game)
	// set the game name
	game.Name = "YLock's"
	// set the game version
	game.Version = "0.2.13"
	// run the game
	ebiten.SetWindowIcon(tools.IconImage)
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle(game.Name + " " + game.Version)
	ebiten.SetWindowResizable(true)
	ebiten.MaximizeWindow()
	ebiten.SetMaxTPS(60)
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}

// Update updates the game state.
func Update(screen *ebiten.Image, g *game.Game) error {
	//fmt.Println(tools.MainMenuID)
	tools.SetMousePosition()
	//CheckButtonID(tools.MainMenuID, screen, tools.Save{})
	Draw(screen,&game.Game{})
	// set frame rate
	if chapter.CanMove {
		if ebiten.IsKeyPressed(ebiten.KeyUp) {
			if g.Player.PlayerY > 16 {
				if tools.MainMenuID == "Chp_1_0" {
					if g.Player.PlayerY < 88 && g.Player.PlayerX < 73  {
						fmt.Print()
					} else if (g.Player.PlayerY <= 88 && (g.Player.PlayerX > 73 && g.Player.PlayerX < 120)) || (g.Player.PlayerY <= 88 && (g.Player.PlayerX > 150 && g.Player.PlayerX < 256)) {
						fmt.Print()
					} else {
						g.Player.PlayerY -= 1
					}
				}
			}
		}
		if ebiten.IsKeyPressed(ebiten.KeyDown) {
			if g.Player.PlayerY < 144-16 {
				g.Player.PlayerY += 1
			}
		}
		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			if g.Player.PlayerX > 16 {
				if tools.MainMenuID == "Chp_1_0" {
					if (g.Player.PlayerY < 88 && g.Player.PlayerX < 74) || (g.Player.PlayerY <= 88 && (g.Player.PlayerX > 73 && g.Player.PlayerX < 120)) {
						fmt.Print()
					} else {
						g.Player.PlayerX -= 1
					}
				}
			}
		}
		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			if g.Player.PlayerX < 256-16 {
				if tools.MainMenuID == "Chp_1_0" {
					if !mob.Mob["Card Reader"].Beaten {
						if g.Player.PlayerX < 74-16 {
							g.Player.PlayerX += 1
						}
					} else if (g.Player.PlayerY <= 88 && (g.Player.PlayerX > 135 && g.Player.PlayerX < 205)) || (g.Player.PlayerX >= 205 ){
						fmt.Print()
					} else {
						g.Player.PlayerX += 1
					}
				}
			}
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			tools.MainMenuID = "Pause"
			chapter.CanMove = false
			button.Pause = true
			button.IsWait = true
			button.WaitDuration = 10
		}
		if button.IsWait {
			button.Wait(button.WaitDuration)
		}
	}
	ebiten.SetMaxTPS(60)
	return nil
}

// Draw draws the game screen.
func Draw(screen *ebiten.Image, g *game.Game) {
	g.Player.PA = 5
	g.Player.PD = 15
	g.Player.PV = 25
	button.CheckButtonID(tools.MainMenuID, screen, &tools.Save{})
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func Layout(outsideWidth, outsideHeight int) (ScreenWidth, ScreenHeight int) {
	return ScreenResWidth, ScreenResHeight
}

func main() {
	tools.Textures_init()
	game := game.Game{}
	MainMenu(&game)
}