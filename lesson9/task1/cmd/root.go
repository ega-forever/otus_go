package cmd

import (
	"fmt"
	"otus/lesson9/task1/api"
	"strconv"
)

func Execute(command string, arg ...string) {

	switch command {
	case "words":
		n, _ := strconv.ParseInt(arg[1], 10, 32)
		res := api.WordCount(arg[0], int(n))
		fmt.Println(res)
	case "short":
		holder := api.LinkHolder{Links: make(map[string]string)}
		shortUrl := api.Shorten(holder, arg[0])
		fmt.Println(shortUrl)
	default:
		fmt.Println("command not found")

	}

}
