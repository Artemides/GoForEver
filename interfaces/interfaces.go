package main

import (
	"fmt"
	"math"
)

/**STRUCT SQUARE */
type Square struct {
	base float64
}

func (Sq Square) area() float64 {
	return Sq.base * Sq.base
}

/**STRUCT RECTANGLE*/
type Rectangle struct {
	base   float64
	height float64
}

func (Rc Rectangle) area() float64 {
	return Rc.base * Rc.height
}

/* Interface */
type figuras2D interface {
	area() float64
}

func calcular(f figuras2D) float64 {
	return f.area()
}

/**/
func main() {
	cuadrado := Square{base: 6}
	rectangulo := Rectangle{base: 4, height: 12}

	fmt.Println(calcular(cuadrado))
	fmt.Println(calcular(rectangulo))

	arrInterface := []interface{}{2, " ", 2.23, math.Pi}
	fmt.Println(arrInterface...)
}
