package main

import (
	"bufio"
	"fmt"
	"github.com/ega-forever/otus_go/internal/client"
	"github.com/ega-forever/otus_go/internal/shared"
	"github.com/spf13/cobra"
	"log"
	"os"
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

			clientInstance, err := client.New(address, port, 10000)

			if err != nil {
				log.Fatalf(err.Error())
			}

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
					case <-clientInstance.Ctx.Done():
						{

							err := clientInstance.Ctx.Err()
							if err != nil {
								fmt.Println("server is down")
								exitChan <- syscall.SIGTERM
							}
						}
					}
				}
			}()

			go func() {
				reader := bufio.NewReader(os.Stdin)

				for {
					text, _ := reader.ReadBytes('\n')
					_, err := clientInstance.Connection.Write(text)

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
