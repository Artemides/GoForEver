package main

import (
	"fmt"

	handlefiles "github.com/Artemides/Go-forever/handleFiles"
	"github.com/Artemides/Go-forever/store"
)

/* CRUD*/

func menu() {
	fmt.Println("**************** MENU ************")
	fmt.Println("1. Insertar")
	fmt.Println("2. Eliminar")
	fmt.Println("3. Mostrar")
	fmt.Println("4. Editar ")
	fmt.Println("*************************************")
}

//
func insert(path string) {
	f := store.Fruits{}
	fmt.Printf(" Ingresar Fruta\n")
	id, name, cost := "", "", 0.0
	fmt.Printf("Id: ")
	fmt.Scan(&id)
	fmt.Printf("Fruta: ")
	fmt.Scan(&name)
	fmt.Printf("Precio S/. : ")
	fmt.Scan(&cost)
	fmt.Println(" ")
	f.SetValues(id, name, cost)
	handlefiles.PushData(f.String(), path)
	fmt.Printf("%v :Agregado.....", f)
}

func mainapp() {
	path := "C:/Go/src/clases/DB/data.txt"
	for true {
		menu()
		i := 0
		fmt.Printf("Opción: ")
		fmt.Scan(&i)
		fmt.Println("")
		switch i {
		case 1:
			insert(path)
		case 2:
			handlefiles.FetchData(path)
		default:
			fmt.Printf("esta opción aun no esta soportada.")
		}
	}

}
