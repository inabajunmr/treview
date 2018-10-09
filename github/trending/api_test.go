package trending

import (
	"testing"
)

func TestFindTrending(t *testing.T) {
	repos, err := FindTrending("", Today)

	if err != nil {
		t.Fatal("Error by find.", err)
	}

	if len(repos) != 25 {
		t.Fatal("Not return 25 repositories.", err)
	}
}

func TestFindSpecificLang(t *testing.T) {
	repos, err := FindTrending("go", Today)

	if err != nil {
		t.Fatal("Error by find.", err)
	}

	if len(repos) != 25 {
		t.Fatal("Not return 25 repositories.", err)
	}
}

func TestPrint(t *testing.T) {
	// This only and Pring
	r := Repository{Name: "testName", URL: "http://example.com/test",
		Description: "desc", Lang: "go", Star: 30, StarBySpan: 10, Fork: 10}

	r.Print()
}

func TestGetQueryForSpan(t *testing.T) {
	if value := getQueryForSpan(Today); value != "today" {
		t.Fatal("It's not expected value by Today. value:" + value)
	}

	if value := getQueryForSpan(Week); value != "week" {
		t.Fatal("It's not expected value by Week. value:" + value)
	}

	if value := getQueryForSpan(Month); value != "month" {
		t.Fatal("It's not expected value by Month. value:" + value)
	}

}

func TestGetSpanByString(t *testing.T) {
	if value := GetSpanByString("today"); value != Today {
		t.Fatal("It's not expected value by today. value:" + getQueryForSpan(value))
	}

	if value := GetSpanByString("week"); value != Week {
		t.Fatal("It's not expected value by week. value:" + getQueryForSpan(value))
	}

	if value := GetSpanByString("month"); value != Month {
		t.Fatal("It's not expected value by month. value:" + getQueryForSpan(value))
	}

	if value := GetSpanByString("other"); value != Today {
		t.Fatal("It's not expected value by today. value:" + getQueryForSpan(value))
	}

}
