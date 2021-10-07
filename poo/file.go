package main

import (
	"fmt"
)

func main() {
	h := NewHero(1, "legion", 375.0, 2500.0, []string{"Manta Style", "Power B"})
	fmt.Println(h)
}
