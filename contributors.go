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
	percentMale := (100 - percentFemale)

	fmt.Println("\nGo Contributors by Gender:")
	fmt.Printf("\n  - Female: %.2f%%\n", percentFemale)
	fmt.Printf("\n  - Male: %.2f%%\n\n", percentMale)
}
