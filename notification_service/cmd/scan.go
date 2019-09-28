package cmd

import (
	"github.com/ega-forever/otus_go/notification_service/internal/queue"
	"github.com/ega-forever/otus_go/notification_service/internal/services"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var ScanCmd = &cobra.Command{
	Use: "scan",
	Run: func(cmd *cobra.Command, args []string) {

		queueUri := viper.GetString("QUEUE_URI")
		queueConn, err := queue.New(queueUri)

		if err != nil {
			log.Fatal(err)
		}

		scanServiceInstance := services.New(queueConn)
		scanServiceInstance.Job()

		sigs := make(chan os.Signal, 1)

		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

		wg := sync.WaitGroup{}
		wg.Add(1)

		go func() {
			for {
				select {
				case <-scanServiceInstance.Ctx.Done():
					wg.Done()
				case <-sigs:
					wg.Done()
				}
			}
		}()

		wg.Wait()
	},
}
