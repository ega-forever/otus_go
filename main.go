package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
	"time"
)

func main() {

	ntpTime, err := ntp.Time("time.apple.com")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("network time %v / local time %v", ntpTime, time.Now())
}
