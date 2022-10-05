package player

import (
	"math/rand"
	"time"

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
	PlayerPV        = 0
)

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