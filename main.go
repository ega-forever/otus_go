package main

import (
	"fmt"
)

func iota(val int) string {

	d := "0123456789abcdef"
	var str = ""

	for i := 0; val > 0 && i < 30; i++ {
		ch := d[val%10]
		str = string(ch) + str
		val /= 10
	}

	return str
}

func main() {

	str := iota(12888)
	fmt.Println(str)

}
