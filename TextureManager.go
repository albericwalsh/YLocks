package RPG

import (
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

//load all the textures

var (
	BackgroundImage   *ebiten.Image
	ButtonImage       *ebiten.Image
	LockedButtonImage *ebiten.Image
	OverButtonImage   *ebiten.Image
	Card_Reader       *ebiten.Image
	Background_Ch1    *ebiten.Image
	PlayerImage       *ebiten.Image
	PaulImage         *ebiten.Image
	Vitaly            *ebiten.Image
	Warning           *ebiten.Image
	Success           *ebiten.Image
	wesh              *ebiten.Image
	IconImage         []image.Image
)

// load the images
func Textures_init() {
	// Gui
	wesh, _, _ = ebitenutil.NewImageFromFile("assets/Icon.png", ebiten.FilterDefault)
	IconImage = append(IconImage, wesh.SubImage(image.Rect(0, 0, 32, 32)).(*ebiten.Image))
	// Background
	BackgroundImage, _, _ = ebitenutil.NewImageFromFile("Assets/Gui_Textures/Background.png", ebiten.FilterDefault)
	// Buttons
	ButtonImage, _, _ = ebitenutil.NewImageFromFile("Assets/Gui_Textures/Button.png", ebiten.FilterDefault)
	LockedButtonImage, _, _ = ebitenutil.NewImageFromFile("Assets/Gui_Textures/LockedButton.png", ebiten.FilterDefault)
	OverButtonImage, _, _ = ebitenutil.NewImageFromFile("Assets/Gui_Textures/OverButton.png", ebiten.FilterDefault)
	// Card Reader
	Card_Reader, _, _ = ebitenutil.NewImageFromFile("Assets/Map_Textures/Card_Reader.png", ebiten.FilterDefault)
	// Chapter Backgrounds
	Background_Ch1, _, _ = ebitenutil.NewImageFromFile("Assets/Map_Textures/Background_Ch1.png", ebiten.FilterDefault)
	// Player
	PlayerImage, _, _ = ebitenutil.NewImageFromFile("Assets/Character_Textures/player.png", ebiten.FilterDefault)
	PaulImage, _, _ = ebitenutil.NewImageFromFile("Assets/Character_Textures/Paul.png", ebiten.FilterDefault)
	Vitaly, _, _ = ebitenutil.NewImageFromFile("Assets/Character_Textures/Vitaly.png", ebiten.FilterDefault)
	// icon
	Warning, _, _ = ebitenutil.NewImageFromFile("Assets/Gui_Textures/Warning.png", ebiten.FilterDefault)
	Success, _, _ = ebitenutil.NewImageFromFile("Assets/Gui_Textures/Success.png", ebiten.FilterDefault)
}
