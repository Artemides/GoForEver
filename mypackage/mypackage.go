package mypackage

import "fmt"

//StudenPublic defines a student struct
type Student struct {
	Id     string
	Name   string
	Career string
}

func Identify() {
	fmt.Println("Hola")
}
