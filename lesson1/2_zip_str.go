package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	str := os.Args[1]
	var finalStr string
	inc := 1

	for i := 0; i < len(str); i++ {

		if i+1 < len(str) && str[i] == str[i+1] {
			inc++
		} else if inc > 1 {
			finalStr += string(str[i-1]) + strconv.Itoa(inc)
			inc = 1
		} else {
			finalStr += string(str[i])
		}

	}

	fmt.Print(finalStr)

}
