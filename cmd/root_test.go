package cmd

import (
	"testing"

	"github.com/inabajunmr/treview/github/trending"
)

func TestDistinctByName(t *testing.T) {
	test1 := trending.Repository{Name: "test1"}
	test2 := trending.Repository{Name: "test2"}
	test3 := trending.Repository{Name: "test3"}

	repos := []trending.Repository{test1, test1, test2, test3, test3}
	results := distinctByName(repos)

	if len(results) != 3 {
		t.Fatal("Unexpected result.")
	}

	if results[0] != test1 {
		t.Fatal("Unexpected result.")
	}

	if results[1] != test2 {
		t.Fatal("Unexpected result.")
	}

	if results[2] != test3 {
		t.Fatal("Unexpected result.")
	}
}
