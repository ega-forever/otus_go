package main

import (
	"fmt"
	"otus/lesson4/task1/api"
	"strconv"
)

func main() {

	elems := api.Container{12, 32, 21321, 43, "", false}

	max := api.Max(func(a interface{}, b interface{}) bool {

		aStr := fmt.Sprintf("%d", a)
		valA, _ := strconv.ParseInt(aStr, 10, 32)

		bStr := fmt.Sprintf("%d", b)
		valB, _ := strconv.ParseInt(bStr, 10, 32)

		return valA > valB

	}, elems)

	fmt.Println(max)

}
