package main

import (
	"fmt"
	"otus/lesson1/task3/api"
)

func main() {
	var str = "https://gobyexample.com/url-parsing"

	var holder = api.LinkHolder{Links: make(map[string]string)}

	shortUrl := api.Shorten(holder, str)
	fmt.Println(shortUrl)

	originalUrl := api.Resolve(holder, shortUrl)

	fmt.Println(originalUrl)

}
