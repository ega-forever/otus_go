package main

import (
	"bufio"
	"github.com/ega-forever/otus_go/internal/client"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func getClientCmd() *cobra.Command {

	var port int
	var address string

	clientCmd := &cobra.Command{
		Use:   "client",
		Short: "the telnet client",
		Long:  `the telnet client`,
		Run: func(cmd *cobra.Command, args []string) {

			wg := sync.WaitGroup{}

			closed := false

			wg.Add(1)
			sigs := make(chan os.Signal, 1)

			signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
			signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

			connection := client.StartClient(address, port)

			go func() {
				<-sigs
				closed = true
				connection.Close()
				wg.Done()
			}()

			go func() {
				reader := bufio.NewReader(os.Stdin)
				for {
					text, _ := reader.ReadBytes('\n')
					_, _ = connection.Write(text)
				}
			}()

			wg.Wait()

		},
	}

	clientCmd.Flags().IntVarP(&port, "port", "p", 2100, "server port")
	clientCmd.Flags().StringVarP(&address, "address", "a", "0.0.0.0", "server address")

	return clientCmd

}
