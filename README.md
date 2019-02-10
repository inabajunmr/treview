# treview

[![codecov](https://codecov.io/gh/inabajunmr/treview/branch/master/graph/badge.svg)](https://codecov.io/gh/inabajunmr/treview)
![CircleCI](https://circleci.com/gh/inabajunmr/treview/tree/master.svg?style=svg)

Viewer for GitHub Trending.

## CLI
![treview](https://user-images.githubusercontent.com/10000393/46803178-6ea89600-cd9a-11e8-8ca6-ad52c58ed942.gif)

## GUI
![treview](https://user-images.githubusercontent.com/10000393/52535598-bdd7f500-2d93-11e9-9f61-82835453a863.png)



# Feature

* shows only today's newcomer repository.(You never see same project everyday!)
* define default langage that you want to see.
* [GUI for macos](https://github.com/inabajunmr/treview/releases)

# Install
## CLI
```
$ go get github.com/inabajunmr/treview
```

## GUI
Download from [GUI for macos](https://github.com/inabajunmr/treview/releases).

# Usage
```
$ treview -h
Usage:
  treview is cli viewer for GitHub Trending. [flags]
  treview [command]

Available Commands:
  config      Setting for default langage configration.
  help        Help about any command

Flags:
  -f, --filter string   all or new (default "new")
  -h, --help            help for treview
  -l, --lang string     filter by lang
  -s, --span string     trending span (default "Today")

Use "treview [command] --help" for more information about a command.
```

## Config
If you have file `~/.treview/.config`, you can set default lang by treview.
You can set config by `treview config` command too.

![treview](https://user-images.githubusercontent.com/10000393/46802798-656af980-cd99-11e8-88fb-a91a72fbfcfd.gif)

### Example
If you have following config, treview show only Golang and JavaScript repositories by `treview` command (without lang flag).
```yaml
lang:  [go, javascript]
```

If you want to find all langage, you set `all` as lang flag.

# Appendix (Using as GitHub Trending API for Golang)
[![GoDoc](https://godoc.org/github.com/inabajunmr/treview/github?status.svg)](https://godoc.org/github.com/inabajunmr/treview/github)

## Sample
```go
package main

import (
	"fmt"
	"github.com/inabajunmr/treview/github/trending"
)

func main() {
	span := trending.GetSpanByString("today")

	repos, err := trending.FindTrending(l, span)
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
