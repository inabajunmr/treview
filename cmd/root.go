package cmd

import (
	"fmt"
	"os"
	"os/user"

	"github.com/inabajunmr/treview/filter"
	"github.com/inabajunmr/treview/github"
	"github.com/jonboulle/clockwork"
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

		// access to github
		repos, err := github.FindTrending(l, span)
		if err != nil {
			println(err)
			os.Exit(1)
		}

		f, err := cmd.Flags().GetString("filter")
		if err != nil {
			os.Exit(1)
		}

		if f == "new" {
			// filter only new comer
			usr, err := user.Current()
			if err != nil {
				os.Exit(1)
			}
			path := usr.HomeDir + "/.treview"
			f := filter.Filter{Time: clockwork.NewRealClock(), Path: path}
			repos = f.OnlyNewComer(repos)
		}

		fmt.Println("■---------------------------------------------------------------------------■")
		for _, repo := range repos {
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
