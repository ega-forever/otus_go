package cmd

import (
	"fmt"
	"github.com/ega-forever/otus_go/scan_service/internal/domain/models"
	db "github.com/ega-forever/otus_go/scan_service/internal/storage/sql"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"math/rand"
	"strconv"
	"time"
)

var FillCmd = &cobra.Command{
	Use: "fill",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			log.Fatal("please specify amount of records")
		}

		fmt.Printf("going to fill the db with %s records", args[0])
		amount, err := strconv.Atoi(args[0])

		if err != nil {
			log.Fatal("please specify amount of records")
		}

		dbUri := viper.GetString("DB_URI")
		conn := db.New(dbUri)

		err = conn.Migrate()

		if err != nil {
			log.Fatal(err)
		}

		timestampMin := time.Now().Unix() - 86400*1000
		timestampMax := time.Now().Unix()
		for i := 0; i < amount; i++ {
			event := models.Event{Text: "random", Timestamp: rand.Int63n(timestampMax-timestampMin) + timestampMin}
			_, err := conn.SaveEvent(&event)
			if err != nil {
				log.Fatal(err)
			}

		}

		fmt.Println("done")

	},
}
