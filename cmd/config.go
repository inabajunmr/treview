package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Setting for default langage configration.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config called")
		os.Exit(1)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
