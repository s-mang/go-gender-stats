package main

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

const contribURL = "https://golang.org/CONTRIBUTORS?m=text"

var firstNameRegexp = regexp.MustCompile(`\n[^#\n ]+`)

func getContributorNames() ([]string, error) {
	resp, err := http.Get(contribURL)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return extractNames(data), nil

}

func extractNames(data []byte) []string {
	firstNames := firstNameRegexp.FindAllString(string(data), -1)

	for i := range firstNames {
		firstNames[i] = strings.Replace(firstNames[i], "\n", "", -1)
	}

	return firstNames
}
