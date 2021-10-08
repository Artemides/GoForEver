package main

import "fmt"

func main() {
	personas := make(map[string]int)
	personas["jose"] = 18
	personas["julio"] = 19
	personas["lucas"] = 12
	personas["maria"] = 13
	alumnosNotas := make(map[string][]uint8)
	alumnosNotas["mijael"] = []uint8{12, 13, 14}
	alumnosNotas["rafael"] = []uint8{11, 14, 17}
	alumnosNotas["lucas"] = []uint8{12, 13, 14}

	for c, v := range alumnosNotas {
		var promedio float64
		for _, e := range v {
			promedio += float64(e)
		}
		fmt.Printf("Alumno: %s  -> promedio: %.2f \n", c, promedio/float64(len(v)))
	}
	value, ok := alumnosNotas["rafael"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("no existe")
	}

}
