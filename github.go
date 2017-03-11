package main

import (
	"fmt"
	"log"
	"strings"

	"google.golang.org/api/iterator"

	"cloud.google.com/go/bigquery"
	"golang.org/x/net/context"
)

// Contributor contains all info of a contributer to a repo
type Contributor struct {
	Name      string
	Firstname string
	Repo      string
	Language  string
}

func getGitHubNamesPerLanguage() []Contributor {
	ctx := context.Background()

	client, err := bigquery.NewClient(ctx, "289749835264")
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	myDataset := client.Dataset("tmp_data")
	if err := myDataset.Create(ctx); err != nil {
		panic(err)
		// TODO: Handle error.
	}

	query := client.Query("SELECT commits.author.name as name, commits.repo_name as repo, language.name as language  FROM FLATTEN([bigquery-public-data:github_repos.commits], repo_name) commits JOIN [bigquery-public-data:github_repos.languages] languages ON commits.repo_name=languages.repo_name LIMIT 10")
	query.AllowLargeResults = true
	query.Dst = myDataset.Table("names_lang")
	it, err := query.Read(ctx)
	if err != nil {
		panic(err)
	}

	names := []Contributor{}

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
		c.Firstname = extractGitHubFirstName(c.Name)
		names = append(names, c)
	}
	return names
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
