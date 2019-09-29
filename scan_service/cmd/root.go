package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "scan",
	Short: "scan service",
}

func init() {
	RootCmd.AddCommand(FillCmd)
	RootCmd.AddCommand(ScanCmd)
}
