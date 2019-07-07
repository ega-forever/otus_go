package main

import (
	"fmt"
	"os"
	"otus/api"
)

func main() {

	text := os.Args[1]
	pairs := api.WordCount(text, 10)

	fmt.Println(pairs)

}
