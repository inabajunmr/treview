package cmd

import (
	"fmt"
	"os"

	"github.com/inabajunmr/treview/github/trending"
	treview "github.com/inabajunmr/treview/service"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "treview is cli viewer for GitHub Trending.",
	Run: func(cmd *cobra.Command, args []string) {

		l, err := cmd.Flags().GetString("lang")
		if err != nil {
			os.Exit(1)
		}

		langs := treview.GetLangs(l)

		s, err := cmd.Flags().GetString("span")
		if err != nil {
			os.Exit(1)
		}

		span := trending.GetSpanByString(s)

		f, err := cmd.Flags().GetString("filter")
		if err != nil {
			os.Exit(1)
		}

		isOnlyNew := false
		if f == "new" {
			// filter only new comer
			if err != nil {
				os.Exit(1)
			}

			isOnlyNew = true
		}

		fmt.Println("■---------------------------------------------------------------------------■")
		for _, repo := range treview.GetRepositories(span, langs, isOnlyNew) {
			repo.Print()
			fmt.Println("■---------------------------------------------------------------------------■")
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
	rootCmd.Flags().StringP("filter", "f", "new", "all or new")
}

func exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
