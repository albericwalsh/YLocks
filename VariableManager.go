package RPG

import "github.com/hajimehoshi/ebiten"

var (
	MouseX = 0
	MouseY = 0
)

func SetMousePosition() {
	MouseX, MouseY = ebiten.CursorPosition()
}
