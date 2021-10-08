package main

import (
	"fmt"
	"os"
	"strings"
)

//000000
//000001
// ...
//ffffff -> 16 777 215
func main() {

	/*         ROUTINE 1 */
	routine1()
	/*****NORMAL ***/
	/*file, err := os.Create("C:/Go/src/clases/DB/routines.txt")
	if err != nil {
		panic(err)
	}
	final := 16777215

	for i := 0; i <= final; i++ {
		_, err = file.WriteString(fmt.Sprintf("%06x\n", i))
		if err != nil {
			panic(err)
		}
	}*/
}

//
func routine1() {
	f, err := os.Create("C:/Go/src/clases/DB/2routinesGO.txt")
	if err != nil {
		panic(err)
	}

	/*Go WriteFile Routine*/
	outCh := make(chan string)
	doneWrite := make(chan struct{})
	go func() {
		for str := range outCh {
			_, err := f.WriteString(str)
			if err != nil {
				panic(err)
			}
		}
		doneWrite <- struct{}{}
	}()
	numGoRoutines := 10
	doneCh := make(chan struct{})
	final := 16777215
	for i := 0; i <= final; i = i + (final / numGoRoutines) + 1 {
		paso := i + (final / numGoRoutines)
		if paso > final {
			paso = final
		}
		fmt.Printf("ejecutanto %d %d\n", i, paso)
		//create go routine
		go calcNums(i, paso, f, outCh, doneCh)
	}
	doneNum := 0
	for doneNum < numGoRoutines {
		<-doneCh
		fmt.Println("Terminó un gopher!")
		doneNum++
	}
	close(outCh)
	<-doneWrite
	fmt.Println("Listo!! :D")
}
func calcNums(start, end int, f *os.File, resultCh chan string, doneCh chan struct{}) {
	var sBuilder strings.Builder
	for i := start; i <= end; i++ {
		_, err := f.WriteString(fmt.Sprintf("%06x\n", i))
		if err != nil {
			panic(err)
			//
		}
		fmt.Fprint(&sBuilder, fmt.Sprintf("%06x\n", i))
	}
	resultCh <- sBuilder.String()
	doneCh <- struct{}{}

}
