package main

import (
	"fmt"
	"os"
	"otus/api"
)

func main() {

	str := os.Args[1]

	holder := api.LinkHolder{Links: make(map[string]string)}

	shortUrl := api.Shorten(holder, str)
	fmt.Println("short url:", shortUrl)

	originalUrl := api.Resolve(holder, shortUrl)

	fmt.Println("original url:", originalUrl)

}
