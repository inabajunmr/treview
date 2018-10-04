# treview

[![codecov](https://codecov.io/gh/inabajunmr/treview/branch/master/graph/badge.svg)](https://codecov.io/gh/inabajunmr/treview)
[![CircleCI](https://circleci.com/gh/inabajunmr/treview/tree/master.svg?style=svg)]

treview is viewer for GitHub Trending.

# Feature

treview shows only today's newcomer repository.

# Install
```
$ go get github.com/inabajunmr/treview
```

# Usage
```
$ treview -h
Usage:
  treview is cli viewer for GitHub Trending. [flags]

Flags:
  -f, --filter string   all or new (default "new")
  -h, --help            help for treview
  -l, --lang string     filter by lang
  -s, --span string     trending span (default "Today")
```

## Config
If you have file `~/.treview/.config`, you can set default lang by treview.

### Example
If you have follow config, treview show only Golang and JavaScript repositories by `treview` command (without lang flag).
```yaml
lang:  [go, javascript]
```

If you want to find all langage, you set `all` as lang flag.

# Example
```
$ treview -l go
■---------------------------------------------------------------------------■
【MontFerret / ferret】(https://github.com/MontFerret/ferret)
Lang:Go	Fork:32	⭐️1009	⭐️915 stars today
Declarative web scraping
■---------------------------------------------------------------------------■
【sourcegraph / sourcegraph】(https://github.com/sourcegraph/sourcegraph)
Lang:Go	Fork:68	⭐️1951	⭐️793 stars today
Code search and intelligence, self-hosted and scalable
■---------------------------------------------------------------------------■
【spiral / roadrunner】(https://github.com/spiral/roadrunner)
Lang:Go	Fork:43	⭐️1277	⭐️156 stars today
High-performance PHP application server, load-balancer and process manager written in Golang
■---------------------------------------------------------------------------■
```

# Appendix (Using as GitHub Trending API for Golang)
[![GoDoc](https://godoc.org/github.com/inabajunmr/treview/github?status.svg)](https://godoc.org/github.com/inabajunmr/treview/github)

## Sample
```go
package main

import (
	"fmt"
	"github.com/inabajunmr/treview/github"
)

func main() {
	span := github.GetSpanByString("today")

	repos, err := github.FindTrending(l, span)
	if err != nil {
		println(err)
		os.Exit(1)
	}

	for _, repo := range repos {
		fmt.Println("------------------------")
		repo.Print()
	}
}
```
