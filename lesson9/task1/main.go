package main

import (
	"os"
	"otus/lesson9/task1/cmd"
)

func main() {

	args := os.Args[1:]

	if len(args) < 1 {
		os.Exit(0)
	}

	cmd.Execute(args[0], args[1:]...)
}
