package github

import (
	"fmt"
	"testing"
)

func TestFind(t *testing.T) {
	repos, err := Find("", Today)

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
	repos, err := Find("go", Today)

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
