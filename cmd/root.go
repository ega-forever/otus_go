package cmd

import (
	"errors"
	"fmt"
	"github.com/ega-forever/otus_go/internal"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "otus_go",
	Short: "Run a program with custom env [env_dir_path] [program_to_run]",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("not enough args")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		env, scanErr := internal.ScanEnv(args[0])

		if scanErr != nil {
			log.Fatal(scanErr)
		}

		runErr := internal.RunProgram(args[1], env)

		if runErr != nil {
			log.Fatal(scanErr)
		}

	},
}

func Execute() {
	err := rootCmd.Execute()
	// fmt.Println(cfgFile)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
