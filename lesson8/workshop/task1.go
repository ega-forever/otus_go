package main

import (
	"fmt"
	"sync"
	"time"
)

type Dog struct {
	name         string
	walkDuration time.Duration
}

func (d Dog) Walk(wg *sync.WaitGroup) {
	fmt.Printf("%s is taking a walk\n", d.name)
	time.Sleep(d.walkDuration)
	fmt.Printf("%s is going home\n", d.name)
	wg.Done()
}
func walkTheDogs(dogs []Dog) {

	var wg sync.WaitGroup
	for _, d := range dogs {
		wg.Add(1)
		go d.Walk(&wg)
	}
	wg.Wait()
	fmt.Println("everybody's home")
}
func main() {
	dogs := []Dog{{"vasya", time.Second}, {"john", time.Second}}
	walkTheDogs(dogs)
}
