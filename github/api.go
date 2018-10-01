package github

import (
	"io/ioutil"
	"net/http"
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
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	body := string(bytes)

	// TODO parse html and struct Repositories
	return nil, nil
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
