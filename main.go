package main

import (
	"github.com/ega-forever/otus_go/cmd"
	"log"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
