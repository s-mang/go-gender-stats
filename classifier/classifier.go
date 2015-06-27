package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	c "github.com/hstove/gender/classifier"
	b "github.com/jbrukh/bayesian"
)

func worker(classifier *b.Classifier, filename string) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		panic(err.Error())
	}

	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		count, err := strconv.ParseInt(record[2], 10, 64)
		if err != nil {
			panic(err.Error())
		}

		name := strings.ToLower(record[0])
		idx := 0
		for idx <= int(count) {
			c.Learn(classifier, name, record[1])
			idx++
		}

	}
	fmt.Println("done.")
}

func main() {
	classifier := c.NewClassifier()

	worker(classifier, "uknames.csv")
	worker(classifier, "usnames.csv")
	classifier.WriteToFile("../classifier.serialized")
}
