package main

import (
	"github.com/goutte/git-spend/cmd"
	"log"
)

func main() {
	log.SetFlags(0)

	err := cmd.Execute()
	if err != nil {
		log.Fatalln("failure:", err)
	}
}
