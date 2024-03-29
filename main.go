package main

import (
	"fmt"
	"log"
)

var logger *log.Logger

func main() {
	// Go contributors
	names, err := getContributorNames()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("\nGo Contributors by Gender:")
	printStats(names)

	// Gophers slack
	names, err = getGopherNames()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("-------------")
	fmt.Println("\nSlack Gophers by Gender:")
	printStats(names)

	// Github Go committers
	// Gophers slack
	names, err = getGoCommitterNames()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("-------------")
	fmt.Println("\nGithub Go Committers by Gender:")
	printStats(names)
}

func printStats(names []string) {
	percentFemale, percentMale := predictGenderStats(names)
	percentUnknown := (100 - percentFemale - percentMale)

	fmt.Printf("\n  - Female: %.2f%%\n", percentFemale)
	fmt.Printf("\n  - Male: %.2f%%\n", percentMale)

	if percentUnknown > 0 {
		fmt.Printf("\n  - Unknown: %.2f%%\n", percentUnknown)
	}
	fmt.Printf("\n  # Total: %d\n", len(names))

	fmt.Printf("\n")

}
