package handlefiles

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func FetchData(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error by Opening the file: %s", path)
	}
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		el := strings.Split(fileScanner.Text(), ",")
		fmt.Printf("Id: %v\tFruta: %v\tS/. %v\n", el[0], el[1], el[2])
	}
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error by reading the file %s", err)
	}
}
