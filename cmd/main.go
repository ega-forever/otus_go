package main

import (
	// log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(getClientCmd())
	rootCmd.AddCommand(getServerCmd())
}

var rootCmd = &cobra.Command{
	Use:   "telnet",
	Short: "super telnet client",
	Long: `
	start as server: telnet server --port=2000 --address=0.0.0.0
	start as client: telnet client --port=2000 --address=0.0.0.0
    `, // todo
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func main() {

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
