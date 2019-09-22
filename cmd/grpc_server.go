package cmd

import (
	"github.com/ega-forever/otus_go/internal/domain/services"
	gServer "github.com/ega-forever/otus_go/internal/protocol/grpc"
	"github.com/ega-forever/otus_go/internal/protocol/grpc/api"
	"github.com/ega-forever/otus_go/internal/storage/memory"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
)

var GrpcServerCmd = &cobra.Command{
	Use: "grpc_server",
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
