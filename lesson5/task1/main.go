package main

import (
	"fmt"
	"otus/lesson5/task1/queue"
)

func main() {

	superQueue := queue.Queue{}

	superQueue.Add(12)
	superQueue.Add(14)
	superQueue.Add(16)
	superQueue.Add(18)

	superQueue.Remove(0)

	fmt.Println(superQueue.Get(0).Value)
	fmt.Println(superQueue.Get(1).Value)
	fmt.Println(superQueue.Get(2).Value)
	fmt.Println("length:", superQueue.Len())

	iterator := superQueue.GetIterator()

	fmt.Println("iterating up")
	fmt.Println(iterator.GetCurrent().Value)
	for iterator.HasNext() {
		fmt.Println(iterator.GetNext().Value)
	}

	fmt.Println("iterating down")
	fmt.Println(iterator.GetCurrent().Value)
	for iterator.HasPrev() {
		fmt.Println(iterator.GetPrev().Value)
	}

}
