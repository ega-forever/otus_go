package main

import (
	"fmt"
	"otus/lesson4/task1/api"
	"reflect"
)

func main() {

	elems := api.Container{12, 32, 21321, 43, 87}

	max := api.Max(func(a interface{}, b interface{}) bool {
		valA := reflect.ValueOf(a).Int()
		valB := reflect.ValueOf(b).Int()
		return valA > valB

	}, elems)

	fmt.Println(max)

}
