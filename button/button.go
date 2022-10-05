package button

import (
	"image/color"
	"os"
	"fmt"

	"RPG/tools"
	"RPG/chapter"
	"RPG/draw"
	"RPG/fight"
	"RPG/game"
	"RPG/mob"
	"RPG/player"
	"RPG/text"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

var (
	ScreenHeight    = 144
	ScreenWidth     = 256
	ScreenResHeight = 144
	ScreenResWidth  = 256

	IsWait          = false
	WaitDuration    = 0
	WaitIndex       = 0

	Pause           = false
)

func CheckButtonID(ID string, screen *ebiten.Image, s *tools.Save) {
	switch ID {
	case "":
		// draw the background and set the position to 0:0
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(0, 0)
		// background size
		op.GeoM.Scale(1, 1)
		screen.DrawImage(tools.BackgroundImage, op)
		// draw the buttons at the center of the screen
		tools.Button(screen, false, tools.Center(tools.ButtonImage, ScreenWidth), 10, "New Game", "New_Game")
		tools.Button(screen, !tools.CanLoad(tools.Save{}), tools.Center(tools.ButtonImage, ScreenWidth), 42, "Load Game", "Load_Game")
		tools.Button(screen, false, tools.Center(tools.ButtonImage, ScreenWidth), 74, "Settings", "Settings")
		tools.Button(screen, false, tools.Center(tools.ButtonImage, ScreenWidth), 106, "Quit", "Quit")
	case "New_Game":
		s.Chapter = "Int_1_P"
		chapter.Current_Level = s.Chapter
		s.CanLoad = false
		game.NewGame(screen, &tools.Save{})
		tools.CreateSave(&tools.Save{CanLoad: true, Chapter: chapter.Current_Level, PlayerX: game.G.Player.PlayerX, PlayerY: game.G.Player.PlayerY, PV: player.PlayerPV, PA: game.G.Player.PA, PD: game.G.Player.PD})
	case "Load_Game":
		test := tools.LoadSave(&tools.Save{})
		game.NewGame(screen, &test)
	case "Settings":
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(0, 0)
		screen.DrawImage(tools.BackgroundImage, op)
		ebitenutil.DrawRect(screen, 0, 0, 256, 144, color.RGBA{0, 0, 0, 170})
		ebitenutil.DebugPrintAt(screen, "Settings", 10, 10)
		tools.UpdateSave(&tools.Save{CanLoad: true, Chapter: chapter.Current_Level, PlayerX: game.G.Player.PlayerX, PlayerY: game.G.Player.PlayerY, PV: player.PlayerPV, PA: game.G.Player.PA, PD: game.G.Player.PD, MobX: mob.MobX, MobY: mob.MobY, MobPV: mob.MobPV, MobPA: mob.MobPA, MobPD: mob.MobPD, MobBeaten: mob.MobBeaten, MobImage: mob.MobImage})
		tools.Button(screen, false, ScreenResWidth-70, ScreenHeight-32, "Ok", "")
		// Settings
		// Fullscreen
		ebitenutil.DebugPrintAt(screen, "Fullscreen", 22, 30)
		if !ebiten.IsFullscreen() {
			ebitenutil.DrawRect(screen, 10, 34, 9, 9, color.RGBA{0, 0, 0, 170})
			if tools.MouseX > 10 && tools.MouseX < 19 && tools.MouseY > 34 && tools.MouseY < 43 && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
				ebitenutil.DrawRect(screen, 10, 34, 9, 9, color.RGBA{255, 255, 255, 170})
				ebiten.SetFullscreen(true)
			}
		} else {
			ebitenutil.DrawRect(screen, 10, 34, 9, 9, color.RGBA{255, 255, 255, 170})
			if tools.MouseX > 10 && tools.MouseX < 19 && tools.MouseY > 34 && tools.MouseY < 43 && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
				ebitenutil.DrawRect(screen, 10, 34, 9, 9, color.RGBA{0, 0, 0, 170})
				ebiten.SetFullscreen(false)
			}
		}
		//
	case "Quit":
		mob.SetMobVariable(mob.Mob, mob.MobName)
		tools.UpdateSave(&tools.Save{CanLoad: true, Chapter: chapter.Current_Level, PlayerX: game.G.Player.PlayerX, PlayerY: game.G.Player.PlayerY, PV: player.PlayerPV, PA: game.G.Player.PA, PD: game.G.Player.PD, MobX: mob.MobX, MobY: mob.MobY, MobPV: mob.MobPV, MobPA: mob.MobPA, MobPD: mob.MobPD, MobBeaten: mob.MobBeaten, MobImage: mob.MobImage})
		os.Exit(0)
	case "Int_1_P":
		op := &ebiten.DrawImageOptions{}
		screen.DrawImage(tools.BackgroundImage, op)
		ebitenutil.DrawRect(screen, 0, 0, 256, 144, color.RGBA{0, 0, 0, 170})
		text.Print(text.Paragraph1, screen)
		tools.Button(screen, false, 256-66, 144-18, "Next", "Int_2_P")
	case "Int_2_P":
		op := &ebiten.DrawImageOptions{}
		screen.DrawImage(tools.BackgroundImage, op)
		ebitenutil.DrawRect(screen, 0, 0, 256, 144, color.RGBA{0, 0, 0, 170})
		text.Print(text.Paragraph2, screen)
		tools.Button(screen, false, 256-66, 144-18, "Next", "Int_3_P")
		tools.Button(screen, false, 2, 144-18, "Previous", "Int_1_P")
	case "Int_3_P":
		op := &ebiten.DrawImageOptions{}
		screen.DrawImage(tools.BackgroundImage, op)
		ebitenutil.DrawRect(screen, 0, 0, 256, 144, color.RGBA{0, 0, 0, 170})
		text.Print(text.Paragraph3, screen)
		tools.Button(screen, false, 256-66, 144-18, "Next", "Int_4_P")
		tools.Button(screen, false, 2, 144-18, "Previous", "Int_2_P")
	case "Int_4_P":
		op := &ebiten.DrawImageOptions{}
		screen.DrawImage(tools.BackgroundImage, op)
		ebitenutil.DrawRect(screen, 0, 0, 256, 144, color.RGBA{0, 0, 0, 170})
		text.Print(text.Paragraph4, screen)
		tools.Button(screen, false, 256-66, 144-18, "Next", "Int_Next_Chapter")
		tools.Button(screen, false, 2, 144-18, "Previous", "Int_3_P")
	case "Int_Next_Chapter":
		s.Chapter = "Chp_1_0"
		chapter.Current_Level = "Chp_1_0"
		game.NewGame(screen, s)
	case "Chp_1_0":
		chapter.CanMove = true
		chapter.Current_Level = "Chp_1_0"
		op := &ebiten.DrawImageOptions{}
		screen.DrawImage(tools.Background_Ch1, op)
		draw.DrawMob(mob.Mob, screen)
		game.SetPlayer(screen, &game.Game{})
		// fmt.Printf("PlayerX: %v, PlayerY: %v \n", g.Player.PlayerX, g.Player.PlayerY)
		for _, v := range mob.Mob {
			v.MaxX = v.PlayerX + v.Image.Bounds().Dx()
			v.MaxY = v.PlayerY + v.Image.Bounds().Dy()
			if (game.G.Player.PlayerX+25 >= v.PlayerX && game.G.Player.PlayerX+25 <= v.MaxX) && (game.G.Player.PlayerY+25 >= v.PlayerY && game.G.Player.PlayerY+25 <= v.MaxY) {
				if !v.Beaten {
					if v.Type != "Event" {
						//fight
						ebitenutil.DrawRect(screen, 2, float64(ScreenHeight)-20, float64(ScreenResWidth), float64(ScreenHeight), color.RGBA{0, 0, 0, 170})
						ebitenutil.DebugPrintAt(screen, "For Fight "+v.Nom+", Press [ENTER]", 2, ScreenHeight-20)
						//tools.Fight(screen, g.Player, v)
						if ebiten.IsKeyPressed(ebiten.KeyEnter) || fight.InFight {
							fight.InFight = true
							tools.MainMenuID = "Init_Fight"
							mob.MobName = v.Nom
						}
					} else {
						ebitenutil.DrawRect(screen, 2, float64(ScreenHeight)-20, float64(ScreenResWidth), float64(ScreenHeight), color.RGBA{0, 0, 0, 170})
						ebitenutil.DebugPrintAt(screen, "C'est de toute beauté", 2, ScreenHeight-20)
					}
				}
			}
		}
	case "Chp_2_0":
		chapter.CanMove = true
		chapter.Current_Level = "Chp_2_0"
		op := &ebiten.DrawImageOptions{}
		screen.DrawImage(tools.Background_Ch2, op)
		draw.DrawMob(mob.Mob, screen)
		game.SetPlayer(screen, &game.Game{})
		// fmt.Printf("PlayerX: %v, PlayerY: %v \n", g.Player.PlayerX, g.Player.PlayerY)
		for _, v := range mob.Mob {
			v.MaxX = v.PlayerX + v.Image.Bounds().Dx()
			v.MaxY = v.PlayerY + v.Image.Bounds().Dy()
			if (game.G.Player.PlayerX+25 >= v.PlayerX && game.G.Player.PlayerX+25 <= v.MaxX) && (game.G.Player.PlayerY+25 >= v.PlayerY && game.G.Player.PlayerY+25 <= v.MaxY) {
				if !v.Beaten {
					//fight
					ebitenutil.DrawRect(screen, 2, float64(ScreenHeight)-20, float64(ScreenResWidth), float64(ScreenHeight), color.RGBA{0, 0, 0, 170})
					ebitenutil.DebugPrintAt(screen, "For Fight "+v.Nom+", Press [ENTER]", 2, ScreenHeight-20)
					//tools.Fight(screen, g.Player, v)
					if ebiten.IsKeyPressed(ebiten.KeyEnter) || fight.InFight {
						fight.InFight = true
						tools.MainMenuID = "Init_Fight"
						mob.MobName = v.Nom
					}
				}
			}
		}
	case "Init_Fight":
		fight.Fight(screen, mob.MobName, mob.Mob, &mob.Pv, &tools.Save{})
	case "Fight":
		fight.Fight(screen, mob.MobName, mob.Mob, &mob.Pv, &tools.Save{})
	case "Pause":
		op := &ebiten.DrawImageOptions{}
		screen.DrawImage(tools.BackgroundImage, op)
		ebitenutil.DrawRect(screen, 0, 0, 256, 144, color.RGBA{0, 0, 0, 170})
		ebitenutil.DrawRect(screen, 0, 0, 0, 144, color.RGBA{0, 0, 0, 170})
		mob.SetMobVariable(mob.Mob, mob.MobName)
		tools.UpdateSave(&tools.Save{CanLoad: true, Chapter: chapter.Current_Level, PlayerX: game.G.Player.PlayerX, PlayerY: game.G.Player.PlayerY, PV: player.PlayerPV, PA: game.G.Player.PA, PD: game.G.Player.PD, MobX: mob.MobX, MobY: mob.MobY, MobPV: mob.MobPV, MobPA: mob.MobPA, MobPD: mob.MobPD, MobBeaten: mob.MobBeaten, MobImage: mob.MobImage})
		tools.Button(screen, false, 256-66, 144-18, "Resume", chapter.Current_Level)
		tools.Button(screen, false, 256-66, 2, "Main Menu", "")
		tools.Button(screen, false, 2, 144-18, "Settings", "Settings")
		tools.Button(screen, false, 2, 2, "Quit", "Quit")
		// fmt.Printf("Pause: %v\n", Pause)
		if inpututil.IsKeyJustReleased(ebiten.KeyEscape) && Pause && !IsWait {
			tools.MainMenuID = chapter.Current_Level
			Pause = false
			IsWait = true
			WaitDuration = 10
		}
	}
}

func Wait(time int) {
	if WaitIndex < time {
		j := 10000000000000000 + 10000000000000000
		fmt.Print(j)
		WaitIndex++
	} else if WaitIndex == time {
		WaitIndex = 0
		IsWait = false
	}
}