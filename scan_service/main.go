package main

import (
	"github.com/ega-forever/otus_go/scan_service/cmd"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {

	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	viper.SetDefault("LOG_LEVEL", 30)
	viper.SetDefault("QUEUE_URI", "amqp://guest:guest@localhost:5672")
	viper.SetDefault("DB_URI", "postgres://user:123@localhost:5432/otus")

	viper.ReadInConfig()
	viper.AutomaticEnv()

	logLevel := viper.GetInt("LOG_LEVEL")
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.Level(logLevel))
}

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
