package main

import (
	"github.com/ega-forever/otus_go/cmd"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {

	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	viper.SetDefault("LOG_LEVEL", 30)
	viper.SetDefault("RPC_URI", "127.0.0.1:8081")

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
