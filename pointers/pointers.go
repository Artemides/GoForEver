package main

import "fmt"

func main() {
	a := 100
	b := &a
	fmt.Println(b)
	*b = 25
	fmt.Println(*b)
}
