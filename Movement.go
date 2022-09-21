package RPG

import (
	"github.com/hajimehoshi/ebiten"
)

const (
	PlayerX = 0.0
	PlayerY = 0.0
)

func updatePlayer() {
	// player movement with ZQSD or arrow keys
	// While is pressed player moves
	if ebiten.IsKeyPressed(ebiten.KeyZ) || ebiten.IsKeyPressed(ebiten.KeyUp) {
		playerY -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown) {
		playerY += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyQ) || ebiten.IsKeyPressed(ebiten.KeyLeft) {
		playerX -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		playerX += 1
	}
}
