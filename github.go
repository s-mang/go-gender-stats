package main

import (
	"database/sql"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// Contributor contains all info of a contributer to a repo
type Contributor struct {
	ID        int
	Name      string
	Firstname string
	Language  string
}

var (
	db     *sql.DB
	nameRE = regexp.MustCompile("^[A-Z]*[a-z]+$")
)

const limit = 10000
const routines = 8

func getGitHubContributors() []Contributor {
	var err error
	db, err = sql.Open("mysql", "user:pass@host/ghtorrent") // local mysql database with the ghtorrent sql dump
	if err != nil {
		panic(err)
	}

	contributors := []Contributor{}

	contributorChan := make(chan Contributor)
	doneChan := make(chan bool)
	offset := 0
	running := 0
	max := getLargestID()

	for i := 0; i < routines; i++ {
		go getContributorsForOffset(offset, limit, contributorChan, doneChan)
		offset += limit
		running++
	}

L:
	for {
		select {
		case contributor := <-contributorChan:
			contributors = append(contributors, contributor)
		case <-doneChan:
			fmt.Println("fetching:", strconv.FormatFloat((float64(offset)/float64(max)*100.0), 'f', 4, 64), "%")
			if offset <= max {
				go getContributorsForOffset(offset, limit, contributorChan, doneChan)
				offset += limit
			} else {
				running--
			}
			if running <= 0 {
				break L
			}
		}
	}
	return contributors
}

func getContributorsForOffset(offset, limit int, contributorsChan chan Contributor, done chan bool) {
	rows, err := db.Query("SELECT * FROM (SELECT u.id, language, u.name FROM commits c JOIN project_languages l ON c.project_id=l.project_id JOIN users u ON c.author_id=u.id WHERE c.id >= ? AND c.id < ? AND u.name is not null) q GROUP BY q.`language`, q.name, q.id", offset, offset+limit)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		contributor := Contributor{}
		err := rows.Scan(&contributor.ID, &contributor.Language, &contributor.Name)
		if err != nil {
			log.Fatal(err)
		}
		nameParts := strings.Split(contributor.Name, " ")
		if len(nameParts) >= 2 && nameRE.Match([]byte(nameParts[0])) {
			contributor.Firstname = nameParts[0]
			contributorsChan <- contributor
		}
	}
	done <- true
}

func getFirstNamesPerLanguage(names []Contributor) map[string][]string {
	num := len(names)
	fmt.Println(num)
	results := map[string][]Contributor{}
	offset := 0
	running := 0
	resultChan := make(chan map[string][]Contributor)
	for i := 0; i < routines; i++ {
		go makeNameLanguageMap(getCutUpSlice(names, offset, limit), resultChan)
		offset += limit
		running++
	}
L:
	for {
		result := <-resultChan
		for lang, names := range result {
			if _, haskey := results[lang]; !haskey {
				results[lang] = []Contributor{}
			}
			results[lang] = append(results[lang], names...)
		}
		fmt.Println("Processing:", strconv.FormatFloat((float64(offset)/float64(num)*100.0), 'f', 4, 64), "%")
		if offset <= num {
			go makeNameLanguageMap(getCutUpSlice(names, offset, limit), resultChan)
			offset += limit
		} else {
			running--
		}
		if running <= 0 {
			break L
		}
	}
	running = 0
	uniqueResults := map[string][]string{}
	namesChan := make(chan map[string][]string)
	for lang, contributors := range results {
		go filterDuplicateNames(lang, contributors, namesChan)
		running++
	}

M:
	for {
		names := <-namesChan
		for lang := range names {
			uniqueResults[lang] = names[lang]
		}
		running--
		if running <= 0 {
			break M
		}
	}

	return uniqueResults
}

func getCutUpSlice(names []Contributor, offset, limit int) []Contributor {
	num := len(names)
	out := []Contributor{}
	for i := offset; i < limit; i++ {
		if i < num {
			out = append(out, names[i])
		}
	}
	return out
}

func filterDuplicateNames(lang string, contributors []Contributor, out chan map[string][]string) {
	idsInSlice := map[int]bool{}
	firstNames := []string{}
	for _, contributor := range contributors {
		if _, exists := idsInSlice[contributor.ID]; !exists {
			firstNames = append(firstNames, contributor.Firstname, strconv.Itoa(contributor.ID))
			idsInSlice[contributor.ID] = true
		}
	}
	output := map[string][]string{}
	output[lang] = firstNames
	out <- output
}

func makeNameLanguageMap(names []Contributor, out chan map[string][]Contributor) {
	results := map[string][]Contributor{}
	for _, contributor := range names {
		if results[contributor.Language] == nil {
			results[contributor.Language] = []Contributor{}
		}
		results[contributor.Language] = append(results[contributor.Language], contributor)
	}
	out <- results
}

func getLargestID() int {
	rows, err := db.Query("SELECT id FROM commits ORDER BY id DESC LIMIT 1")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var num int
	rows.Next()
	rows.Scan(&num)
	return num
}
