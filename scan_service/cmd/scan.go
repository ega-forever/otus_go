package cmd

import (
	"github.com/ega-forever/otus_go/scan_service/internal/queue"
	"github.com/ega-forever/otus_go/scan_service/internal/services"
	db "github.com/ega-forever/otus_go/scan_service/internal/storage/sql"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
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
		// todo implement

		if len(args) < 2 {
			log.Fatal("specify scan_period earliest_timestamp")
		}

		seconds, err := strconv.Atoi(args[0])

		if err != nil {
			log.Fatal(err)
		}

		timestamp, err := strconv.Atoi(args[1])

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

		scanServiceInstance := services.New(dbConn, queueConn)
		scanServiceInstance.Job(time.Duration(seconds), int64(timestamp))

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
