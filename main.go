package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	str := os.Args[1]
	var finalStrBuilder strings.Builder
	inc := 1

	for i := 0; i < len(str); i++ {

		if i+1 < len(str) && str[i] == str[i+1] {
			inc++
		} else if inc > 1 {
			finalStrBuilder.WriteString(string(str[i-1]))
			finalStrBuilder.WriteString(strconv.Itoa(inc))
			inc = 1
		} else {
			finalStrBuilder.WriteString(string(str[i]))
		}

	}

	fmt.Print(finalStrBuilder.String())

}
