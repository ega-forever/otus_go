package shared

import (
	"os"
	"os/signal"
	"syscall"
)

func ListenExitSignal() chan os.Signal {

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	return sigs
}
