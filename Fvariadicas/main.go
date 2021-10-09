package main

import (
	"fmt"
	"math"
)

func sum(values ...int) (res int) {
	for _, v := range values {
		res += v
	}
	return
}
func exp(values ...int) (double, tripple, quad int) {
	for _, v := range values {
		double += int(math.Pow(float64(v), 2))
		tripple += int(math.Pow(float64(v), 3))
		quad += int(int(math.Pow(float64(v), 3)))
	}
	return
}
func main() {
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8))
	fmt.Println(exp(2, 4, 8, 16))
}
