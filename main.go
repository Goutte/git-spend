package main

import (
	"github.com/goutte/git-spend/cmd"
	"log"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalln("Failure:", err)
	}
}
