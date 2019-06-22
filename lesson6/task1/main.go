package main

import (
	"os"
	"otus/lesson6/task1/api"
)

func main() {

	api.LogOtusEvent(api.HwAccepted{Id: 12, Grade: 15}, os.Stdout)
	api.LogOtusEvent(api.HwSubmitted{Id: 12, Code: "123", Comment: "super comment"}, os.Stdout)

}
