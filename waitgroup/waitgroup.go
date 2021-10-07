package main

import (
	"fmt"
	"sync"
	"time"
)

func Working(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Gopher %d Starting Sleep\n", id)
	time.Sleep(time.Second * 5)
	fmt.Printf("Gopher %d Done\n", id)
}
func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go Working(i, &wg)
	}
	wg.Wait()

}
