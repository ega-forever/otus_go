package main

import (
	"fmt"
	"reflect"
	"time"
)

func MergeChans(ins ...<-chan interface{}) <-chan interface{} {

	multiCh := make(chan interface{})

	cases := make([]reflect.SelectCase, len(ins))
	for i, ch := range ins {
		cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch)}
	}

	go func() {
		remaining := len(cases)
		for remaining > 0 {
			chosen, value, ok := reflect.Select(cases)
			if !ok {
				cases[chosen].Chan = reflect.ValueOf(nil)
				remaining -= 1
				continue
			}

			multiCh <- value
		}
	}()

	return multiCh
}

func main() {

	in1 := make(chan interface{}, 10)
	in2 := make(chan interface{}, 10)

	out1 := MergeChans(in1, in2)

	go func() {

		for data := range out1 {
			fmt.Println(data)
		}

	}()

	for i := 0; i < 100; i++ {
		in1 <- i
		in2 <- i
	}

	time.Sleep(1 * time.Second)

}
