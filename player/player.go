package player

import (
	"github.com/hajimehoshi/ebiten"
)

type Player struct {
	PlayerX int
	PlayerY int
	MaxX    int
	MaxY    int
	Nom     string
	PV      int
	PA      int
	PD      int
	Beaten  bool
	Type    string
	Image   *ebiten.Image
}

var (
	Mob map[string]Player
)

