package fight

import (
	"strings"
	"image/color"
	"fmt"
	
	"RPG/chapter"
	"RPG/player"
	"RPG/mob"
	"RPG/game"
	"RPG/tools"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	ScreenHeight    = 144
	ScreenWidth     = 256
	ScreenResHeight = 144
	ScreenResWidth  = 256

	InFight         = false
	YourTurn        = true
)

func Fight(screen *ebiten.Image, v string, m map[string]player.Player, PV *int, s *tools.Save) {
	if tools.MainMenuID == "Init_Fight" {
		*PV = m[v].PV
		player.PlayerPV = game.G.Player.PV
		tools.MainMenuID = "Fight"
	} else if tools.MainMenuID == "Fight" {
		if *PV > 0 {
			chapter.CanMove = false
			// draw the Background
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(0, 0)
			screen.DrawImage(tools.BackgroundImage, op)
			// draw the Player at the left bottom of the screen
			op = &ebiten.DrawImageOptions{}
			op.GeoM.Scale(2, 2)
			op.GeoM.Translate(10, float64(ScreenResHeight)-74)
			screen.DrawImage(tools.PlayerImage, op)
			ebitenutil.DrawRect(screen, 9, float64(ScreenResHeight)-79, 100, 5, color.RGBA{0, 0, 0, 170})
			ebitenutil.DrawLine(screen, 10, float64(ScreenResHeight)-78, 100, float64(ScreenResHeight)-78, color.RGBA{200, 200, 200, 170})
			ebitenutil.DrawLine(screen, 10, float64(ScreenResHeight)-78, (float64(player.PlayerPV)*100)/float64(game.G.Player.PV), float64(ScreenResHeight)-78, color.RGBA{0, 255, 0, 170})
			ebitenutil.DrawLine(screen, 10, float64(ScreenResHeight)-76, 100, float64(ScreenResHeight)-76, color.RGBA{200, 200, 200, 170})
			ebitenutil.DrawLine(screen, 10, float64(ScreenResHeight)-76, float64(game.G.Player.PA)*10, float64(ScreenResHeight)-76, color.RGBA{255, 0, 0, 170})
			// draw the Mob at the Right top of the screen
			op = &ebiten.DrawImageOptions{}
			op.GeoM.Scale(2, 2)
			op.GeoM.Translate(float64(ScreenResWidth)-74, 10)
			screen.DrawImage(m[v].Image, op)
			// draw the name of the mob 2:2
			ebitenutil.DrawRect(screen, 2, 2, float64(ScreenResWidth)-32, 16, color.RGBA{0, 0, 0, 170})
			ebitenutil.DebugPrintAt(screen, m[v].Nom, 2, 2)
			ebitenutil.DrawRect(screen, 2, 19, 100, 5, color.RGBA{0, 0, 0, 170})
			ebitenutil.DrawLine(screen, 3, 20, 100, 20, color.RGBA{200, 200, 200, 170})
			ebitenutil.DrawLine(screen, 3, 20, (float64(*PV)*100)/float64(m[v].PV), 20, color.RGBA{0, 255, 0, 170})
			ebitenutil.DrawLine(screen, 3, 22, 100, 22, color.RGBA{200, 200, 200, 170})
			ebitenutil.DrawLine(screen, 3, 22, float64(m[v].PA)*10, 22, color.RGBA{255, 0, 0, 170})
			// draw button
			tools.Button(screen, !YourTurn, ScreenResWidth-180, ScreenHeight-70, "Attack", "Attack")
			tools.Button(screen, !YourTurn, ScreenResWidth-180, ScreenHeight-50, "Regen", "Regen")
			tools.Button(screen, !YourTurn, ScreenResWidth-180, ScreenHeight-30, "Run", "Run")
			//---------------------------------------------------------------------------------------------------------
			fmt.Println(player.PlayerPV, "pv actuel")
			fmt.Println((float64(player.PlayerPV)*100)/(float64(game.G.Player.PV)), "pv actuel en %")
			switch tools.MainMenuID {
			case "Attack":
				fmt.Println((float64(player.PlayerPV)*100)/(float64(game.G.Player.PV)), "pv%")
				miss := player.GetMiss(0, 2)
				critical := player.Getcritical(0, 3)
				// fmt.Println("debut ", *PV)
				if miss != 0 {
					*PV = *PV
					// fmt.Println("miss atk: ", *PV)
				} else {
					if critical == 0 || critical == 1 {
						*PV -= game.G.Player.PA
						// fmt.Println("pas crit atk:", *PV)
					} else {
						*PV -= game.G.Player.PA * critical
						// fmt.Println("crit atk: ", *PV)
					}
				}
				// fmt.Println("fin ", *PV)
				YourTurn = false
				tools.MainMenuID = "Fight"
			case "Regen":
				if (player.PlayerPV+game.G.Player.PD) > 25 {
					player.PlayerPV = game.G.Player.PV
					// tools.PrintonTime(screen, "You are full HP",10,10, 0)
				} else if (player.PlayerPV+game.G.Player.PD) <= game.G.Player.PV {
					player.PlayerPV = player.PlayerPV + game.G.Player.PD
				}
				YourTurn = false
				tools.MainMenuID = "Fight"
			case "Run":
				InFight = false
				tools.MainMenuID = "Chp_1_0"
				chapter.CanMove = true
			}
			if !YourTurn {
				player.PlayerPV -= m[v].PA
				tools.MainMenuID = "Fight"
				// tools.PrintonTime(screen, m[v].Nom+" attack you, damage "+string(m[v].PA), 10, 10, 2)
				YourTurn = true
			}
		}
		if *PV <= 0 {
			// draw the Background
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(0, 0)
			screen.DrawImage(tools.BackgroundImage, op)
			ebitenutil.DrawRect(screen, 0, 0, 256, 144, color.RGBA{0, 0, 0, 170})
			ebitenutil.DebugPrintAt(screen, "You've beaten "+v, 10, 10)
			tools.Button(screen, false, ScreenResWidth-70, ScreenHeight-32, "Continue", "Beaten")
			if tools.MainMenuID == "Beaten" {
				InFight = false
				chapter.CanMove = true
				temp := mob.Mob[mob.MobName]
				temp.Beaten = true
				mob.Mob[mob.MobName] = temp
				game.G.Mob = mob.Mob
				if mob.Mob[mob.MobName].Type == "Master Boss" {
					CheckChapter := strings.Split(chapter.Current_Level, "_")
					for _, v := range CheckChapter {
						if v == "1" {
							CheckChapter[1] = "2"
						}
						if v == "2" {
							CheckChapter[1] = "3"
						}
					}
					chapter.NextChapter = strings.Join(CheckChapter[:], "_")
					chapter.Current_Level = chapter.NextChapter
					game.NewGame(screen, s)
				} else {
					chapter.NextChapter = chapter.Current_Level
				}
				tools.MainMenuID = chapter.NextChapter 
				mob.SetMobVariable(mob.Mob, mob.MobName)
				tools.UpdateSave(&tools.Save{CanLoad: true, Chapter: chapter.Current_Level, PlayerX: game.G.Player.PlayerX, PlayerY: game.G.Player.PlayerY, PV: player.PlayerPV, PA: game.G.Player.PA, PD: game.G.Player.PD, MobX: mob.MobX, MobY: mob.MobY, MobPV: mob.MobPV, MobPA: mob.MobPA, MobPD: mob.MobPD, MobBeaten: mob.MobBeaten, MobImage: mob.MobImage})
				// fmt.Println(Mob)
			}
		} else if player.PlayerPV < 0 {
			// draw the Background
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(0, 0)
			screen.DrawImage(tools.BackgroundImage, op)
			ebitenutil.DrawRect(screen, 0, 0, 256, 144, color.RGBA{0, 0, 0, 170})
			ebitenutil.DebugPrintAt(screen, "You are beaten by "+v, 10, 10)
			tools.DeleteSave()
			tools.Button(screen, false, ScreenResWidth-70, ScreenHeight-32, "Return Main Menu", "")
		}
	}
}