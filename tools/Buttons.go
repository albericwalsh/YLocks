package tools

// make a setting button

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)
 var (
	MainMenuID = ""
	Fix = false
 )
func Center(Object *ebiten.Image, ref int) int {
	Object_Width := Object.Bounds().Dx()
	// Center the object
	coor := (ref / 2) - (Object_Width / 2)
	return coor
}

func CheckMouseOverButton(x,y int, Button *ebiten.Image) bool {
	//check if the mouse is over the button
	if MouseX > x && MouseX < x+Button.Bounds().Dx() && MouseY > y && MouseY < y+Button.Bounds().Dy() {
		return true
	}
	return false
}

func Button(screen *ebiten.Image, Locked bool, x, y int, text string, ID string) {
	if Locked {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(x), float64(y))
		screen.DrawImage(LockedButtonImage, op)
		// draw the text and set the position to Center of the button
		ebitenutil.DebugPrintAt(screen, text, x+Center(LockedButtonImage, LockedButtonImage.Bounds().Dx()), y)
	} else if CheckMouseOverButton(x, y, ButtonImage) && !Locked {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x+1), float64(y+1))
			screen.DrawImage(OverButtonImage, op)
			// draw the text
			ebitenutil.DebugPrintAt(screen, text, x+1+Center(LockedButtonImage, LockedButtonImage.Bounds().Dx()), y+1)
			Fix = true
		} else if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x-1), float64(y-1))
			screen.DrawImage(OverButtonImage, op)
			// draw the text
			ebitenutil.DebugPrintAt(screen, text, x-1+Center(LockedButtonImage, LockedButtonImage.Bounds().Dx()), y-1)
			if Fix {
				Fix = false
				MainMenuID = ID
			}
		} else {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x), float64(y))
			screen.DrawImage(OverButtonImage, op)
			// draw the text
			ebitenutil.DebugPrintAt(screen, text, x+Center(LockedButtonImage, LockedButtonImage.Bounds().Dx()), y)
		}
	} else {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(x), float64(y))
		screen.DrawImage(ButtonImage, op)
		// draw the text
		ebitenutil.DebugPrintAt(screen, text, x+Center(LockedButtonImage, LockedButtonImage.Bounds().Dx()), y)
	}
}
