package main

import (
	"fmt"
	"github.com/ega-forever/otus_go/internal/server"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func getServerCmd() *cobra.Command {

	var port int
	var address string

	serverCmd := &cobra.Command{
		Use:   "server",
		Short: "the telnet server",
		Long:  `the telnet server`,
		Run: func(cmd *cobra.Command, args []string) {

			wg := sync.WaitGroup{}

			closed := false

			wg.Add(1)
			sigs := make(chan os.Signal, 1)

			signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
			signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

			listener := server.StartServer(address, port)

			go func() {
				<-sigs
				closed = true
				listener.Close()
				wg.Done()
			}()

			go func() {
				msg := make([]byte, 1024, 1024)
				for {
					length, fromAddr, err := listener.ReadFromUDP(msg)

					if closed {
						return
					}

					if err != nil {
						log.Fatalf("Error emitted %s", err)
					}
					fmt.Printf("Message from %s with length %d: %s\n", fromAddr.String(), length, string(msg))
				}
			}()

			wg.Wait()
		},
	}

	serverCmd.Flags().IntVarP(&port, "port", "p", 2100, "server port")
	serverCmd.Flags().StringVarP(&address, "address", "a", "0.0.0.0", "server address")

	return serverCmd

}
