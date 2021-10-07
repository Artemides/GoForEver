package main

import (
	"fmt"
	"sync"
)

func say(str string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(str)
}
func main9() {
	var wg sync.WaitGroup
	wg.Add(1)
	fmt.Println("Hello")
	go say("World", &wg)
	wg.Wait()
}
