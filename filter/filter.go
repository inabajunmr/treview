package filter

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/inabajunmr/kevaf"
	"github.com/inabajunmr/treview/github/trending"
	"github.com/jonboulle/clockwork"
)

// Filter is filter for repository
type Filter struct {
	Time clockwork.Clock
	Path string
}

type readRepository struct {
	Repository trending.Repository `json:"repository"`
	ReadDate   time.Time           `json:"time"`
}

// OnlyNewComer filter only new comer(user never see)
func (f *Filter) OnlyNewComer(repos []trending.Repository) []trending.Repository {

	if !exists(f.Path) {
		if err := os.Mkdir(f.Path, 0777); err != nil {
			fmt.Println(err)
			print("Can not make dir for treview. path:"+f.Path, err)
			os.Exit(1)
		}
	}
	kvs, err := kevaf.NewMap(f.Path)
	if err != nil {
		print("Unexpected error. Can not initialize kevaf.", err)
		os.Exit(1)
	}

	filteredRepos := make([]trending.Repository, 0)

	for _, repo := range repos {
		key := createKey(repo.Name)

		v, err := kvs.Get(key)
		if err != nil {
			// if can not get, user has never seen this repository
			filteredRepos = append(filteredRepos, repo)
			markAsRead(kvs, repo, f.Time)
			continue
		}

		if f.isVisible(v, repo) {
			filteredRepos = append(filteredRepos, repo)
		}
	}

	return filteredRepos
}

func (f Filter) isVisible(savedRepo []byte, repo trending.Repository) bool {
	var readRepo readRepository
	err := json.Unmarshal(savedRepo, &readRepo)
	if err != nil {
		print("Unexpected error. Can't deserialize this repository.")
		print(repo.ToString())
		return true
	}

	// I wanna see repository as new comer that I first see this date.
	if isSameDate(readRepo.ReadDate, f.Time.Now()) {
		return true
	}

	return false
}

func markAsRead(kvs *kevaf.Map, repo trending.Repository, time clockwork.Clock) {
	val, err := json.Marshal(&readRepository{repo, time.Now()})
	if err != nil {
		print("Unexpected error. Can't serialize this repository.")
		print(repo.ToString())
		fmt.Println(err)
		return
	}

	err = kvs.Put(createKey(repo.Name), val)
	if err != nil {
		print("Unexpected error. Can't mark as read this repository.")
		print(repo.ToString())
		return
	}
}

func createKey(v string) string {
	return strings.Replace(strings.Replace(v, " ", "", -1), "/", "_", -1)
}

func isSameDate(time1 time.Time, time2 time.Time) bool {
	if time1.Year() != time2.Year() {
		return false
	}

	if time1.YearDay() != time2.YearDay() {
		return false
	}

	return true
}

func exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

// DistinctRepository by repository name
func DistinctRepository(repos []trending.Repository) []trending.Repository {

	var distinctRepos []trending.Repository
	repoNames := map[string]bool{}
	for _, repo := range repos {
		if !repoNames[repo.Name] {
			repoNames[repo.Name] = true
			distinctRepos = append(distinctRepos, repo)
		}
	}

	return distinctRepos

}
