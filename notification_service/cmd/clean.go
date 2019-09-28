package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

var CleanCmd = &cobra.Command{
	Use: "clean",
	Run: func(cmd *cobra.Command, args []string) {

		log.Println("going to drop all records from queue")

		fmt.Println("done")

	},
}
