package filter

import (
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/inabajunmr/treview/github/trending"
	"github.com/jonboulle/clockwork"
)

func TestOnlyNewComer(t *testing.T) {
	test1 := trending.Repository{Name: uuid.New().String()}
	test2 := trending.Repository{Name: uuid.New().String()}
	test3 := trending.Repository{Name: uuid.New().String()}
	test4 := trending.Repository{Name: uuid.New().String()}

	initialRepos := []trending.Repository{test1, test2, test3}

	// First day
	dir := os.TempDir()

	f1 := Filter{Time: clockwork.NewFakeClockAt(time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC)), Path: dir}
	repos := f1.OnlyNewComer(initialRepos)

	// All new comer
	if len(repos) != 3 {
		t.Fatal("Not return all repositories.")
	}

	// First day(but second time)
	secondRepos := []trending.Repository{test1, test2}
	repos = f1.OnlyNewComer(secondRepos)

	// All new comer
	if len(repos) != 2 {
		t.Fatal("Not return all repositories.")
	}

	// Second day
	f2 := Filter{Time: clockwork.NewFakeClockAt(time.Date(2001, 1, 2, 0, 0, 0, 0, time.UTC)), Path: dir}
	thirdRepos := []trending.Repository{test1, test4}

	repos = f2.OnlyNewComer(thirdRepos)
	if len(repos) != 1 {
		t.Fatal("Not return only 1 repository.")
	}

	if repos[0] != test4 {
		t.Fatal("Not return expected repository.")
	}
}

func TestIsSameDate(t *testing.T) {
	d1 := time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC)
	d2 := time.Date(2001, 1, 1, 12, 0, 0, 0, time.UTC)
	d3 := time.Date(2001, 1, 2, 0, 0, 0, 0, time.UTC)
	d4 := time.Date(2002, 1, 1, 0, 0, 0, 0, time.UTC)

	if !isSameDate(d1, d2) {
		// Different hour but same date.
		t.Fatal("Same date.")
	}

	if isSameDate(d1, d3) {
		// Different date
		t.Fatal("Same date.")
	}

	if isSameDate(d1, d4) {
		// Different year
		t.Fatal("Same date.")
	}

}

func TestDistinctByName(t *testing.T) {
	test1 := trending.Repository{Name: "test1"}
	test2 := trending.Repository{Name: "test2"}
	test3 := trending.Repository{Name: "test3"}

	repos := []trending.Repository{test1, test1, test2, test3, test3}
	results := DistinctRepository(repos)

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
