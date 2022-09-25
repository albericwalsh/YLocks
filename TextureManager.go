package RPG

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

//load all the textures

var (
	backgroundImage   *ebiten.Image
	buttonImage       *ebiten.Image
	LockedButtonImage *ebiten.Image
	OverButtonImage   *ebiten.Image
)

// load the images
func Textures_init() {
	// Background
	backgroundImage, _, _ = ebitenutil.NewImageFromFile("Assets/Gui_Textures/Background.png", ebiten.FilterDefault)
	// Buttons
	buttonImage, _, _ = ebitenutil.NewImageFromFile("Assets/Gui_Textures/Button.png", ebiten.FilterDefault)
	LockedButtonImage, _, _ = ebitenutil.NewImageFromFile("Assets/Gui_Textures/LockedButton.png", ebiten.FilterDefault)
	OverButtonImage, _, _ = ebitenutil.NewImageFromFile("Assets/Gui_Textures/OverButton.png", ebiten.FilterDefault)
}
