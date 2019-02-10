package trending

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Span is span for trending search
type Span int

const (
	// Today span
	Today Span = iota
	// Week span
	Week
	// Month span
	Month
)

// Repository is expression type for github repository on trending.
type Repository struct {
	Name        string
	URL         string
	Description string
	Lang        string
	Star        int
	StarBySpan  int
	Fork        int
}

// FindTrending repositories by lang and span
func FindTrending(lang string, span Span) ([]Repository, error) {
	// assemble url for trending
	url := "https://github.com/trending/" + lang + "?" + "since=" + getQueryForSpan(span)

	// access to github
	resp, err := http.Get(url)
	if err != nil {
		print("Can not access to " + url)
		print(err)
		os.Exit(1)
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		return nil, err
	}

	// correct repositories
	repos := make([]Repository, 25)
	doc.Find("div.explore-content > ol > li").Each(func(i int, s *goquery.Selection) {
		repos[i] = getRepositoryBySelection(s)
	})

	return repos, nil
}

// FindLangs from GitHub trending
func FindLangs() []string {
	url := "https://github.com/trending"

	// access to github
	resp, err := http.Get(url)
	if err != nil {
		print("Can not access to " + url)
		print(err)
		os.Exit(1)
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		print("Can not get langs from " + url)
		print(err)
		os.Exit(1)
	}

	var langs []string

	doc.Find("details > details-menu > div.select-menu-list > div > a").Each(func(i int, s *goquery.Selection) {
		lang, _ := s.Attr("href")
		langURL := cleansing(lang)
		rep := regexp.MustCompile(`.*/(.*)\?since=daily`)
		query := rep.ReplaceAllString(langURL, "$1")
		langs = append(langs, query)
	})

	return langs

}

func getRepositoryBySelection(s *goquery.Selection) Repository {
	name := cleansing(s.Find("div.d-inline-block.col-9.mb-1 > h3 > a").Text())
	url, _ := s.Find("a").Attr("href")
	description := cleansing(s.Find("div.py-1").Text())
	lang := cleansing(s.Find("span[itemprop='programmingLanguage']").Text())
	star := cleansingNum(s.Find("div.f6.text-gray.mt-2 > a:nth-child(2)").Text())
	fork := cleansingNum(s.Find("div.f6.text-gray.mt-2 > a:nth-child(3)").Text())
	starBySpan := cleansingNum(
		strings.Replace(s.Find(".float-sm-right").Text(), "stars today", "", -1))

	return Repository{Name: name, URL: "https://github.com" + url,
		Description: description,
		Lang:        lang, Star: star,
		Fork:       fork,
		StarBySpan: starBySpan}

}

// remove \n and trim
func cleansing(value string) string {
	return strings.Trim(strings.Replace(value, "\n", "", -1), " ")
}

// cleansing and parse int
func cleansingNum(value string) int {
	val, _ := strconv.Atoi(strings.Replace(cleansing(value), ",", "", -1))
	return val
}

// Print data for Repository
func (repo *Repository) Print() {
	fmt.Println(repo.ToString())
}

// ToString for Repository
func (repo *Repository) ToString() string {
	val := "【" + repo.Name + "】(" + repo.URL + ")\n"
	val = val + "Lang:" + repo.Lang + "\t" +
		"Fork:" + strconv.Itoa(repo.Fork) + "\t" +
		"⭐️" + strconv.Itoa(repo.Star) + "\t" +
		"⭐️" + strconv.Itoa(repo.StarBySpan) + " stars today" + "\n"
	val = val + repo.Description

	return val
}

func getQueryForSpan(span Span) string {
	switch span {
	case Today:
		return "daily"
	case Week:
		return "weekly"
	case Month:
		return "monthly"
	default:
		return "daily"
	}
}

// GetSpanByString Span const by string
func GetSpanByString(span string) Span {
	switch span {
	case "today":
		return Today
	case "week":
		return Week
	case "month":
		return Month
	default:
		return Today
	}
}
