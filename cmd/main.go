package main

import (
	"github.com/ega-forever/otus_go/internal"
	eventpb "github.com/ega-forever/otus_go/proto"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
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

	rpcUri := viper.GetString("RPC_URI")
	listen, err := net.Listen("tcp", rpcUri)

	if err != nil {
		log.Fatalf(err.Error())
	}

	grpcServer := grpc.NewServer()
	srv := internal.NewEventService()

	eventpb.RegisterEventServiceServer(grpcServer, srv)

	log.Info("hosting server on: ", listen.Addr().String())
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
