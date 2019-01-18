package cmd

import (
	"fmt"
	"os"
	"os/user"

	"github.com/inabajunmr/treview/config"
	"github.com/inabajunmr/treview/filter"
	"github.com/inabajunmr/treview/github/trending"
	"github.com/jonboulle/clockwork"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "treview is cli viewer for GitHub Trending.",
	Run: func(cmd *cobra.Command, args []string) {
		usr, err := user.Current()
		if err != nil {
			os.Exit(1)
		}

		path := usr.HomeDir + "/.treview"

		l, err := cmd.Flags().GetString("lang")
		if err != nil {
			os.Exit(1)
		}

		var langs []string
		cpath := path + "/.config"
		if len(l) == 0 && exists(cpath) {
			// using default from conf
			langs = config.GetLangs(cpath)
		} else if l == "all" {
			// when exists config, I wanna see all
			langs = append(langs, "")
		} else {
			langs = append(langs, l)
		}

		s, err := cmd.Flags().GetString("span")
		if err != nil {
			os.Exit(1)
		}

		span := trending.GetSpanByString(s)

		// access to github
		c := make(chan []trending.Repository, len(langs))

		var repos []trending.Repository
		for _, lang := range langs {
			go func(l string) {
				findRepos, err := trending.FindTrending(l, span)
				if err != nil {
					println(err)
					os.Exit(1)
				}
				c <- findRepos
			}(lang)
		}

		for range langs {
			repos = append(repos, <-c...)
		}

		f, err := cmd.Flags().GetString("filter")
		if err != nil {
			os.Exit(1)
		}

		if f == "new" {
			// filter only new comer
			if err != nil {
				os.Exit(1)
			}
			f := filter.Filter{Time: clockwork.NewRealClock(), Path: path}
			repos = f.OnlyNewComer(repos)
		}

		// distinct by name
		distinctRepos := filter.DistinctRepository(repos)

		fmt.Println("■---------------------------------------------------------------------------■")
		for _, repo := range distinctRepos {
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
