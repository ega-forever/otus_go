package main

import (
	"fmt"
	"otus/lesson5/workshop/api"
)

func main() {

	s := api.Stack{}
	s.Push(10)
	s.Push(20)
	s.Push(30)
	fmt.Printf("expected 30, got %d\n", s.Pop())
	fmt.Printf("expected 20, got %d\n", s.Pop())
	fmt.Printf("expected 10, got %d\n", s.Pop())

}
