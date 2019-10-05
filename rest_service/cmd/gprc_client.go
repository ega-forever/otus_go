package cmd

import (
	"context"
	"github.com/ega-forever/otus_go/rest_service/internal/protocol/grpc/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
	"strconv"
)

var GrpcClientCmd = &cobra.Command{
	Use: "grpc_client",
	Run: func(cmd *cobra.Command, args []string) {
		// todo implement
	},
}

var createEventCmd = &cobra.Command{
	Use: "create",
	Run: func(cmd *cobra.Command, args []string) {

		rpcUri := viper.GetString("RPC_URI")

		conn, err := grpc.Dial(rpcUri, grpc.WithInsecure())
		if err != nil {
			log.Fatal(err)
		}

		client := event.NewEventServiceClient(conn)

		timestamp, err := strconv.Atoi(args[1])

		if err != nil {
			log.Fatal(err)
		}

		req := &event.CreateEventReq{
			Event: &event.Event{Text: args[0], Timestamp: int64(timestamp)},
		}

		resp, err := client.CreateEvent(context.Background(), req)

		if err != nil {
			log.Fatal(err)
		}

		if resp.GetError() != "" {
			log.Fatal(resp)
		} else {
			log.Println(resp.GetResult())
		}
	},
}

var updateEventCmd = &cobra.Command{
	Use: "update",
	Run: func(cmd *cobra.Command, args []string) {

		rpcUri := viper.GetString("RPC_URI")

		conn, err := grpc.Dial(rpcUri, grpc.WithInsecure())
		if err != nil {
			log.Fatal(err)
		}

		client := event.NewEventServiceClient(conn)

		id, err := strconv.Atoi(args[0])

		if err != nil {
			log.Fatal(err)
		}

		timestamp, err := strconv.Atoi(args[2])

		if err != nil {
			log.Fatal(err)
		}

		req := &event.UpdateEventReq{
			Event: &event.Event{Id: int64(id), Text: args[1], Timestamp: int64(timestamp)},
		}

		resp, err := client.UpdateEvent(context.Background(), req)

		if err != nil {
			log.Fatal(err)
		}

		if resp.GetError() != "" {
			log.Fatal(resp)
		} else {
			log.Println(resp.GetResult())
		}
	},
}

var getEventCmd = &cobra.Command{
	Use: "get",
	Run: func(cmd *cobra.Command, args []string) {

		rpcUri := viper.GetString("RPC_URI")

		conn, err := grpc.Dial(rpcUri, grpc.WithInsecure())
		if err != nil {
			log.Fatal(err)
		}

		client := event.NewEventServiceClient(conn)

		id, err := strconv.Atoi(args[0])

		if err != nil {
			log.Fatal(err)
		}

		req := &event.GetEventReq{
			Id: int64(id),
		}

		resp, err := client.GetEvent(context.Background(), req)

		if err != nil {
			log.Fatal(err)
		}

		if resp.GetError() != "" {
			log.Fatal(resp)
		} else {
			log.Println(resp.GetResult())
		}
	},
}

var deleteEventCmd = &cobra.Command{
	Use: "delete",
	Run: func(cmd *cobra.Command, args []string) {

		rpcUri := viper.GetString("RPC_URI")

		conn, err := grpc.Dial(rpcUri, grpc.WithInsecure())
		if err != nil {
			log.Fatal(err)
		}

		client := event.NewEventServiceClient(conn)

		id, err := strconv.Atoi(args[0])

		if err != nil {
			log.Fatal(err)
		}

		req := &event.DeleteEventReq{
			Id: int64(id),
		}

		resp, err := client.DeleteEvent(context.Background(), req)

		if err != nil {
			log.Fatal(err)
		}

		if resp.GetError() != "" {
			log.Fatal(resp)
		} else {
			log.Println(resp.GetResult())
		}
	},
}

var listEventsCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {

		rpcUri := viper.GetString("RPC_URI")

		conn, err := grpc.Dial(rpcUri, grpc.WithInsecure())
		if err != nil {
			log.Fatal(err)
		}

		client := event.NewEventServiceClient(conn)

		req := &event.ListEventReq{}

		resp, err := client.ListEvents(context.Background(), req)

		if err != nil {
			log.Fatal(err)
		}

		if resp.GetError() != "" {
			log.Fatal(resp)
		} else {
			log.Println(resp.GetResult())
		}
	},
}

func init() {
	GrpcClientCmd.AddCommand(createEventCmd)
	GrpcClientCmd.AddCommand(updateEventCmd)
	GrpcClientCmd.AddCommand(getEventCmd)
	GrpcClientCmd.AddCommand(deleteEventCmd)
	GrpcClientCmd.AddCommand(listEventsCmd)
}
