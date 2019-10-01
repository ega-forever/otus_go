package main

import (
	"github.com/ega-forever/otus_go/notification_service/cmd"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {

	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	viper.SetDefault("REST_PORT", 8080)
	viper.SetDefault("LOG_LEVEL", 30)
	viper.SetDefault("QUEUE_URI", "amqp://guest:guest@localhost:5672")

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
