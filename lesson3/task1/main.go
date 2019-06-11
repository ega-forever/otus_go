package main

import "fmt"
import "otus/lesson3/task1/api"

func main() {

	var text = "super text awesome super duper map inside stack overflow queue size go"
	var pairs = api.WordCount(text, 10)

	fmt.Println(pairs)

}
