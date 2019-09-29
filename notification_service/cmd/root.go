package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "notification",
	Short: "notification service",
}

func init() {
	RootCmd.AddCommand(ScanCmd)
}
