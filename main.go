package main

import (
	"os"
	"otus/api"
)

func main() {

	api.LogOtusEvent(api.HwAccepted{Id: 12, Grade: 15}, os.Stdout)
	api.LogOtusEvent(api.HwSubmitted{Id: 12, Code: "123", Comment: "super comment"}, os.Stdout)

}
