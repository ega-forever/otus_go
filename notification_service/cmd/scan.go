package cmd

import (
	"github.com/ega-forever/otus_go/notification_service/internal/queue"
	"github.com/ega-forever/otus_go/notification_service/internal/services"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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

		scanServiceInstance := services.NewScanService(queueConn)
		scanServiceInstance.Job()

		restPort := viper.GetInt("REST_PORT")

		restServiceInstance := services.NewRestService(restPort)
		err = restServiceInstance.Start()

		if err != nil {
			log.Fatal(err)
		}

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
