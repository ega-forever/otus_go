package cmd

import (
	"fmt"
	"github.com/ega-forever/otus_go/rest_service/internal/domain/services"
	gServer "github.com/ega-forever/otus_go/rest_service/internal/protocol/grpc"
	"github.com/ega-forever/otus_go/rest_service/internal/protocol/grpc/api"
	"github.com/ega-forever/otus_go/rest_service/internal/storage/memory"
	"github.com/ega-forever/otus_go/rest_service/internal/storage/sql"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
)

var GrpcServerCmd = &cobra.Command{
	Use: "grpc_server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("please specify adapter memory|database")
	},
}

var grpcServerMemory = &cobra.Command{
	Use: "memory",
	Run: func(cmd *cobra.Command, args []string) {
		rpcUri := viper.GetString("RPC_URI")
		listen, err := net.Listen("tcp", rpcUri)

		if err != nil {
			log.Fatalf(err.Error())
		}

		grpcServer := grpc.NewServer()
		storage := memory.New()
		service := services.NewEventService(storage)
		srv := gServer.New(service)

		event.RegisterEventServiceServer(grpcServer, srv)

		log.Info("hosting server on: ", listen.Addr().String())
		if err := grpcServer.Serve(listen); err != nil {
			log.Fatalf("failed to serve: %s", err)
		}
	},
}

var grpcServerDB = &cobra.Command{
	Use: "database",
	Run: func(cmd *cobra.Command, args []string) {
		rpcUri := viper.GetString("RPC_URI")
		dbUri := viper.GetString("DB_URI")

		listen, err := net.Listen("tcp", rpcUri)

		if err != nil {
			log.Fatalf(err.Error())
		}

		grpcServer := grpc.NewServer()
		storage := db.New(dbUri)
		err = storage.Migrate()

		if err != nil {
			log.Fatal(err)
		}

		service := services.NewEventService(storage)
		srv := gServer.New(service)

		event.RegisterEventServiceServer(grpcServer, srv)

		log.Info("hosting server on: ", listen.Addr().String())
		if err := grpcServer.Serve(listen); err != nil {
			log.Fatalf("failed to serve: %s", err)
		}
	},
}

func init() {
	GrpcServerCmd.AddCommand(grpcServerMemory)
	GrpcServerCmd.AddCommand(grpcServerDB)
}
