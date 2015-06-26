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

	percentFemale, percentMale := predictGenderStats(names)
	percentUnknown := (100 - percentFemale - percentMale)

	fmt.Println("\nGo Contributors by Gender:")
	fmt.Printf("\n  - Female: %.2f%%\n", percentFemale)
	fmt.Printf("\n  - Male: %.2f%%\n", percentMale)
	fmt.Printf("\n  - Unknown: %.2f%%\n\n", percentUnknown)
}
