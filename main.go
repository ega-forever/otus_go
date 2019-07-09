package main

import (
	"flag"
	"github.com/ega-forever/otus_go/api"
)

var from string
var to string
var offset int
var limit int

func init() {
	flag.StringVar(&from, "from", "", "source file")
	flag.StringVar(&to, "to", "", "destination file")
	flag.IntVar(&offset, "offset", 0, "offset")
	flag.IntVar(&limit, "limit", 0, "limit")
	flag.Parse()
}

func main() {
	api.Copy(from, to, offset, limit)
}
