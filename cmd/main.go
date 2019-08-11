package main

import (
	"github.com/ega-forever/otus_go/internal/app"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
)

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetDefault("PORT", 8080)
	viper.SetDefault("LOG_LEVEL", 6)

	_ = viper.ReadInConfig()
	viper.AutomaticEnv()

	logLevel := viper.GetString("LOG_LEVEL")

	log.SetFormatter(&log.JSONFormatter{})

	parsedLevel, err := log.ParseLevel(logLevel)

	if err != nil {
		log.Panic(err)
	}

	log.SetLevel(parsedLevel)

}

func main() {

	port := viper.GetString("PORT")
	log.Info(port)

	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(app.LoggingMiddleware)

	app.SetProductRouter(r)

	err := http.ListenAndServe(":"+port, r)

	if err != nil {
		log.Panic(err)
	}

}
