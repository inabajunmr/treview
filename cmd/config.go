package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Interactive setting for default langage configration.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config called")

		// TODO
		// check config file
		// if exist
		// load and modify and write
		// if not exist
		// create file and write config
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
