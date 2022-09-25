package RPG

type Game struct {
	Name string
	Version string
	PlayerX int
	PlayerY int
	ScreenHeight int
	ScreenWidth int
}

func (c *Game) Master() {
	c.MainMenu()
}