package main

import (
	"fmt"
	"log"
)

var logger *log.Logger

func main() {
	names, err := getContributorNames()
	if err != nil {
		panic(err.Error())
	}

	percentFemale := predictPercentFemale(names)

	fmt.Printf("\nFemale Go Contributors: %.2f%%\n\n", percentFemale)
}
