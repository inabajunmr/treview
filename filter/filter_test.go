package filter

import (
	"os"
	"testing"
	"time"

	"github.com/inabajunmr/kevaf"
	"github.com/inabajunmr/treview/github"
	"github.com/jonboulle/clockwork"
)

func TestOnlyNewComer(t *testing.T) {
	test1 := github.Repository{Name: "test1"}
	test2 := github.Repository{Name: "test2"}
	test3 := github.Repository{Name: "test3"}
	test4 := github.Repository{Name: "test4"}

	initialRepos := []github.Repository{test1, test2, test3}

	// First day
	dir := os.TempDir()
	m, _ := kevaf.NewMap(dir)
	m.RemoveAll()

	f1 := Filter{Time: clockwork.NewFakeClockAt(time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC)), Path: dir}
	repos := f1.OnlyNewComer(initialRepos)

	// All new comer
	if len(repos) != 3 {
		t.Fatal("Not return all repositories.")
	}

	// First day(but second time)
	secondRepos := []github.Repository{test1, test2}
	repos = f1.OnlyNewComer(secondRepos)

	// All new comer
	if len(repos) != 2 {
		t.Fatal("Not return all repositories.")
	}

	// Second day
	f2 := Filter{Time: clockwork.NewFakeClockAt(time.Date(2001, 1, 2, 0, 0, 0, 0, time.UTC)), Path: dir}
	thirdRepos := []github.Repository{test1, test4}

	repos = f2.OnlyNewComer(thirdRepos)
	if len(repos) != 1 {
		t.Fatal("Not return only 1 repository.")
	}

	if repos[0] != test4 {
		t.Fatal("Not return expected repository.")
	}
}
