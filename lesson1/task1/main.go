package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"time"
)

func main() {

	ntpTime, err := ntp.Time("time.apple.com")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("network time %v / local time %v", ntpTime, time.Now())
}
