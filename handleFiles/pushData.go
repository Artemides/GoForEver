package handlefiles

import (
	"log"
	"os"
)

func PushData(str, path string) {
	/*file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		log.Fatalf("Erro by Opening the File %s", path)
	}
	file.WriteString(str)*/
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0644)
	defer file.Close()
	if err != nil {
		log.Fatalf("Error by Opening the file %s", path)
	}
	file.WriteString(str + "\n")

}
