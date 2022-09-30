package RPG

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"time"
)

var (
	MouseX = 0
	MouseY = 0
)

func SetMousePosition() {
	MouseX, MouseY = ebiten.CursorPosition()
}

func PrintonTime(screen *ebiten.Image, text string, x, y, timeint int ) {
	for i:= 0; i < timeint; i++ {
		ebitenutil.DebugPrint(screen, text)
		ebitenutil.DebugPrintAt(screen, text, x, y)
		time.Sleep(1 * time.Millisecond)
	}
}