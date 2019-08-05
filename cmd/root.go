package main

import (
	"github.com/ega-forever/otus_go/internal"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
)

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetDefault("PORT", 8080)
	viper.SetDefault("LOG_LEVEL", 6)

	_ = viper.ReadInConfig()
	viper.AutomaticEnv()

	logLevel := viper.GetInt("LOG_LEVEL")

	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.Level(uint32(logLevel)))

}

func main() {

	port := viper.GetInt("PORT")
	log.Info(port)

	r := mux.NewRouter()
	r.Use(mux.CORSMethodMiddleware(r))
	r.Use(internal.LoggingMiddleware)

	internal.SetProductRouter(r)

	httpError := http.ListenAndServe(":"+strconv.Itoa(port), r)

	if httpError != nil {
		log.Panic(httpError)
	}

}
