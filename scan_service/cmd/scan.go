package cmd

import (
	"github.com/ega-forever/otus_go/scan_service/internal/queue"
	"github.com/ega-forever/otus_go/scan_service/internal/services"
	db "github.com/ega-forever/otus_go/scan_service/internal/storage/sql"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

var ScanCmd = &cobra.Command{
	Use: "scan",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 2 {
			log.Fatal("specify scan_period earliest period")
		}

		scanPeriodSeconds, err := strconv.Atoi(args[0])

		if err != nil {
			log.Fatal(err)
		}

		eventCreatedEarliestSeconds, err := strconv.Atoi(args[1])

		if err != nil {
			log.Fatal(err)
		}

		dbUri := viper.GetString("DB_URI")
		dbConn := db.New(dbUri)

		err = dbConn.Migrate()

		if err != nil {
			log.Fatal(err)
		}

		queueUri := viper.GetString("QUEUE_URI")
		queueConn, err := queue.New(queueUri)

		if err != nil {
			log.Fatal(err)
		}

		restPort := viper.GetInt("REST_PORT")

		scanServiceInstance := services.NewScanService(dbConn, queueConn)
		scanServiceInstance.Job(time.Duration(scanPeriodSeconds), int64(eventCreatedEarliestSeconds))

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
					_ = restServiceInstance.Disconnect()
				case <-sigs:
					wg.Done()
					_ = restServiceInstance.Disconnect()
				}
			}
		}()

		wg.Wait()
	},
}
