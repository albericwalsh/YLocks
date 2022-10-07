package main

import (
	"RPG"
	"fmt"
	"image/color"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"

	"RPG/chapter"
	"RPG/draw"
	"RPG/player"
)

type Game struct {
	Name    string
	Version string
	Player  player.Player
	Mob     map[string]player.Player
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
	Gaga            = true
	InFight         = false
	YourTurn        = true
	Pause           = false
	MobName         = ""
	Pv              = 0
	PlayerPV        = 0
	MobX            = map[string]int{}
	MobY            = map[string]int{}
	MobPV           = map[string]int{}
	MobPA           = map[string]int{}
	MobPD           = map[string]int{}
	MobBeaten       = map[string]bool{}
	MobImage        = map[string]*ebiten.Image{}
	IsWait          = false
	WaitDuration    time.Duration
	WaitTemp        time.Time

	Mob = map[string]player.Player{}
)

func print(s string, screen *ebiten.Image) {
	// draw a white text
	ebitenutil.DebugPrint(screen, s)
}

func (g *Game) NewGame(screen *ebiten.Image, s *RPG.Save) {
	if s.CanLoad {
		chapter.Current_Level = s.Chapter
		g.Player.PlayerX = s.PlayerX
		g.Player.PlayerY = s.PlayerY
		g.Player.PV = s.PV
		g.Player.PA = s.PA
		g.Player.PD = s.PD
	}
	if !s.CanLoad {
		g.Player.PlayerX = 16
		g.Player.PlayerY = 144 - 16
	}
	switch chapter.Current_Level {
	case "Int_1_P":
		RPG.MainMenuID = "Int_1_P"
	case "Chp_1_0":
		RPG.MainMenuID = "Chp_1_0"
		s := RPG.LoadSave(&RPG.Save{})
		if !s.CanLoad {
			Mob["Card Reader"] = player.Player{PlayerX: 58, PlayerY: 90, Nom: "Card Reader", PV: 15, PA: 3, PD: 1, Beaten: false, Type: "Machine", Image: RPG.Card_Reader}
			Mob["Kog'Maw"] = player.Player{PlayerX: 160, PlayerY: 50, Nom: "Kog'Maw", PV: 35, PA: 6, PD: 5, Beaten: false, Type: "Master Boss", Image: RPG.PaulImage}
			Mob["Avatar"] = player.Player{PlayerX: 225, PlayerY: 95, Nom: "Avatar", Type: "Event", Image: RPG.Avatar}
		} else if s.CanLoad {
			Mob["Card Reader"] = player.Player{PlayerX: 58, PlayerY: 90, Nom: "Card Reader", PV: 15, PA: 3, PD: 1, Beaten: s.MobBeaten["Card Reader"], Type: "Machine", Image: RPG.Card_Reader}
			Mob["Kog'Maw"] = player.Player{PlayerX: 160, PlayerY: 50, Nom: "Kog'Maw", PV: 35, PA: 6, PD: 5, Beaten: s.MobBeaten["Kog'Maw"], Type: "Master Boss", Image: RPG.PaulImage}
			Mob["Avatar"] = player.Player{PlayerX: 225, PlayerY: 95, Nom: "Avatar", Type: "Event", Image: RPG.Avatar}
		}
	case "Chp_2_0":
		RPG.MainMenuID = "Chp_2_0"
		s := RPG.LoadSave(&RPG.Save{})
		if !s.CanLoad {
			Mob["Vitaly"] = player.Player{PlayerX: 120, PlayerY: 50, Nom: "Vitaly", PV: 50, PA: 8, PD: 5, Beaten: false, Type: "Master Boss", Image: RPG.Vitaly}
		} else if s.CanLoad {
			Mob["Vitaly"] = player.Player{PlayerX: 120, PlayerY: 50, Nom: "Vitaly", PV: 50, PA: 8, PD: 5, Beaten: s.MobBeaten["Vitaly"], Type: "Master Boss", Image: RPG.Vitaly}
		}
	case "Chp_3_0":
		//start 3rd chapter (Classes)
	case "Chp_4_0":
		//start Final chapter (Final Dungeon)
	}

}

func SetMobVariable(m map[string]player.Player, name string) {
	MobX[name] = m[name].PlayerX
	MobY[name] = m[name].PlayerY
	MobPV[name] = m[name].PV
	MobPA[name] = m[name].PA
	MobPD[name] = m[name].PD
	MobBeaten[name] = m[name].Beaten
}

func (g *Game) Fight(screen *ebiten.Image, v string, m map[string]player.Player, PV *int, s *RPG.Save) {
	if RPG.MainMenuID == "Init_Fight" {
		*PV = m[v].PV
		PlayerPV = g.Player.PV
		RPG.MainMenuID = "Fight"
	} else if RPG.MainMenuID == "Fight" {
		if *PV > 0 {
			chapter.CanMove = false
			// draw the Background
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(0, 0)
			screen.DrawImage(RPG.BackgroundImage, op)
			// draw the Player at the left bottom of the screen
			op = &ebiten.DrawImageOptions{}
			op.GeoM.Scale(2, 2)
			op.GeoM.Translate(10, float64(ScreenResHeight)-74)
			screen.DrawImage(RPG.PlayerImage, op)
			ebitenutil.DrawRect(screen, 9, float64(ScreenResHeight)-79, 100, 5, color.RGBA{0, 0, 0, 170})
			ebitenutil.DrawLine(screen, 10, float64(ScreenResHeight)-78, 100, float64(ScreenResHeight)-78, color.RGBA{200, 200, 200, 170})
			ebitenutil.DrawLine(screen, 10, float64(ScreenResHeight)-78, (float64(PlayerPV)*100)/float64(g.Player.PV), float64(ScreenResHeight)-78, color.RGBA{0, 255, 0, 170})
			ebitenutil.DrawLine(screen, 10, float64(ScreenResHeight)-76, 100, float64(ScreenResHeight)-76, color.RGBA{200, 200, 200, 170})
			ebitenutil.DrawLine(screen, 10, float64(ScreenResHeight)-76, float64(g.Player.PA)*10, float64(ScreenResHeight)-76, color.RGBA{255, 0, 0, 170})
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
			RPG.Button(screen, !YourTurn, ScreenResWidth-180, ScreenHeight-70, "Attack", "Attack")
			RPG.Button(screen, !YourTurn, ScreenResWidth-180, ScreenHeight-50, "Regen", "Regen")
			RPG.Button(screen, !YourTurn, ScreenResWidth-180, ScreenHeight-30, "Run", "Run")
			//---------------------------------------------------------------------------------------------------------
			fmt.Println(PlayerPV, "pv actuel")
			fmt.Println((float64(PlayerPV)*100)/(float64(g.Player.PV)), "pv actuel en %")
			switch RPG.MainMenuID {
			case "Attack":
				fmt.Println((float64(PlayerPV)*100)/(float64(g.Player.PV)), "pv%")
				miss := GetMiss(0, 2)
				critical := Getcritical(0, 3)
				// fmt.Println("debut ", *PV)
				if miss != 0 {
					*PV = *PV
					// fmt.Println("miss atk: ", *PV)
				} else {
					if critical == 0 || critical == 1 {
						*PV -= g.Player.PA
						// fmt.Println("pas crit atk:", *PV)
					} else {
						*PV -= g.Player.PA * critical
						// fmt.Println("crit atk: ", *PV)
					}
				}
				// fmt.Println("fin ", *PV)
				YourTurn = false
				RPG.MainMenuID = "Fight"
			case "Regen":
				if (PlayerPV + g.Player.PD) > 25 {
					PlayerPV = g.Player.PV
					// RPG.PrintonTime(screen, "You are full HP",10,10, 0)
				} else if (PlayerPV + g.Player.PD) <= g.Player.PV {
					PlayerPV = PlayerPV + g.Player.PD
				}
				YourTurn = false
				RPG.MainMenuID = "Fight"
			case "Run":
				InFight = false
				RPG.MainMenuID = "Chp_1_0"
				chapter.CanMove = true
			}
			if !YourTurn {
				PlayerPV -= m[v].PA
				RPG.MainMenuID = "Fight"
				// RPG.PrintonTime(screen, m[v].Nom+" attack you, damage "+string(m[v].PA), 10, 10, 2)
				YourTurn = true
			}
		}
		if *PV <= 0 {
			// draw the Background
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(0, 0)
			screen.DrawImage(RPG.BackgroundImage, op)
			ebitenutil.DrawRect(screen, 0, 0, 256, 144, color.RGBA{0, 0, 0, 170})
			ebitenutil.DebugPrintAt(screen, "You've beaten "+v, 10, 10)
			RPG.Button(screen, false, ScreenResWidth-70, ScreenHeight-32, "Continue", "Beaten")
			if RPG.MainMenuID == "Beaten" {
				InFight = false
				chapter.CanMove = true
				temp := Mob[MobName]
				temp.Beaten = true
				Mob[MobName] = temp
				g.Mob = Mob
				if Mob[MobName].Type == "Master Boss" {
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
					g.NewGame(screen, s)
				} else {
					chapter.NextChapter = chapter.Current_Level
				}
				RPG.MainMenuID = chapter.NextChapter
				SetMobVariable(Mob, MobName)
				RPG.UpdateSave(&RPG.Save{CanLoad: true, Chapter: chapter.Current_Level, PlayerX: g.Player.PlayerX, PlayerY: g.Player.PlayerY, PV: PlayerPV, PA: g.Player.PA, PD: g.Player.PD, MobX: MobX, MobY: MobY, MobPV: MobPV, MobPA: MobPA, MobPD: MobPD, MobBeaten: MobBeaten, MobImage: MobImage})
				// fmt.Println(Mob)
			}
		} else if PlayerPV < 0 {
			// draw the Background
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(0, 0)
			screen.DrawImage(RPG.BackgroundImage, op)
			ebitenutil.DrawRect(screen, 0, 0, 256, 144, color.RGBA{0, 0, 0, 170})
			ebitenutil.DebugPrintAt(screen, "You are beaten by "+v, 10, 10)
			RPG.DeleteSave()
			RPG.Button(screen, false, ScreenResWidth-70, ScreenHeight-32, "Return Main Menu", "")
		}
	}
}

func Getcritical(minLimit int, maxlimit int) int {
	rand.Seed(time.Now().UnixNano())
	rndCrit := rand.Intn(maxlimit-minLimit) + minLimit
	// fmt.Println("nb crit ", rndCrit)
	return rndCrit
}

func GetMiss(minLimit int, maxlimit int) int {
	rand.Seed(time.Now().UnixNano())
	rndMiss := rand.Intn(maxlimit-minLimit) + minLimit
	// fmt.Println("nb miss ", rndMiss)
	return rndMiss
}

func SetPlayer(screen *ebiten.Image, g *Game) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.Player.PlayerX)-15, float64(g.Player.PlayerY)-15)
	screen.DrawImage(RPG.PlayerImage, op)
}

func (g *Game) CheckButtonID(ID string, screen *ebiten.Image, s *RPG.Save) {
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
		s.Chapter = "Int_1_P"
		chapter.Current_Level = s.Chapter
		s.CanLoad = false
		g.NewGame(screen, &RPG.Save{})
		RPG.CreateSave(&RPG.Save{CanLoad: true, Chapter: chapter.Current_Level, PlayerX: g.Player.PlayerX, PlayerY: g.Player.PlayerY, PV: PlayerPV, PA: g.Player.PA, PD: g.Player.PD})
	case "Load_Game":
		test := RPG.LoadSave(&RPG.Save{})
		g.NewGame(screen, &test)
	case "Settings":
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(0, 0)
		screen.DrawImage(RPG.BackgroundImage, op)
		ebitenutil.DrawRect(screen, 0, 0, 256, 144, color.RGBA{0, 0, 0, 170})
		ebitenutil.DebugPrintAt(screen, "Settings", 10, 10)
		RPG.UpdateSave(&RPG.Save{CanLoad: true, Chapter: chapter.Current_Level, PlayerX: g.Player.PlayerX, PlayerY: g.Player.PlayerY, PV: PlayerPV, PA: g.Player.PA, PD: g.Player.PD, MobX: MobX, MobY: MobY, MobPV: MobPV, MobPA: MobPA, MobPD: MobPD, MobBeaten: MobBeaten})
		RPG.Button(screen, false, ScreenResWidth-70, ScreenHeight-32, "Ok", "")
		// Settings
		// Fullscreen
		ebitenutil.DebugPrintAt(screen, "Fullscreen", 22, 30)
		if !ebiten.IsFullscreen() {
			ebitenutil.DrawRect(screen, 10, 34, 9, 9, color.RGBA{0, 0, 0, 170})
			if RPG.MouseX > 10 && RPG.MouseX < 19 && RPG.MouseY > 34 && RPG.MouseY < 43 && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
				ebitenutil.DrawRect(screen, 10, 34, 9, 9, color.RGBA{255, 255, 255, 170})
				ebiten.SetFullscreen(true)
			}
		} else {
			ebitenutil.DrawRect(screen, 10, 34, 9, 9, color.RGBA{255, 255, 255, 170})
			if RPG.MouseX > 10 && RPG.MouseX < 19 && RPG.MouseY > 34 && RPG.MouseY < 43 && ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
				ebitenutil.DrawRect(screen, 10, 34, 9, 9, color.RGBA{0, 0, 0, 170})
				ebiten.SetFullscreen(false)
			}
		}
		//
	case "Quit":
		SetMobVariable(Mob, MobName)
		RPG.UpdateSave(&RPG.Save{CanLoad: true, Chapter: chapter.Current_Level, PlayerX: g.Player.PlayerX, PlayerY: g.Player.PlayerY, PV: PlayerPV, PA: g.Player.PA, PD: g.Player.PD, MobX: MobX, MobY: MobY, MobPV: MobPV, MobPA: MobPA, MobPD: MobPD, MobBeaten: MobBeaten})
		os.Exit(0)
	case "Int_1_P":
		op := &ebiten.DrawImageOptions{}
		screen.DrawImage(RPG.BackgroundImage, op)
		ebitenutil.DrawRect(screen, 0, 0, 256, 144, color.RGBA{0, 0, 0, 170})
		print(paragraph1, screen)
		RPG.Button(screen, false, 256-66, 144-18, "Next", "Int_2_P")
	case "Int_2_P":
		op := &ebiten.DrawImageOptions{}
		screen.DrawImage(RPG.BackgroundImage, op)
		ebitenutil.DrawRect(screen, 0, 0, 256, 144, color.RGBA{0, 0, 0, 170})
		print(paragraph2, screen)
		RPG.Button(screen, false, 256-66, 144-18, "Next", "Int_3_P")
		RPG.Button(screen, false, 2, 144-18, "Previous", "Int_1_P")
	case "Int_3_P":
		op := &ebiten.DrawImageOptions{}
		screen.DrawImage(RPG.BackgroundImage, op)
		ebitenutil.DrawRect(screen, 0, 0, 256, 144, color.RGBA{0, 0, 0, 170})
		print(paragraph3, screen)
		RPG.Button(screen, false, 256-66, 144-18, "Next", "Int_4_P")
		RPG.Button(screen, false, 2, 144-18, "Previous", "Int_2_P")
	case "Int_4_P":
		op := &ebiten.DrawImageOptions{}
		screen.DrawImage(RPG.BackgroundImage, op)
		ebitenutil.DrawRect(screen, 0, 0, 256, 144, color.RGBA{0, 0, 0, 170})
		print(paragraph4, screen)
		RPG.Button(screen, false, 256-66, 144-18, "Next", "Int_Next_Chapter")
		RPG.Button(screen, false, 2, 144-18, "Previous", "Int_3_P")
	case "Int_Next_Chapter":
		s.Chapter = "Chp_1_0"
		chapter.Current_Level = "Chp_1_0"
		g.NewGame(screen, s)
	case "Chp_1_0":
		chapter.CanMove = true
		chapter.Current_Level = "Chp_1_0"
		op := &ebiten.DrawImageOptions{}
		screen.DrawImage(RPG.Background_Ch1, op)
		draw.DrawMob(Mob, screen)
		SetPlayer(screen, g)
		// fmt.Printf("PlayerX: %v, PlayerY: %v \n", g.Player.PlayerX, g.Player.PlayerY)
		for _, v := range Mob {
			v.MaxX = v.PlayerX + v.Image.Bounds().Dx()
			v.MaxY = v.PlayerY + v.Image.Bounds().Dy()
			if (g.Player.PlayerX+25 >= v.PlayerX && g.Player.PlayerX+25 <= v.MaxX) && (g.Player.PlayerY+25 >= v.PlayerY && g.Player.PlayerY+25 <= v.MaxY) {
				if !v.Beaten {
					if v.Type != "Event" {
						//fight
						ebitenutil.DrawRect(screen, 2, float64(ScreenHeight)-20, float64(ScreenResWidth), float64(ScreenHeight), color.RGBA{0, 0, 0, 170})
						ebitenutil.DebugPrintAt(screen, "For Fight "+v.Nom+", Press [ENTER]", 2, ScreenHeight-20)
						//RPG.Fight(screen, g.Player, v)
						if ebiten.IsKeyPressed(ebiten.KeyEnter) || InFight {
							InFight = true
							RPG.MainMenuID = "Init_Fight"
							MobName = v.Nom
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
		screen.DrawImage(RPG.Background_Ch2, op)
		draw.DrawMob(Mob, screen)
		SetPlayer(screen, g)
		// fmt.Printf("PlayerX: %v, PlayerY: %v \n", g.Player.PlayerX, g.Player.PlayerY)
		for _, v := range Mob {
			v.MaxX = v.PlayerX + v.Image.Bounds().Dx()
			v.MaxY = v.PlayerY + v.Image.Bounds().Dy()
			if (g.Player.PlayerX+25 >= v.PlayerX && g.Player.PlayerX+25 <= v.MaxX) && (g.Player.PlayerY+25 >= v.PlayerY && g.Player.PlayerY+25 <= v.MaxY) {
				if !v.Beaten {
					//fight
					ebitenutil.DrawRect(screen, 2, float64(ScreenHeight)-20, float64(ScreenResWidth), float64(ScreenHeight), color.RGBA{0, 0, 0, 170})
					ebitenutil.DebugPrintAt(screen, "For Fight "+v.Nom+", Press [ENTER]", 2, ScreenHeight-20)
					//RPG.Fight(screen, g.Player, v)
					if ebiten.IsKeyPressed(ebiten.KeyEnter) || InFight {
						InFight = true
						RPG.MainMenuID = "Init_Fight"
						MobName = v.Nom
					}
				}
			}
		}
	case "Init_Fight":
		g.Fight(screen, MobName, Mob, &Pv, &RPG.Save{})
	case "Fight":
		g.Fight(screen, MobName, Mob, &Pv, &RPG.Save{})
	case "Pause":
		op := &ebiten.DrawImageOptions{}
		screen.DrawImage(RPG.BackgroundImage, op)
		ebitenutil.DrawRect(screen, 0, 0, 256, 144, color.RGBA{0, 0, 0, 170})
		ebitenutil.DrawRect(screen, 0, 0, 0, 144, color.RGBA{0, 0, 0, 170})
		SetMobVariable(Mob, MobName)
		RPG.UpdateSave(&RPG.Save{CanLoad: true, Chapter: chapter.Current_Level, PlayerX: g.Player.PlayerX, PlayerY: g.Player.PlayerY, PV: PlayerPV, PA: g.Player.PA, PD: g.Player.PD, MobX: MobX, MobY: MobY, MobPV: MobPV, MobPA: MobPA, MobPD: MobPD, MobBeaten: MobBeaten})
		RPG.Button(screen, false, 256-66, 144-18, "Resume", chapter.Current_Level)
		RPG.Button(screen, false, 256-66, 2, "Main Menu", "")
		RPG.Button(screen, false, 2, 144-18, "Settings", "Settings")
		RPG.Button(screen, false, 2, 2, "Quit", "Quit")
		// fmt.Printf("Pause: %v\n", Pause)
		if inpututil.IsKeyJustReleased(ebiten.KeyEscape) && Pause && !IsWait {
			RPG.MainMenuID = chapter.Current_Level
			Pause = false
			Wait(1 * time.Second)
		}
	}
}

// MainMenu is the main menu of the game
func (g *Game) MainMenu() {
	// set the game name
	g.Name = "YLock's"
	// set the game version
	g.Version = "0.2.13"
	// run the game
	ebiten.SetWindowIcon(RPG.IconImage)
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle(g.Name + " " + g.Version)
	ebiten.SetWindowResizable(true)
	ebiten.MaximizeWindow()
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}

// Update updates the game state.
func (g *Game) Update(screen *ebiten.Image) error {
	//fmt.Println(RPG.MainMenuID)
	RPG.SetMousePosition()
	//CheckButtonID(RPG.MainMenuID, screen, RPG.Save{})
	g.Draw(screen)
	// set frame rate
	if chapter.CanMove {
		if ebiten.IsKeyPressed(ebiten.KeyUp) {
			if g.Player.PlayerY > 16 {
				if g.Player.PlayerY < 88 && g.Player.PlayerX < 73 && RPG.MainMenuID == "Chp_1_0" {
					fmt.Print()
				} else if (g.Player.PlayerY <= 88 && (g.Player.PlayerX > 73 && g.Player.PlayerX < 120)) || (g.Player.PlayerY <= 88 && (g.Player.PlayerX > 150 && g.Player.PlayerX < 256)) && RPG.MainMenuID == "Chp_1_0" {
					fmt.Print()
				} else {
					g.Player.PlayerY -= 1
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
				if (g.Player.PlayerY < 88 && g.Player.PlayerX < 74) || (g.Player.PlayerY <= 88 && (g.Player.PlayerX > 73 && g.Player.PlayerX < 120)) && RPG.MainMenuID == "Chp_1_0" {
					fmt.Print()
				} else {
					g.Player.PlayerX -= 1
				}
			}
		}
		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			if g.Player.PlayerX < 256-16 {
				if !Mob["Card Reader"].Beaten {
					if g.Player.PlayerX < 74-16 {
						g.Player.PlayerX += 1
					}
				} else if (g.Player.PlayerY <= 88 && (g.Player.PlayerX > 135 && g.Player.PlayerX < 205)) || (g.Player.PlayerX >= 205) && RPG.MainMenuID == "Chp_1_0" {
					fmt.Print()
				} else {
					g.Player.PlayerX += 1
				}
			}
		}
		// fmt.Printf("Pause: %v\n", Pause)
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			RPG.MainMenuID = "Pause"
			chapter.CanMove = false
			Pause = true
		}
	}
	if IsWait {
		WaitSys()
	}
	// fmt.Print(g.PlayerX, g.Player.PlayerY)
	ebiten.SetMaxTPS(60)
	return nil
}

// Draw draws the game screen.
func (g *Game) Draw(screen *ebiten.Image) {
	g.Player.PA = 5
	g.Player.PD = 15
	g.Player.PV = 25
	g.CheckButtonID(RPG.MainMenuID, screen, &RPG.Save{})
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

func Wait(Duration time.Duration) {
	if !IsWait {
		IsWait = true
		WaitDuration = Duration
		WaitTemp = time.Now()
		fmt.Println("Wait init")
	}
}

func WaitSys() {
	j := time.Now()
	if j.Sub(WaitTemp) >= WaitDuration {
		IsWait = false
		fmt.Println("Wait finish")
	}
}
