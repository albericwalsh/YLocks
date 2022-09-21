package RPG


import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var (
	// player
	playerImage *ebiten.Image
	playerX     = 0.0
	playerY     = 0.0
)

func init() {
	// player
	playerImage, _ = ebiten.NewImage(16, 16, ebiten.FilterDefault)
	playerImage.Fill(color.White)
}

func update(screen *ebiten.Image) error {
	// player
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		playerY--
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		playerY++
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		playerX--
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		playerX++
	}

	// draw
	screen.Fill(color.Black)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(playerX, playerY)
	screen.DrawImage(playerImage, op)

	ebitenutil.DebugPrint(screen, "Hello, World!")

	return nil
}

func CreateWindow() {
	rand.Seed(time.Now().UnixNano())
	if err := ebiten.Run(update, screenWidth, screenHeight, 2, "Hello, World!"); err != nil {
		log.Fatal(err)
	}
}

