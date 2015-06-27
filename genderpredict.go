package main

import (
	"fmt"
	"strings"

	c "github.com/hstove/gender/classifier"
)

func predictGenderStats(names []string) (f, m float64) {
	classifier := c.Classifier()

	var numFemale int
	var femaleNames []string

	var numMale int
	var maleNames []string

	for _, name := range names {
		gender, _ := c.Classify(classifier, name)

		if gender == string(c.Girl) {
			numFemale += 1
			femaleNames = append(femaleNames, name)
		}

		if gender == string(c.Boy) {
			numMale += 1
			maleNames = append(maleNames, name)
		}
	}

	printNames(maleNames, femaleNames)

	numTotal := len(names)

	f = percent(numFemale, numTotal)
	m = percent(numMale, numTotal)

	return
}

func printNames(maleNames, femaleNames []string) {
	fmt.Printf(
		"\nMALE (%d):\n%s\n",
		len(maleNames),
		strings.Join(maleNames, "\n"),
	)

	fmt.Printf(
		"\nFEMALE (%d):\n%s\n",
		len(femaleNames),
		strings.Join(femaleNames, "\n"),
	)
}

func percent(numForGender, numTotal int) float64 {
	return (float64(numForGender) / float64(numTotal) * 100)
}
