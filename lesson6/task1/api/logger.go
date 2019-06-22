package api

import (
	"fmt"
	"io"
	"time"
)

type HwAccepted struct {
	Id    int
	Grade int
}

type HwSubmitted struct {
	Id      int
	Code    string
	Comment string
}

type OtusEvent interface {
	getLog() string
}

func (log HwAccepted) getLog() string {

	currentTime := time.Now()

	return fmt.Sprintf(
		"%d-%d-%d accepted %d %d",
		currentTime.Year(),
		currentTime.Month(),
		currentTime.Day(),
		log.Id,
		log.Grade)
}

func (log HwSubmitted) getLog() string {

	currentTime := time.Now()

	return fmt.Sprintf(
		"%d-%d-%d submitted %d %s",
		currentTime.Year(),
		currentTime.Month(),
		currentTime.Day(),
		log.Id,
		log.Comment)
}

func LogOtusEvent(e OtusEvent, w io.Writer) {

	log := e.getLog() + "\n"

	_, _ = w.Write([]byte(log))

}
