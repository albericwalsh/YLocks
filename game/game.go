package game

import (
	"RPG/tools"
	"RPG/chapter"
	"RPG/mob"
	"RPG/player"

	"github.com/hajimehoshi/ebiten"
)

type Game struct {
	Name    string
	Version string
	Player  player.Player
	Mob map[string]player.Player
}

var (
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
