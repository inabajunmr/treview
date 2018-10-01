package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hoge",
	Short: "hoge hoge",
	Long:  `hoge hoge hoge hoge`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hogehogehogehoge")
	},
}

// Execute cmd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
