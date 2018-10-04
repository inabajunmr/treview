package github

import (
	"fmt"
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

	for _, repo := range repos {
		fmt.Println("------------------------")
		repo.Print()
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

	for _, repo := range repos {
		fmt.Println("========================")
		repo.Print()
	}
}
