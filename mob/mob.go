package mob

import (
	"RPG/player"

	"github.com/hajimehoshi/ebiten"
)

var (
	MobName         = ""
	Pv              = 0
	MobX            = map[string]int{}
	MobY            = map[string]int{}
	MobPV           = map[string]int{}
	MobPA           = map[string]int{}
	MobPD           = map[string]int{}
	MobBeaten       = map[string]bool{}
	MobImage        = map[string]*ebiten.Image{}

	Mob = map[string]player.Player{}
)

func SetMobVariable(m map[string]player.Player, name string) {
	MobX[name] = m[name].PlayerX
	MobY[name] = m[name].PlayerY
	MobPV[name] = m[name].PV
	MobPA[name] = m[name].PA
	MobPD[name] = m[name].PD
	MobBeaten[name] = m[name].Beaten
}