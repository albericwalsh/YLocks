package RPG

// make a setting button

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

func Button(screen *ebiten.Image, Locked, Over bool, x, y int, text string) {
	if Locked {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(x), float64(y))
		op.GeoM.Scale(3, 3)
		screen.DrawImage(LockedButtonImage, op)
		// draw the text
		ebitenutil.DebugPrintAt(screen, text, x+10, y+10)
	} else if Over {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(x), float64(y))
		op.GeoM.Scale(3, 3)
		screen.DrawImage(OverButtonImage, op)
		// draw the text
		ebitenutil.DebugPrintAt(screen, text, x+10, y+10)
	} else {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(x), float64(y))
		op.GeoM.Scale(3, 3)
		screen.DrawImage(buttonImage, op)
		// draw the text
		ebitenutil.DebugPrintAt(screen, text, x+10, y+10)
	}
}
