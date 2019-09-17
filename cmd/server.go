package main

import (
	"fmt"
	"github.com/ega-forever/otus_go/internal/server"
	"github.com/ega-forever/otus_go/internal/shared"
	"github.com/spf13/cobra"
	"log"
	"sync"
)

func getServerCmd() *cobra.Command {

	var port int
	var address string

	serverCmd := &cobra.Command{
		Use:   "server",
		Short: "the telnet server",
		Long:  `the telnet server`,
		Run: func(cmd *cobra.Command, args []string) {

			serverInstance, err := server.New(address, port)

			if err != nil {
				log.Fatal(err)
			}

			serverInstance.Listen()

			wg := sync.WaitGroup{}
			wg.Add(1)
			exitChan := shared.ListenExitSignal()

			go func() {
				for {
					select {
					case c := <-serverInstance.MsgC:
						fmt.Printf("Message from %s: %s\n", c.Address, c.Msg)
					case <-exitChan:
						{
							wg.Done()
							_ = serverInstance.Listener.Close()
						}
					}
				}
			}()
			wg.Wait()
		},
	}

	serverCmd.Flags().IntVarP(&port, "port", "p", 2100, "server port")
	serverCmd.Flags().StringVarP(&address, "address", "a", "0.0.0.0", "server address")

	return serverCmd

}
