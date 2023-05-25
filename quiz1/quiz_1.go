package quiz1

import (
	"encoding/csv"
	"fmt"
	"os"
)

func quiz1() {
	fileName := "problems.csv"
	file, err := os.Open(fileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s", fileName))

	}
	r := csv.NewReader(file)
	_ = r
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
