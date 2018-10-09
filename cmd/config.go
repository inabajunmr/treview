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
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
