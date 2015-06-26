package main

import c "github.com/hstove/gender/classifier"

func predictPercentFemale(names []string) float64 {
	classifier := c.Classifier()

	var numFemale int
	for _, name := range names {
		gender, _ := c.Classify(classifier, name)

		if gender == string(c.Girl) {
			numFemale += 1
		}
	}

	return (float64(numFemale) / float64(len(names)) * 100)
}
