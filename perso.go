package RPG

import (
	"fmt"
)

// Attributs des personnages
type personnage struct {
	Nom   string
	PV    int
	PA    int
	PD    int
	Type  string
	Battu bool
}

// affichage des informations sur le personnage
func (p personnage) Affinfo() {
	fmt.Println("Vie du personnage", p.Nom, ":", p.PV)
	fmt.Println("Puissance du personnage", p.Nom, ":", p.PA)
	fmt.Println("Defense du personnage", p.Nom, ":", p.PD)
	fmt.Println("Type du personnage", p.Nom, ":", p.Type)
}
