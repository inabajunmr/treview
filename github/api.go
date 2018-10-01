package github

import (
	"fmt"
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

// Find trending repositories by lang and span
func Find(lang string, span Span) ([]Repository, error) {
	// assemble url for trending
	url := "https://github.com/trending/" + lang + "?" + "since=" + getQueryForSpan(span)

	// access to github
	doc, err := goquery.NewDocument(url)

	if err != nil {
		return nil, err
	}

	// correct repositories
	repos := make([]Repository, 25)
	doc.Find("div.explore-content > ol > li").Each(func(i int, s *goquery.Selection) {
		name := cleansing(s.Find("div.d-inline-block.col-9.mb-1 > h3 > a").Text())
		url, _ := s.Find("a").Attr("href")
		description := cleansing(s.Find("div.py-1").Text())
		lang := cleansing(s.Find("span[itemprop='programmingLanguage']").Text())
		star := cleansingNum(s.Find("div.f6.text-gray.mt-2 > a:nth-child(2)").Text())
		fork := cleansingNum(s.Find("div.f6.text-gray.mt-2 > a:nth-child(3)").Text())
		starBySpan := cleansingNum(
			strings.Replace(s.Find(".float-sm-right").Text(), "stars today", "", -1))

		repo := Repository{Name: name, URL: "https://github.com" + url,
			Description: description,
			Lang:        lang, Star: star,
			Fork:       fork,
			StarBySpan: starBySpan}

		repos[i] = repo
	})

	return repos, nil
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
	fmt.Println("Name:" + repo.Name)
	fmt.Println("URL:" + repo.URL)
	fmt.Println("Description:" + repo.Description)
	fmt.Println("Lang:" + repo.Lang)
	fmt.Println("Star:" + strconv.Itoa(repo.Star))
	fmt.Println("StarBySpan:" + strconv.Itoa(repo.StarBySpan))
	fmt.Println("Fork:" + strconv.Itoa(repo.Fork))
}

func getQueryForSpan(span Span) string {
	switch span {
	case Today:
		return "today"
	case Week:
		return "week"
	case Month:
		return "month"
	default:
		return "today"
	}
}
