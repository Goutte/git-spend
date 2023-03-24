package main

import (
	"fmt"
	"github.com/goutte/gitime/gitime"
	"github.com/tsuyoshiwada/go-gitlog"
	"log"
)

func main() {
	git := gitlog.New(&gitlog.Config{})

	commits, err := git.Log(nil, nil)
	if err != nil {
		log.Fatalln("Cannot read git log:", err)
	}

	ts := &gitime.TimeSpent{}
	for _, commit := range commits {
		ts.Add(gitime.CollectTimeSpent(commit.Subject))
		ts.Add(gitime.CollectTimeSpent(commit.Body))
	}

	fmt.Printf(ts.String() + "\n")
	fmt.Printf("%d minutes\n", ts.ToMinutes())
}
