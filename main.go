package main

import (
	"github.com/goutte/gitime/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalln("Failure:", err)
	}
}
