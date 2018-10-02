package cmd

import (
	"fmt"
	"os"

	"github.com/inabajunmr/treview/github"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "treview is cli viewer for GitHub Trending.",
	Run: func(cmd *cobra.Command, args []string) {
		l, err := cmd.Flags().GetString("lang")
		if err != nil {
			os.Exit(1)
		}

		s, err := cmd.Flags().GetString("span")
		if err != nil {
			os.Exit(1)
		}

		span := github.GetSpanByString(s)

		repos, err := github.Find(l, span)
		if err != nil {
			println(err)
			os.Exit(1)
		}

		for _, repo := range repos {
			fmt.Println("------------------------")
			repo.Print()
		}
	},
}

// Execute cmd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("lang", "l", "", "filter by lang")
	rootCmd.Flags().StringP("span", "s", "Today", "trending span")
}
