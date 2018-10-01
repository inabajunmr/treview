package github

import (
	"strconv"

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
		name := s.Find("a").Text()
		url, _ := s.Find("a").Attr("href")
		description := s.Find("div.py-1").Text()
		lang := s.Find("span[itemprop='programmingLanguage']").Text()
		star, _ := strconv.Atoi(s.Find("div.f6.text-gray.mt-2 > a:nth-child(2)").Text())
		fork, _ := strconv.Atoi(s.Find("div.f6.text-gray.mt-2 > a:nth-child(3)").Text())
		starBySpan, _ := strconv.Atoi(s.Find("ddiv.f6.text-gray.mt-2 > span.d-inline-block.float-sm-right").Text())

		repo := Repository{Name: name, URL: url, Description: description,
			Lang: lang, Star: star,
			Fork:       fork,
			StarBySpan: starBySpan}

		repos[i] = repo
	})

	return repos, nil
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
