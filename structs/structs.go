package main

import (
	"fmt"
)

type man struct {
	dni      string
	name     string
	lastName string
	age      int
	city     string
	state    string
}

/*func main() {
	s := []pk.Student{}
	var ss pk.Student
	ss.Id = "asdasd"
	s = append(s, ss)
	fmt.Println(s)
	pk.Identify()
}*/

func main() {
	people := []man{}
	fmt.Println("how many people? ")
	fmt.Printf(" n: ")
	n := 0
	fmt.Scan(&n)
	for n > 0 {
		fmt.Println("DNI Name LastName age city state")
		var pple man
		fmt.Scanf("%s %s %s %d %s %s", &pple.dni, &pple.name, &pple.lastName, &pple.age, &pple.city, &pple.state)
		people = append(people, pple)
		n--
	}
	fmt.Println(people)
}
