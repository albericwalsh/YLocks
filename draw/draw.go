package draw

import (
	"RPG/tools"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"

	"RPG/player"
)

func DrawMob(m map[string]player.Player, screen *ebiten.Image) {
	for _, v := range m {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(v.PlayerX)-15, float64(v.PlayerY)-15)
		// log.Fatal(v.Image)
		screen.DrawImage(v.Image, op)
		// draw the name of the mob at the top of the mob
		ebitenutil.DebugPrintAt(screen, v.Nom, v.PlayerX-15, v.PlayerY-32)
		if v.Type == "Event"{
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(v.PlayerX+10), float64(v.PlayerY)+10)
			screen.DrawImage(tools.Info, op)
		} else if v.Beaten {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(v.PlayerX+10), float64(v.PlayerY)+10)
			screen.DrawImage(tools.Success, op)
		} else {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(v.PlayerX+10), float64(v.PlayerY)+10)
			screen.DrawImage(tools.Warning, op)
		}
	}
}