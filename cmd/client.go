package main

import (
	"bufio"
	"github.com/ega-forever/otus_go/internal/client"
	"github.com/ega-forever/otus_go/internal/shared"
	"github.com/spf13/cobra"
	"log"
	"os"
	"sync"
)

func getClientCmd() *cobra.Command {

	var port int
	var address string

	clientCmd := &cobra.Command{
		Use:   "client",
		Short: "the telnet client",
		Long:  `the telnet client`,
		Run: func(cmd *cobra.Command, args []string) {

			clientInstance, err := client.New(address, port)

			if err != nil {
				log.Fatalf(err.Error())
			}

			clientInstance.Ping(2000)

			wg := sync.WaitGroup{}
			wg.Add(1)
			exitChan := shared.ListenExitSignal()

			go func() {
				for {
					select {
					case <-exitChan:
						{
							_ = clientInstance.Connection.Close()
							wg.Done()
						}
					case err := <-clientInstance.ErrC:
						log.Println(err)
						_ = clientInstance.Connection.Close()
						wg.Done()
					}
				}
			}()

			go func() {
				reader := bufio.NewReader(os.Stdin)

				for {
					text, _ := reader.ReadBytes('\n')
					_, err := clientInstance.Connection.Write([]byte(string(text) + "\n"))

					if err != nil {
						log.Fatal(err)
					}

				}
			}()

			wg.Wait()

		},
	}

	clientCmd.Flags().IntVarP(&port, "port", "p", 2100, "server port")
	clientCmd.Flags().StringVarP(&address, "address", "a", "0.0.0.0", "server address")

	return clientCmd

}
