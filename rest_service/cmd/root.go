package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "calendar",
	Short: "calendar service",
}

func init() {
	RootCmd.AddCommand(GrpcClientCmd)
	RootCmd.AddCommand(GrpcServerCmd)
}
