package main

// See bigquery_go_committers.sql for bigquery table creation.

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"google.golang.org/api/iterator"

	"cloud.google.com/go/bigquery"
)

var (
	nameRE    = regexp.MustCompile("^[A-Z]*[a-z]+$")
	projectID = os.Getenv("GOCLOUD_PROJECT_ID")
	bqDataset = os.Getenv("GH_BIGQUERY_DATASET")
)

func getGoCommitterNames() ([]string, error) {
	ctx := context.Background()
	client, err := bigquery.NewClient(ctx, os.Getenv("GOCLOUD_PROJECT_ID"))
	if err != nil {
		return nil, err
	}

	qstr := fmt.Sprintf(`
    SELECT first_name
    FROM [%s:%s.go_committers]
`, projectID, bqDataset)

	q := client.Query(qstr)

	it, err := q.Read(ctx)
	if err != nil {
		return nil, err
	}

	var names []string

	for {
		var c GoCommitter
		err := it.Next(&c)
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Println(err.Error())
			continue
		}

		var committerNames []string

		// some have comma separated, multiple individuals
		if strings.Contains(c.FirstName, ", ") {
			commitNames := strings.Split(c.FirstName, ", ")
			for _, c := range commitNames {
				c = strings.TrimPrefix(c, "and ")
				committerNames = append(committerNames, c)
			}
		} else if strings.Contains(c.FirstName, " ") {
			// some people put their full name as first name
			firstName := strings.Split(c.FirstName, " ")[0]
			committerNames = append(committerNames, firstName)
		}

		for _, n := range committerNames {
			// if does not follow nameRE, cannot be safely considered
			if !nameRE.Match([]byte(n)) {
				continue
			}
			names = append(names, n)
		}
	}

	return names, nil
}

type GoCommitter struct {
	FirstName string `bigquery:"first_name"`
}
