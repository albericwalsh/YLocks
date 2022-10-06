package game

import (
	"fmt"

	"RPG/tools"
	"RPG/chapter"
	"RPG/mob"
	"RPG/player"
	"RPG/button"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type Game struct {
	Name    string
	Version string
	Player  player.Player
	Mob map[string]player.Player
}

var (
	ScreenHeight    = 144
	ScreenWidth     = 256
	ScreenResHeight = 144
	ScreenResWidth  = 256

	G = Game{}
)

func NewGame(screen *ebiten.Image, s *tools.Save) {
	if s.CanLoad {
		chapter.Current_Level = s.Chapter
		G.Player.PlayerX = s.PlayerX
		G.Player.PlayerY = s.PlayerY
		G.Player.PV = s.PV
		G.Player.PA = s.PA
		G.Player.PD = s.PD
	}
	if !s.CanLoad {
		G.Player.PlayerX = 16
		G.Player.PlayerY = 144 - 16
	}
	switch chapter.Current_Level {
	case "Int_1_P":
		tools.MainMenuID = "Int_1_P"
	case "Chp_1_0":
		tools.MainMenuID = "Chp_1_0"
		s := tools.LoadSave(&tools.Save{})
		if !s.CanLoad {
			mob.Mob["Card Reader"] = player.Player{PlayerX: 58, PlayerY: 90, Nom: "Card Reader", PV: 15, PA: 3, PD: 1, Beaten: false, Type: "Machine", Image: tools.Card_Reader}
			mob.Mob["Kog'Maw"] = player.Player{PlayerX: 160, PlayerY: 50, Nom: "Kog'Maw", PV: 35, PA: 6, PD: 5, Beaten: false, Type: "Master Boss", Image: tools.PaulImage}
			mob.Mob["Avatar"] = player.Player{PlayerX: 225, PlayerY: 95, Nom: "Avatar", Type: "Event", Image: tools.Avatar}
		} else if s.CanLoad {
			mob.Mob["Card Reader"] = player.Player{PlayerX: 58, PlayerY: 90, Nom: "Card Reader", PV: 15, PA: 3, PD: 1, Beaten: s.MobBeaten["Card Reader"], Type: "Machine", Image: tools.Card_Reader}
			mob.Mob["Kog'Maw"] = player.Player{PlayerX: 160, PlayerY: 50, Nom: "Kog'Maw", PV: 35, PA: 6, PD: 5, Beaten: s.MobBeaten["Kog'Maw"], Type: "Master Boss", Image: tools.PaulImage}
			mob.Mob["Avatar"] = player.Player{PlayerX: 225, PlayerY: 95, Nom: "Avatar", Type: "Event", Image: tools.Avatar}
		}
	case "Chp_2_0":
		tools.MainMenuID = "Chp_2_0"
		s := tools.LoadSave(&tools.Save{})
		if !s.CanLoad {
			mob.Mob["Vitaly"] = player.Player{PlayerX: 120, PlayerY: 50, Nom: "Vitaly", PV: 50, PA: 8, PD: 5, Beaten: false, Type: "Master Boss", Image: tools.Vitaly}
		} else if s.CanLoad {
			mob.Mob["Vitaly"] = player.Player{PlayerX: 120, PlayerY: 50, Nom: "Vitaly", PV: 50, PA: 8, PD: 5, Beaten: s.MobBeaten["Vitaly"], Type: "Master Boss", Image: tools.Vitaly}
		}
	case "Chp_3_0":
		//start 3rd chapter (Classes)
	case "Chp_4_0":
		//start Final chapter (Final Dungeon)
	}
}

func SetPlayer(screen *ebiten.Image, g *Game) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.Player.PlayerX)-15, float64(g.Player.PlayerY)-15)
	screen.DrawImage(tools.PlayerImage, op)
}

// MainMenu is the main menu of the game
func (g *Game) MainMenu() {
	// set the game name
	g.Name = "YLock's"
	// set the game version
	g.Version = "0.2.13"
	// run the game
	ebiten.SetWindowIcon(tools.IconImage)
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle(g.Name + " " + g.Version)
	ebiten.SetWindowResizable(true)
	ebiten.MaximizeWindow()
	ebiten.SetMaxTPS(60)
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}

// Update updates the game state.
func (g *Game) Update(screen *ebiten.Image) error {
	//fmt.Println(tools.MainMenuID)
	tools.SetMousePosition()
	//CheckButtonID(tools.MainMenuID, screen, tools.Save{})
	g.Draw(screen)
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
func (g *Game) Draw(screen *ebiten.Image) {
	g.Player.PA = 5
	g.Player.PD = 15
	g.Player.PV = 25
	button.CheckButtonID(tools.MainMenuID, screen, &tools.Save{})
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (ScreenWidth, ScreenHeight int) {
	return ScreenResWidth, ScreenResHeight
}