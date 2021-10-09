package main

import (
	"fmt"

	h "github.com/Artemides/GoForEver/tree/master/poo/heroes"
)

func main() {
	h := h.NewHero(1, "legion", 375.0, 2500.0, []string{"Manta Style", "Power B"})
	fmt.Println(h)
}
