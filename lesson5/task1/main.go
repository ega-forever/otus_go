package main

import (
	"fmt"
	"otus/lesson5/task1/api"
)

func main() {

	queue := api.Queue{}

	queue.Add(12)
	queue.Add(14)
	queue.Add(16)
	queue.Add(18)

	queue.Remove(0)

	fmt.Println(queue.Get(0).Value)
	fmt.Println(queue.Get(1).Value)
	fmt.Println(queue.Get(2).Value)
	fmt.Println(queue.Len())

	iterator := queue.GetIterator()
	fmt.Println(iterator.GetCurrent().Value)
	fmt.Println(iterator.GetNext().Value)
	fmt.Println(iterator.HasNext())

}
