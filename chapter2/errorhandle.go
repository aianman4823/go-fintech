package chapter2

import (
	"fmt"
	"log"
	"os"
)

func ErrorHandle() {
	file, err := os.Open("./main.go")
	if err != nil {
		log.Fatal("Error")
	}
	defer file.Close()

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln("Error")
	}
	fmt.Println(count, string(data))

	err = os.Chdir("test")
	if err != nil {
		log.Fatalln("Error")
	}

}
