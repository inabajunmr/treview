package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"

	"github.com/ghodss/yaml"
	"github.com/inabajunmr/treview/filter"
	"github.com/inabajunmr/treview/github"
	"github.com/jonboulle/clockwork"
	"github.com/spf13/cobra"
)

type config struct {
	Lang []string `yaml:"lang"`
}

var rootCmd = &cobra.Command{
	Use: "treview is cli viewer for GitHub Trending.",
	Run: func(cmd *cobra.Command, args []string) {
		usr, err := user.Current()
		path := usr.HomeDir + "/.treview"

		l, err := cmd.Flags().GetString("lang")
		if err != nil {
			os.Exit(1)
		}

		var langs []string
		cpath := path + "/.config"
		if len(l) == 0 && exists(cpath) {
			// using default from conf
			langs = readConfig(cpath).Lang
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

		span := github.GetSpanByString(s)

		// access to github
		var repos []github.Repository
		for _, lang := range langs {
			findRepos, err := github.FindTrending(lang, span)
			if err != nil {
				println(err)
				os.Exit(1)
			}

			repos = append(repos, findRepos...)
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

func exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func readConfig(path string) config {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		print("Config is something wrong.")
		os.Exit(1)
	}

	var d config
	err = yaml.Unmarshal(buf, &d)
	if err != nil {
		print("Config is something wrong.")
		os.Exit(1)
	}

	return d
}
