package treview

import (
	"os"
	"os/user"

	"github.com/inabajunmr/treview/filter"
	"github.com/inabajunmr/treview/github/trending"
	"github.com/jonboulle/clockwork"
)

func GetRepositories(span trending.Span, langs []string, isOnlyNew bool) []trending.Repository {
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

	if isOnlyNew {
		usr, err := user.Current()
		if err != nil {
			os.Exit(1)
		}
		path := usr.HomeDir + "/.treview"
		f := filter.Filter{Time: clockwork.NewRealClock(), Path: path}
		repos = f.OnlyNewComer(repos)
	}

	// distinct by name
	return filter.DistinctRepository(repos)
}

func exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}
