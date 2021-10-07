package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Printf("steps: %d", stepsToZero(-63))
	//fmt.Printf(" 2658 : %d", reverse(2658))
	n := 20
	sum, mul, sq, cube := sumN(n)
	fmt.Printf("Suma: %d , multiplicación: %d , cuadrado: %d, cubo: %d", sum, mul, sq, cube)
}
func reverse(n int) int {
	aux := 0
	for n%10 > 0 {
		aux = aux*10 + n%10
		n = n / 10
	}
	return aux
}
func sumN(n int) (sum int, mul int, sq int, cube int) {
	i := 1
	mul = 1
	for i <= n {
		sum += i
		mul *= i
		sq = int(math.Pow(float64(i), 2)) + sq
		cube = int(math.Pow(float64(i), 3)) + sq
		i++
	}
	return sum, mul, sq, cube
}
func stepsToZero(n int) int {
	i := 0
	n = int(math.Abs(float64(n)))
	for n >= 0 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = n - 1
		}
		fmt.Println("n: ", n)
		i++
	}
	return i
}
func dist2Points(x1, x2, y1, y2 float64) float64 {
	return math.Pow(math.Pow(x2-x1, 2)+math.Pow(y2-y1, 2), 0.5)
}
func minMax(v1, v2 float64) (min, max float64) {
	min, max = v1, v1
	if v1 >= v2 {
		min, max = v2, v1
	} else {
		min, max = v1, v2
	}
	return min, max
}

func main4() {
	x, y := 2, 3

	fmt.Println("testing Numbers")
	fmt.Println("es par: ", even(x))
	fmt.Println("es par: ", even(y))
	user, password := "jose", "192168101"
	fmt.Println("Iniciando sesión....")
	if validatePassword(user, password) {
		println(" Logeado")
	} else {
		fmt.Println("Clave incorrecta")
	}

}
func even(n int) bool {
	if n%2 == 0 {
		return true
	} else {
		return false
	}
}
func validatePassword(user, password string) bool {
	if user == "jose" && password == "192168101" {
		return true
	} else {
		return false
	}
}
func main3() {
	x1, x2, y1, y2 := 12.3, 4.2, 25.7, -32.2
	fmt.Printf("distancia: %.3f\n", dist2Points(x1, x2, y1, y2))
	min, _ := minMax(x1, y1)
	_, max := minMax(x2, y2)
	fmt.Printf("mayor: %.1f, mayor: %.1f\n", min, max)
	fmt.Println("Promedio", promedio(x1, x2, y1, y2))

	for i := 9; i >= 0; i-- {
		fmt.Printf("i: %d\n", i)
	}
	index := 9
	for index >= 0 {
		fmt.Println(index)
		index--
	}

}
func promedio(x1, x2, x3, x4 float64) float64 {
	return (x1 + x2 + x3 + x4) / 4
}

//creación de funciones
//funciones con parámetros
//funciones con return
//funciones con return multiple
//funciones anonimas

func main2() {
	x1 := 2
	x2 := 4
	y1 := 8
	y2 := 12
	distancia := math.Round(math.Pow(math.Pow(float64(x2-x1), 2)+math.Pow(float64(y2-y1), 2), 0.5))

	fmt.Println("hellor world")
	fmt.Println("distancia: ", distancia)
	/*
		uint8= enteros positivos desde 0 a 2^7
		uint16	0 "" false
		uint32
		uint64
		int16
		int32
		int64
		flot32
		float64
	*/
	d := 10.0
	h := 18.0
	r := d / 2
	AB := math.Pi * math.Pow(r, 2)
	vol := AB * h
	fmt.Printf("Radio: %.2f , Area Base: %.2f, volumen: %.2f", r, AB, vol)
	// buena práctica indicar el valor en el printf
	message := fmt.Sprintf("Radio: %.2f , Area Base: %.2f, volumen: %.2f", r, AB, vol)
	fmt.Println(message)
}
