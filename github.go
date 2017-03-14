package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"golang.org/x/net/context"

	"google.golang.org/api/iterator"

	"cloud.google.com/go/bigquery"
)

// Contributor contains all info of a contributer to a repo
type Contributor struct {
	Name      string
	Firstname string
	Repo      string
	Language  string
}

type language struct {
	Name string
}

var (
	bqClient      *bigquery.Client
	hasToContinue bool
)

const authorsQueryString = "SELECT author.name as name, repo_name as repo FROM [bigquery-public-data:github_repos.commits] GROUP BY name, repo ORDER BY repo LIMIT 500 OFFSET "
const languagesQueryString = "SELECT language.name as name from [bigquery-public-data:github_repos.languages] WHERE repo_name="

func getGitHubNamesPerLanguage() []Contributor {
	var err error
	bqClient, err = bigquery.NewClient(context.Background(), "289749835264")

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	runningQueries := 0
	offset := 0
	contributorsChan := make(chan Contributor)
	doneChan := make(chan bool)
	names := []Contributor{}

	hasToContinue = true

	for i := 0; i < 30; i++ { // start 30 routines
		runningQueries++
		go getContributorsForOffset(offset, contributorsChan, doneChan)
		offset += 500
	}
L:
	for {
		select {
		case <-doneChan:
			fmt.Println("worker finished")
			if hasToContinue {
				go getContributorsForOffset(offset, contributorsChan, doneChan)
				offset += 100
			} else {
				runningQueries--
				if runningQueries == 0 {
					break L
				}
			}

		case contributor := <-contributorsChan:
			names = append(names, contributor)
		}
	}

	return names
}

func getContributorsForOffset(offset int, contributorsChan chan Contributor, done chan bool) {
	projectLanguages := map[string][]string{} // cache languages inside the goroutine

	query := bqClient.Query(authorsQueryString + strconv.Itoa(offset))
	it, err := query.Read(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	count := 0
	for {
		c := Contributor{}
		err := it.Next(&c)
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		count++
		c.Firstname = extractGitHubFirstName(c.Name)
		addLanguagesForContributor(c, contributorsChan, projectLanguages)
	}
	if count < 100 {
		hasToContinue = false
	}
	done <- true
}

func addLanguagesForContributor(c Contributor, contributorsChan chan Contributor, projectLanguages map[string][]string) {
	if projectLanguages[c.Repo] != nil {
		for _, language := range projectLanguages[c.Repo] {
			c.Language = language
			contributorsChan <- c
		}
		return
	}
	query := bqClient.Query(languagesQueryString + "\"" + c.Repo + "\"")
	it, err := query.Read(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	languages := []string{}
	for {
		var values []bigquery.Value
		err := it.Next(&values)
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		for _, language := range values {
			if language != nil {
				c.Language = language.(string)
				languages = append(languages, language.(string))
				contributorsChan <- c
			}
		}
	}
	projectLanguages[c.Repo] = languages
}

func getFirstNamesPerLanguage(names []Contributor) map[string][]string {
	results := map[string][]string{}

	for _, contributor := range names {
		if results[contributor.Language] == nil {
			results[contributor.Language] = []string{}
		}
		results[contributor.Language] = append(results[contributor.Language], contributor.Firstname)
	}
	return results
}

func extractGitHubFirstName(data string) string {
	return strings.Split(data, " ")[0]
}
