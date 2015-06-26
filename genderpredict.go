package main

import c "github.com/hstove/gender/classifier"

func predictGenderStats(names []string) (f, m float64) {
	classifier := c.Classifier()

	var numFemale int
	var numMale int
	for _, name := range names {
		gender, _ := c.Classify(classifier, name)

		if gender == string(c.Girl) {
			numFemale += 1
		}

		if gender == string(c.Boy) {
			numMale += 1
		}
	}

	numTotal := len(names)

	f = percent(numFemale, numTotal)
	m = percent(numMale, numTotal)

	return
}

func percent(numForGender, numTotal int) float64 {
	return (float64(numForGender) / float64(numTotal) * 100)
}
